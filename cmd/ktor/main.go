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
	"github.com/ktorio/ktor-cli/internal/app/utils"
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
	case cli.VersionCommand:
		fmt.Printf(i18n.Get(i18n.VersionInfo, getVersion()))
	case cli.HelpCommand:
		cli.WriteUsage(os.Stdout)
	case cli.NewCommand:
		if len(args.CommandArgs) > 0 {
			projectName := utils.CleanProjectName(filepath.Base(args.CommandArgs[0]))
			projectDir, err := filepath.Abs(args.CommandArgs[0])

			if err != nil {
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.CannotDetermineProjectDir, projectName))
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
			fmt.Fprintf(os.Stderr, "Unable to determine project directory %s\n", projectDir)
			os.Exit(1)
		}

		projectName := utils.CleanProjectName(filepath.Base(projectDir))

		if _, err := os.Stat(specPath); errors.Is(err, os.ErrNotExist) {
			fmt.Printf("OpenAPI spec file %s does not exist\n", specPath)
			os.Exit(1)
		}

		err = command.OpenApi(client, specPath, projectName, projectDir, homeDir, verboseLogger)

		if err != nil {
			cli.ExitWithError(err, projectDir, hasGlobalLog, homeDir)
		}

		fmt.Printf("Project %s has been generated in the directory %s\n", projectName, projectDir)
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
