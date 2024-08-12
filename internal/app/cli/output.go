package cli

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"os"
	"runtime"
)

func HandleAppError(projectDir string, err error) (reportLog bool) {
	if err == nil {
		return
	}

	reportLog = true
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
	}

	return
}

func HandleArgsValidation(err error, command *string) {
	switch err {
	case NoCommandError:
		fmt.Fprintln(os.Stderr, "Expected a command")
	case CommandNotFoundError:
		fmt.Fprintf(os.Stderr, "Command %s not found\n", *command)
	case WrongNumberOfArgumentsError:
		if Command(*command) == NewCommand {
			fmt.Fprintln(os.Stderr, "Expected one argument (project name) for the new command")
		}
	default:
		// do nothing
	}

	WriteUsage(os.Stderr)
}

func PrintSuccessGen(projectDir, projectName string) {
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
