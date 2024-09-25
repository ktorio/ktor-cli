package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/cli/command"
	"github.com/ktorio/ktor-cli/internal/app/config"
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
		verboseLogger.Println("Cannot determine home directory")
	}

	hasGlobalLog := false
	if hasHomeDir {
		logHandle, err := utils.PrepareGlobalLog(homeDir)

		if err != nil {
			verboseLogger.Printf("Unable to initialize log file: %s\n", config.LogPath(homeDir))
		}

		if err == nil {
			hasGlobalLog = true
			defer logHandle.Close()
		}
	}

	ctx := context.WithValue(context.Background(), "user-agent", fmt.Sprintf("KtorCLI/%s", getVersion()))

	switch args.Command {
	case cli.VersionCommand:
		fmt.Printf("Ktor CLI version %s\n", getVersion())
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
				fmt.Fprintf(os.Stderr, "Unable to determine project directory for project %s\n", projectName)
				os.Exit(1)
			}

			command.Generate(client, projectDir, projectName, verboseLogger, hasGlobalLog, ctx)
			return
		}

		if err := interactive.Run(client); err != nil {
			fmt.Fprintf(os.Stderr, "Interactive error: %v\n", err)
			os.Exit(1)
		}
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
