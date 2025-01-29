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
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"
	"time"
)

var Version string

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Unrecoverable error occurred: %s\n", e)
			fmt.Println("It seems to be a bug so please file an issue at https://youtrack.jetbrains.com/newIssue?project=ktor")
			fmt.Printf("Please put the following stack trace into the issue's description: \n\n%s", string(debug.Stack()))
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
	case cli.AddCommand:
		modules := args.CommandArgs

		projectDir := "."
		if dir, ok := args.CommandOptions[cli.ProjectDir]; ok {
			projectDir = dir
		}

		tomlPath, tomlFound := toml.FindVersionsPath(projectDir)
		var tomlDoc *toml.Document
		var tomlSuccessParsed bool

		buildPath := filepath.Join(projectDir, "build.gradle.kts")
		buildFound := utils.Exists(buildPath)
		var buildRoot *gradle.BuildRoot

		if buildFound {
			var syntaxErrors []lang.SyntaxError
			buildRoot, err, syntaxErrors = gradle.ParseBuildFile(buildPath)

			if len(syntaxErrors) > 0 {
				log.Println(lang.StringifySyntaxErrors(syntaxErrors[:5]))
			}

			if err != nil {
				cli.ExitWithError(err, hasGlobalLog, homeDir)
			} else {
				if tomlFound {
					tomlDoc, err, syntaxErrors = toml.ParseCatalogToml(tomlPath)

					if len(syntaxErrors) > 0 {
						log.Println(lang.StringifySyntaxErrors(syntaxErrors[:5]))
					}

					if err != nil {
						log.Println(err)
					} else {
						tomlSuccessParsed = true
					}
				}

				if project.IsKmp(buildRoot, tomlDoc, tomlSuccessParsed) {
					fmt.Fprintln(os.Stderr, "Unable to add the Ktor module to a Kotlin Multiplatform project (not supported yet).")
					os.Exit(1)
				}
			}
		} else {
			if utils.Exists(filepath.Join(projectDir, "pom.xml")) {
				fmt.Fprintln(os.Stderr, "Unable to add the Ktor module to a Maven project (not supported yet).")
			} else if utils.Exists(filepath.Join(projectDir, "build.gradle")) {
				fmt.Fprintln(os.Stderr, "Unable to add the Ktor module to a Gradle project with Groovy DSL (not supported yet).")
			} else {
				fmt.Fprintf(os.Stderr, "Unable to find build.gradle.kts file in the project directory %s.\n", projectDir)
			}
			os.Exit(1)
		}

		var ktorVersion string
		if v, ok := project.SearchKtorVersion(projectDir, buildRoot, tomlDoc, tomlSuccessParsed); ok {
			ktorVersion = v
			verboseLogger.Printf("Detected Ktor version: %s\n", ktorVersion)
		} else {
			settings, err := network.FetchSettings(client)

			if err != nil {
				cli.ExitWithError(err, hasGlobalLog, homeDir)
			} else {
				ktorVersion = settings.KtorVersion.DefaultId
				verboseLogger.Printf("Using the latest stable Ktor version: %s\n", ktorVersion)
			}
		}

		artifacts, err := network.SearchArtifacts(client, ktorVersion, modules)

		if err != nil {
			cli.ExitWithError(err, hasGlobalLog, homeDir)
		}

		for _, mod := range modules {
			mc, modResult, candidates := ktor.FindModule(artifacts[mod])

			if mc.Version == "" {
				mc.Version = ktorVersion
			}

			switch modResult {
			case ktor.ModuleNotFound:
				fmt.Fprintf(os.Stderr, "Cannot recognize Ktor module '%s'.\n", mod)
				os.Exit(1)
			case ktor.ModuleAmbiguity:
				var names []string
				for _, c := range candidates {
					if !slices.Contains(names, c.Artifact) {
						names = append(names, c.Artifact)
					}
				}
				fmt.Fprintf(os.Stderr, "Ktor Module ambiguity. Applicable candidates: %s.", strings.Join(names, ", "))
				os.Exit(1)
			case ktor.SimilarModulesFound:
				fmt.Fprintf(os.Stderr, "Cannot recognize module '%s'. ", mod)

				if len(candidates) > 0 {
					fmt.Fprintf(os.Stderr, "Did you mean '%s'?\n", candidates[0].Artifact)
				}
			case ktor.ModuleFound:
				verboseLogger.Printf("The chosen module is %s.\n", mc.String())
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
					fmt.Printf("Below you can find suggested changes to add '%s' into the project.\n", mc.String())
					fmt.Println("If you consider them incorrect, please file an issue at https://youtrack.jetbrains.com/newIssue?project=ktor.")
					fmt.Println()
					for _, f := range files {
						fmt.Println(utils.GetDiff(f.Path, f.Content))
					}

					fmt.Print("Do you want to apply the changes (y/n)? ")
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					answer := scanner.Text()

					if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" {
						err = project.ApplyChanges(files)

						if err == nil {
							fmt.Println("The changes have been successfully applied.")
						} else {
							cli.ExitWithError(err, hasGlobalLog, homeDir)
						}
					} else {
						fmt.Println("GoodBye!")
					}
				} else {
					fmt.Println("Nothing to change.")
				}
			}
		}
	case cli.VersionCommand:
		fmt.Printf(i18n.Get(i18n.VersionInfo, getVersion()))
	case cli.HelpCommand:
		cli.WriteUsage(os.Stdout)
	case cli.CompletionCommand:
		log.SetOutput(os.Stderr)

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
