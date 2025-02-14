package main

import (
	"bufio"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/cli/command"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/interactive"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"github.com/ktorio/ktor-cli/internal/app/project"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"
	"time"
)

var Version string

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Print(i18n.Get(i18n.UnrecoverableErrorBlock, e, string(debug.Stack())))
		}
	}()

	args, err := cli.ProcessArgs(cli.ParseArgs(os.Args))

	if err != nil {
		cli.HandleArgsValidation(err)
	}

	homeDir, err := os.UserHomeDir()
	hasHomeDir := err == nil

	verboseLogger := log.New(os.Stdout, "", 0)

	if !args.Verbose {
		verboseLogger.SetOutput(io.Discard)
	}

	if !hasHomeDir {
		verboseLogger.Println(i18n.Get(i18n.CannotDetermineHomeDir))
	}

	hasGlobalLog := false
	if hasHomeDir {
		logHandle, err := utils.PrepareGlobalLog(homeDir)

		if err != nil {
			verboseLogger.Printf(i18n.Get(i18n.ErrorInitLogFile, config.LogPath(homeDir)))
		}

		if err == nil {
			hasGlobalLog = true
			defer logHandle.Close()
		}
	}

	ctx := context.WithValue(context.Background(), "user-agent", fmt.Sprintf("KtorCLI/%s", getVersion()))

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, 5*time.Second)
			},
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}

	switch args.Command {
	case cli.DevCommand:
		log.SetOutput(os.Stderr)
		projectDir := "."

		if dir, ok := args.CommandOptions[cli.ProjectDir]; ok {
			projectDir = dir
		} else if dir, err = filepath.Abs(projectDir); err == nil {
			projectDir = dir
		}

		wrapper := "./gradlew"
		if runtime.GOOS == "windows" {
			wrapper = ".\\gradlew.bat"
		}

		wrapperPath := filepath.Join(projectDir, wrapper)

		if !utils.Exists(wrapperPath) {
			fmt.Printf("Gradle wrapper %s doesn't exist in the project directory %s.\n", wrapper, projectDir) // TODO: i18n
			os.Exit(1)
		}

		runTask, buildTask, guessed := project.GuessGradleTasks(projectDir)

		if !guessed {
			fmt.Println("Cannot find the Ktor gradle plugin with appropriate version") // TODO: i18n
			os.Exit(1)
		}

		_, jdkPath, err := cli.ObtainJdk(client, verboseLogger, homeDir)

		if err != nil {
			cli.ExitWithError(err, hasGlobalLog, homeDir)
		}

		env := os.Environ()
		env = append(env, fmt.Sprintf("JAVA_HOME=%s", jdkPath))

		buildCmd := exec.Command(wrapper, buildTask, "--continuous")
		buildCmd.Env = env
		buildCmd.Dir = projectDir
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		verboseLogger.Printf("Starting build command: %s (JAVA_HOME=%s)\n", buildCmd.String(), jdkPath) // TODO: i18n

		err = buildCmd.Start() // TODO: Handle when build exits with error
		if err != nil {
			// TODO: Handle permissions error
			log.Fatal(err) // TODO: Handle error
		}

		runCmd := exec.Command(wrapper, runTask, "-Pdevelopment") // TODO: Fix after the Gradle plugin update
		runCmd.Env = env
		runCmd.Dir = projectDir
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr

		verboseLogger.Printf("Starting run command: %s (JAVA_HOME=%s)\n", runCmd.String(), jdkPath) // TODO: i18n

		err = runCmd.Start()
		if err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				os.Exit(exitErr.ExitCode())
			}

			log.Fatal(err) // TODO: Handle error
		}
		err = runCmd.Wait()
		if err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				os.Exit(exitErr.ExitCode())
			}

			log.Fatal(err) // TODO: Handle error
		}
	case cli.AddCommand:
		modules := args.CommandArgs

		projectDir := "."
		if dir, ok := args.CommandOptions[cli.ProjectDir]; ok {
			projectDir = dir
		}

		projectDir, err = filepath.Abs(projectDir)

		if err != nil {
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.CannotDetermineProjectDir, projectDir))
			os.Exit(1)
		}

		verboseLogger.Print(i18n.Get(i18n.ProjectAddMessage, projectDir))

		tomlPath, tomlFound := toml.FindVersionsPath(projectDir)
		var tomlDoc *toml.Document
		var tomlSuccessParsed bool

		buildPath := filepath.Join(projectDir, "build.gradle.kts")
		buildFound := utils.Exists(buildPath)
		var buildRoot *gradle.BuildRoot

		if buildFound {
			var syntaxErrors []lang.SyntaxError
			buildRoot, err, syntaxErrors = gradle.ParseBuildFile(buildPath)

			var sErrors []lang.SyntaxError
			if len(sErrors) > 0 {
				if len(syntaxErrors) < 5 {
					sErrors = syntaxErrors
				} else {
					sErrors = syntaxErrors[:5]
				}

				log.Println(lang.StringifySyntaxErrors(sErrors))
			}

			if err != nil {
				cli.ExitWithError(err, hasGlobalLog, homeDir)
			} else {
				if tomlFound {
					tomlDoc, err, syntaxErrors = toml.ParseCatalogToml(tomlPath)

					if len(syntaxErrors) > 0 {
						if len(syntaxErrors) < 5 {
							sErrors = syntaxErrors
						} else {
							sErrors = syntaxErrors[:5]
						}

						log.Println(lang.StringifySyntaxErrors(sErrors))
					}

					if err != nil {
						log.Println(err)
					} else {
						tomlSuccessParsed = true
					}
				}

				if project.IsKmp(buildRoot, tomlDoc, tomlSuccessParsed) {
					fmt.Fprintln(os.Stderr, i18n.Get(i18n.AddKtorModulesToKmpError))
					os.Exit(1)
				}
			}
		} else {
			if utils.Exists(filepath.Join(projectDir, "pom.xml")) {
				fmt.Fprintln(os.Stderr, i18n.Get(i18n.AddKtorModulesToMavenError))
			} else if utils.Exists(filepath.Join(projectDir, "build.gradle")) {
				fmt.Fprintln(os.Stderr, i18n.Get(i18n.AddKtorModulesToGradleGroovyError))
			} else {
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnableToFindBuildGradleKts, projectDir))
			}
			os.Exit(1)
		}

		var ktorVersion string
		if v, ok := project.SearchKtorVersion(projectDir, buildRoot, tomlDoc, tomlSuccessParsed); ok {
			ktorVersion = v
			verboseLogger.Printf(i18n.Get(i18n.DetectedKtorVersion, ktorVersion))
		} else {
			settings, err := network.FetchSettings(client)

			if err != nil {
				cli.ExitWithError(err, hasGlobalLog, homeDir)
			} else {
				ktorVersion = settings.KtorVersion.DefaultId
				verboseLogger.Printf(i18n.Get(i18n.UseLatestKtorVersion, ktorVersion))
			}
		}

		artifacts, err := network.SearchArtifacts(client, ktorVersion, modules)

		if err != nil {
			cli.ExitWithError(err, hasGlobalLog, homeDir)
		}

		fmt.Print(i18n.Get(i18n.ChangesWarningBlock, filepath.Base(projectDir)))

		for i, mod := range modules {
			if i > 0 {
				buildRoot, err, _ = gradle.ParseBuildFile(buildPath)

				if err != nil {
					cli.ExitWithError(err, hasGlobalLog, homeDir)
				}

				tomlDoc, err, _ = toml.ParseCatalogToml(tomlPath)
				tomlSuccessParsed = err == nil
			}

			mc, modResult, candidates := ktor.FindModule(artifacts[mod])

			if mc.Version == "" {
				mc.Version = ktorVersion
			}

			switch modResult {
			case ktor.ModuleNotFound:
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnableToRecognizeKtorModule, mod))
				os.Exit(1)
			case ktor.ModuleAmbiguity:
				var names []string
				for _, c := range candidates {
					if !slices.Contains(names, c.Artifact) {
						names = append(names, c.Artifact)
					}
				}
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.KtorModuleAmbiguity, strings.Join(names, ", ")))
				os.Exit(1)
			case ktor.SimilarModulesFound:
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnableToRecognizeKtorModule, mod))

				if len(candidates) > 0 {
					fmt.Fprintf(os.Stderr, i18n.Get(i18n.SimilarModuleQuestion, candidates[0].Artifact))
				}
				os.Exit(1)
			case ktor.ModuleFound:
				verboseLogger.Printf(i18n.Get(i18n.ChosenKtorModule, mc.String()))
				depPlugins := ktor.DependentPlugins(mc)
				var serPlugin *ktor.GradlePlugin
				if len(depPlugins) > 0 {
					serPlugin = &depPlugins[0]
				}

				files, err := project.AddKtorModule(mc, buildRoot, tomlDoc, tomlSuccessParsed, serPlugin, buildPath, tomlPath, projectDir)

				if err != nil {
					cli.ExitWithError(err, hasGlobalLog, homeDir)
				}

				if len(files) > 0 {
					fmt.Println()
					for _, f := range files {
						fmt.Println(utils.GetDiff(f.Path, f.Content))
					}

					fmt.Print(i18n.Get(i18n.ApplyChangesQuestion))
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					answer := scanner.Text()

					if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" {
						err = project.ApplyChanges(files)

						if err == nil {
							fmt.Println(i18n.Get(i18n.ChangesApplied))
						} else {
							cli.ExitWithError(err, hasGlobalLog, homeDir)
						}
					}
				} else {
					fmt.Println()
					fmt.Println(i18n.Get(i18n.NoChanges, mc.String()))
				}
			}
		}
	case cli.VersionCommand:
		fmt.Printf(i18n.Get(i18n.VersionInfo, getVersion()))
	case cli.HelpCommand:
		cli.WriteUsage(os.Stdout)
	case cli.CompletionCommand:
		settings, err := network.FetchSettings(client)

		if err != nil {
			cli.ExitWithError(err, hasGlobalLog, homeDir)
		}

		shell := args.CommandArgs[0]
		modules, err := network.ListArtifacts(client, settings.KtorVersion.DefaultId)

		if err != nil {
			cli.ExitWithError(err, hasGlobalLog, homeDir)
		}

		s, err := command.Complete(modules, shell)

		if err != nil {
			cli.ExitWithError(err, hasGlobalLog, homeDir)
		}

		fmt.Print(s)
	case cli.NewCommand:
		if len(args.CommandArgs) > 0 {
			projectName := utils.CleanProjectName(filepath.Base(args.CommandArgs[0]))
			projectDir, err := filepath.Abs(args.CommandArgs[0])

			if err != nil {
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.CannotDetermineProjectDirOfProject, projectName))
				os.Exit(1)
			}

			command.Generate(client, projectDir, projectName, []string{}, verboseLogger, hasGlobalLog, ctx)
			return
		}

		result, err := interactive.Run(client, ctx)
		if err != nil {
			cli.ExitWithError(err, hasGlobalLog, homeDir)
		}

		if result.Quit {
			fmt.Println(i18n.Get(i18n.ByeMessage))
			return
		}

		command.Generate(client, result.ProjectDir, result.ProjectName, result.Plugins, verboseLogger, hasGlobalLog, ctx)
	case cli.OpenAPI:
		specPath := args.CommandArgs[0]
		projectDir, err := filepath.Abs(".")

		if dir, ok := args.CommandOptions[cli.OutDir]; ok {
			projectDir, err = filepath.Abs(dir)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.CannotDetermineProjectDir, projectDir))
			os.Exit(1)
		}

		projectName := utils.CleanProjectName(filepath.Base(projectDir))

		if _, err := os.Stat(specPath); errors.Is(err, os.ErrNotExist) {
			fmt.Printf(i18n.Get(i18n.OpenApiSpecNotExist, specPath))
			os.Exit(1)
		}

		err = command.OpenApi(client, specPath, projectName, projectDir, homeDir, verboseLogger)

		if err != nil {
			cli.ExitWithProjectError(err, projectDir, hasGlobalLog, homeDir)
		}

		fmt.Printf(i18n.Get(i18n.ProjectCreated, projectName, projectDir))
	}
}

func getVersion() string {
	if Version != "" {
		return strings.Trim(Version, "\n\r")
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return fmt.Sprintf("dev-%s", setting.Value[:7])
			}
		}
	}

	return "dev"
}
