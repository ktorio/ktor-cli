package main

import (
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
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"golang.org/x/exp/slices"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

var Version string

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Unrecoverable error occurred: %s\n", e)
			fmt.Println("This looks like a bug so please file an issue at https://youtrack.jetbrains.com/newIssue?project=ktor.")
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
		log.SetOutput(os.Stderr)
		modules := args.CommandArgs
		projectDir := "."

		for _, mod := range modules {
			mc, modResult, candidates := ktor.FindModule(mod)

			fmt.Printf("Changes for module '%s':\n", mod)

			switch modResult {
			case ktor.ModuleNotFound:
				log.Fatal(fmt.Sprintf("Cannot recongnize Ktor module %s", mod))
			case ktor.ModuleAmbiguity:
				var names []string
				for _, c := range candidates {
					if !slices.Contains(names, c.Artifact) {
						names = append(names, c.Artifact)
					}
				}
				log.Fatal(fmt.Sprintf("Module ambiguity. Candidates: %s", strings.Join(names, ", ")))
			case ktor.AlikeModuleFound:
				log.Fatal(fmt.Sprintf("Cannot recognize the '%s' module.\nDid you mean '%s'?\n", mod, mc.Artifact))
			case ktor.ModuleFound:
				depPlugins := ktor.DependentPlugins(mc)
				var serPlugin *ktor.GradlePlugin
				if len(depPlugins) > 0 {
					serPlugin = &depPlugins[0]
				}

				err = command.Add(mc, projectDir, serPlugin)

				if err != nil {
					log.Fatal(err)
				}
			}
		}
	case cli.VersionCommand:
		fmt.Printf(i18n.Get(i18n.VersionInfo, getVersion()))
	case cli.HelpCommand:
		cli.WriteUsage(os.Stdout)
	case cli.CompletionCommand:
		log.SetOutput(os.Stderr)
		shell := args.CommandArgs[0]
		s, err := command.Complete(shell)

		if err != nil {
			log.Fatal(err)
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
			cli.ExitWithError(err, "", hasGlobalLog, homeDir)
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
			cli.ExitWithError(err, projectDir, hasGlobalLog, homeDir)
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
