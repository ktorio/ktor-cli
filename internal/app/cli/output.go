package cli

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"os"
	"runtime"
	"strings"
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
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.GenServerError))
		case app.GenServerTimeoutError:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.GenServerTimeoutError))
		case app.NetworkError:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.NetworkError))
		case app.InternalError:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.InternalError))
		case app.ProjectDirError:
			reportLog = false
			var pe *os.PathError
			errors.As(e.Err, &pe)

			switch {
			case errors.Is(pe.Err, os.ErrExist):
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.ProjectDirExist, pe.Path))
			case errors.Is(pe.Err, os.ErrPermission):
				fmt.Fprintf(os.Stderr, i18n.Get(i18n.NoPermsCreateProjectDir, pe.Path))
			}
		case app.ProjectExtractError:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.ProjectExtractError, projectDir))
		case app.JdkExtractError:
			if je, ok := e.Err.(interface{ Unwrap() []error }); ok && len(je.Unwrap()) > 0 {
				errs := je.Unwrap()
				var pe *os.PathError
				var appErr *app.Error
				if errors.As(errs[0], &pe) || (errors.As(errs[0], &appErr) && appErr.Kind == app.ExtractRootDirExistError && errors.As(appErr.Err, &pe)) {
					fmt.Fprintf(os.Stderr, i18n.Get(i18n.JdkExtractError, pe.Path))

					if errors.Is(pe.Err, os.ErrExist) {
						fmt.Fprintln(os.Stderr, i18n.Get(i18n.DirAlreadyExist))
					}
				}

				return
			}

			fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnableExtractJdk))
		case app.UnableLocateJdkError:
			var je jdk.Error
			errors.As(e.Err, &je)

			fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnableDownloadJdk, je.Descriptor.Version, je.Descriptor.Platform, je.Arch))
		case app.JdkServerError:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.JdkServerError))
		case app.JdkServerDownloadError:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.JdkServerDownloadError))
		case app.JdkVerificationFailed:
			fmt.Fprintln(os.Stderr, i18n.Get(i18n.ChecksumVerificationFailed))
		case app.GradlewChmodError:
			var pe *os.PathError
			errors.As(e.Err, &pe)
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnableMakeFileExec, pe.Path))
		case app.UnknownError:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnexpectedError))
		case app.JdksDirError:
			var pe *os.PathError
			errors.As(e.Err, &pe)

			fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnableCreateStoreJdkDir, pe.Path))
		default:
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnexpectedError))
		}
	}

	return
}

func HandleArgsValidation(err error) {
	var e *Error
	if !errors.As(err, &e) {
		fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnexpectedErrorWithArg, err))
	}

	switch e.Kind {
	case UnrecognizedFlagsError:
		var fe UnrecognizedFlags
		errors.As(e.Err, &fe)
		fmt.Fprintf(os.Stderr, i18n.Get(i18n.UnrecognizedFlagsError, strings.Join(fe, ", ")))
	case NoCommandError:
		fmt.Fprintln(os.Stderr, i18n.Get(i18n.NoCommandError))
	case CommandNotFoundError:
		var ce CommandError
		errors.As(e.Err, &ce)
		fmt.Fprintf(os.Stderr, i18n.Get(i18n.CommandNotFoundError, ce.Command))
	case WrongNumberOfArgumentsError:
		var ce CommandError
		errors.As(e.Err, &ce)

		if spec, ok := allCommandsSpec[ce.Command]; ok {
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.CommandArgumentsError, len(spec.args), formatArgs(spec.args), ce.Command))
		}
	default:
		// do nothing
	}

	fmt.Fprintln(os.Stderr)

	WriteUsage(os.Stderr)
	os.Exit(1)
}

func formatArgs(args map[string]Arg) string {
	sep := ""
	var sb strings.Builder
	for name, arg := range args {
		sb.WriteString(sep)

		if arg.required {
			sb.WriteString("<")
		} else {
			sb.WriteString("[")
		}

		sb.WriteString(name)

		if arg.required {
			sb.WriteString(">")
		} else {
			sb.WriteString("]")
		}

		sep = " "
	}

	return sb.String()
}

func PrintCommands(projectDir string, javaHomeSet bool, jdkPath string) {
	fmt.Print(i18n.Get(i18n.ToRunProject))

	if runtime.GOOS == "windows" {
		fmt.Printf("cd %s\n", projectDir)

		if javaHomeSet {
			fmt.Println(".\\gradlew.bat run")
		} else {
			fmt.Printf("cmd /C \"set JAVA_HOME=%s&& .\\gradlew.bat run\"\n\n", jdkPath)
			fmt.Printf(i18n.Get(i18n.JavaHomeJdkIdeaInstruction))
			fmt.Println(jdkPath)
		}
	} else {
		fmt.Printf("cd %s\n", projectDir)

		if javaHomeSet {
			fmt.Println("./gradlew run")
		} else {
			fmt.Printf("JAVA_HOME=%s ./gradlew run\n\n", jdkPath)
			fmt.Printf(i18n.Get(i18n.JavaHomeJdkIdeaInstruction))
			fmt.Println(jdkPath)
		}
	}
}
