package main

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/generate"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	args, err := cli.ParseArgs(os.Args)

	if err != nil {
		cli.WriteUsage(os.Stderr)
		os.Exit(1)
	}

	if err := cli.ValidateArgs(args); err != nil {
		switch err {
		case cli.NoCommandError:
			fmt.Fprintln(os.Stderr, "Expected a command")
		case cli.CommandNotFoundError:
			fmt.Fprintf(os.Stderr, "Command %s not found\n", *args.Command)
		case cli.WrongNumberOfArgumentsError:
			if *args.Command == "new" {
				fmt.Fprintln(os.Stderr, "Expected one argument (project name) for the new command")
			}
		default:
			// do nothing
		}

		cli.WriteUsage(os.Stderr)
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
	case "new":
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
			reportLog := true
			var e *app.Error
			if ok := errors.As(err, &e); ok {
				switch e.Kind {
				case app.ServerError:
					fmt.Fprintf(os.Stderr, "Unexpected error occurred while connecting to the generation server. Please try again later.\n")
				case app.NetworkError:
					fmt.Fprintf(os.Stderr, "Unexpected network error occurred while connecting to the generation server. Please check your Internet connection.\n")
				case app.InternalError:
					fmt.Fprintf(os.Stderr, "An intenal error occurred. Please file an issue on https://youtrack.jetbrains.com/newIssue?project=ktor\n")
				case app.ProjectDirError:
					reportLog = false
					var pe *os.PathError
					errors.As(e.Err, &pe)

					switch {
					case errors.Is(pe.Err, os.ErrExist):
						fmt.Fprintf(os.Stderr, "The project directory %s already exists.\n", pe.Path)
					case errors.Is(pe.Err, os.ErrPermission):
						fmt.Fprintf(os.Stderr, "Not enough permissions to create project directory %s.\n", pe.Path)
					}
				case app.ExtractError:
					fmt.Fprintf(os.Stderr, "Unable to extract downloaded archive to the directory %s.\n", projectDir)
				case app.GradlewChmod:
					var pe *os.PathError
					errors.As(e.Err, &pe)
					fmt.Fprintf(os.Stderr, "Unable to make %s file executable.\n", pe.Path)
				case app.UnknownError:
					fmt.Fprintf(os.Stderr, "Unexpected error occurred.\n")
				}

				if hasGlobalLog && reportLog {
					fmt.Fprintf(os.Stderr, "You can find more information in the log: %s\n", config.LogPath(homeDir))
				}
			}

			log.Fatal(err)
		}

		initialProjectDir := projectName
		fmt.Printf("Project \"%s\" has been created in the directory %s.\n", projectName, projectDir)
		fmt.Print("To run the project use the following commands:\n\n")

		if runtime.GOOS == "windows" {
			fmt.Printf("cd %s\n", initialProjectDir)
			fmt.Println(".\\gradlew.bat run")
		} else {
			fmt.Printf("cd %s\n", initialProjectDir)
			fmt.Println("./gradlew run")
		}
	}
}
