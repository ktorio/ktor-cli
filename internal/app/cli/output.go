package cli

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"os"
	"runtime"
)

func HandleAppError(projectDir string, err error) (reportLog bool) {
	if err == nil {
		return
	}

	reportLog = true
	var e *app.Error
	if errors.As(err, &e) {
		switch e.Kind {
		case app.GenServerError:
			fmt.Fprintf(os.Stderr, "Unexpected error occurred while connecting to the generation server. Please try again later.\n")
		case app.NetworkError:
			fmt.Fprintf(os.Stderr, "Unexpected network error occurred while connecting to the generation server. Please check your Internet connection.\n")
		case app.InternalError:
			fmt.Fprintf(os.Stderr, "An internal error occurred. Please file an issue on https://youtrack.jetbrains.com/newIssue?project=ktor\n")
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
		case app.ProjectExtractError:
			fmt.Fprintf(os.Stderr, "Unable to extract downloaded archive to the directory %s.\n", projectDir)
		case app.JdkExtractError:
			if je, ok := e.Err.(interface{ Unwrap() []error }); ok && len(je.Unwrap()) > 0 {
				errs := je.Unwrap()
				var pe *os.PathError
				if errors.As(errs[0], &pe) {
					fmt.Fprintf(os.Stderr, "Unable to extract downloaded JDK to the directory %s.\n", pe.Path)
				}

				return
			}

			fmt.Fprintf(os.Stderr, "Unable to extract downloaded JDK.\n")
		case app.UnableLocateJdkError:
			var je jdk.Error
			errors.As(e.Err, &je)

			fmt.Fprintf(os.Stderr, "Unable to download JDK %s for %s %s\n", je.Descriptor.Version, je.Descriptor.Platform, je.Arch)
		case app.JdkServerError:
			fmt.Fprintf(os.Stderr, "Unexpected error occurred while connecting to a JDK server. Please try again later.\n")
		case app.JdkVerificationFailed:
			fmt.Fprintln(os.Stderr, "Checksum verification for the downloaded JDK failed")
		case app.GradlewChmodError:
			var pe *os.PathError
			errors.As(e.Err, &pe)
			fmt.Fprintf(os.Stderr, "Unable to make %s file executable.\n", pe.Path)
		case app.UnknownError:
			fmt.Fprintf(os.Stderr, "Unexpected error occurred.\n")
		case app.JdksDirError:
			var pe *os.PathError
			errors.As(e.Err, &pe)

			fmt.Fprintf(os.Stderr, "Unable to create a root directory %s to store downloaded JDKs.\n", pe.Path)
		default:
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

	fmt.Fprintln(os.Stderr)

	UsageTerminate(os.Stderr)
}

func PrintCommands(projectName string, javaHomeSet bool, jdkPath string) {
	initialProjectDir := projectName
	fmt.Print("To run the project use the following commands:\n\n")

	if runtime.GOOS == "windows" {
		fmt.Printf("cd %s\n", initialProjectDir)
		fmt.Println(".\\gradlew.bat run")

		if javaHomeSet {
			fmt.Println(".\\gradlew.bat run")
		} else {
			fmt.Printf("set \"JAVA_HOME=%s\" && .\\gradlew.bat run", jdkPath)
			fmt.Printf("You can also set the JAVA_HOME environment variable permanently or add this JDK in the IntelliJ IDEA\n")
		}
	} else {
		fmt.Printf("cd %s\n", initialProjectDir)

		if javaHomeSet {
			fmt.Println("./gradlew run")
		} else {
			fmt.Printf("JAVA_HOME=%s ./gradlew run\n\n", jdkPath)
			fmt.Printf("You can also set the JAVA_HOME environment variable permanently or add this JDK in the IntelliJ IDEA\n")
		}
	}
}
