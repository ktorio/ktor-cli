package main

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/generate"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	args, err := cli.ParseArgs(os.Args)

	if err != nil {
		cli.WriteUsage(os.Stderr)
		os.Exit(1)
	}

	if err := cli.ValidateArgs(args); err != nil {
		cli.HandleArgsValidation(err, args.Command)
		os.Exit(1)
	}

	homeDir, err := os.UserHomeDir()
	hasHomeDir := err == nil

	verboseLogger := log.New(os.Stdout, "", 0)

	if !*args.Options.Verbose {
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

	switch *args.Command {
	case string(cli.NewCommand):
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		projectName := utils.CleanProjectName(args.CommandArgs[0])
		projectDir, err := filepath.Abs(projectName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to determine project directory for project %s\n", projectName)
			os.Exit(1)
		}

		err = generate.Project(client, verboseLogger, projectDir, projectName)
		if err != nil {
			reportLog := cli.HandleAppError(projectDir, err)

			if hasGlobalLog && reportLog {
				fmt.Fprintf(os.Stderr, "You can find more information in the log: %s\n", config.LogPath(homeDir))
			}

			log.Fatal(err)
		}

		cli.PrintSuccessGen(projectDir, projectName)
	}
}
