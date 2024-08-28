package main

import (
	_ "embed"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/cli/command"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	switch args.Command {
	case cli.VersionCommand:
		command.Version(Version)
	case cli.HelpCommand:
		cli.WriteUsage(os.Stdout)
	case cli.NewCommand:
		client := &http.Client{
			Timeout: 30 * time.Second,
		}

		projectName := utils.CleanProjectName(args.CommandArgs[0])
		projectDir, err := filepath.Abs(projectName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to determine project directory for project %s\n", projectName)
			os.Exit(1)
		}

		command.Generate(client, projectDir, projectName, verboseLogger, hasGlobalLog)
	}
}
