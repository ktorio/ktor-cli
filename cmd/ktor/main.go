package main

import (
	"context"
	_ "embed"
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

	switch args.Command {
	case cli.VersionCommand:
		fmt.Printf(i18n.Get(i18n.VersionInfo, getVersion()))
	case cli.HelpCommand:
		cli.WriteUsage(os.Stdout)
	case cli.NewCommand:
		client := &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, 5*time.Second)
				},
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
			},
		}

		if len(args.CommandArgs) > 0 {
			projectName := utils.CleanProjectName(args.CommandArgs[0])
			projectDir, err := filepath.Abs(projectName)

			if err != nil {
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.CannotDetermineProjectDir, projectName))
				os.Exit(1)
			}

			command.Generate(client, projectDir, projectName, []string{}, verboseLogger, hasGlobalLog, ctx)
			return
		}

		result, err := interactive.Run(client, ctx)
		if err != nil {
			reportLog := cli.HandleAppError("", err)

			if hasGlobalLog && reportLog {
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.LogHint, config.LogPath(homeDir)))
			}

			if hasGlobalLog {
				log.Fatal(err)
			}

			os.Exit(1)
		}

		if result.Quit {
			fmt.Println(i18n.Get(i18n.ByeMessage))
			return
		}

		command.Generate(client, result.ProjectDir, result.ProjectName, result.Plugins, verboseLogger, hasGlobalLog, ctx)
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
