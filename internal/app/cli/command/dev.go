package command

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/project"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
)

func Dev(projectDir string, client *http.Client, verboseLogger *log.Logger, hasGlobalLog bool, homeDir string) {
	wrapper := "./gradlew"
	if runtime.GOOS == "windows" {
		wrapper = ".\\gradlew.bat"
	}

	wrapperPath := filepath.Join(projectDir, wrapper)

	if !utils.Exists(wrapperPath) {
		fmt.Printf(i18n.Get(i18n.GradleWrapperNotExistErr, wrapper, projectDir))
		os.Exit(1)
	}

	runTask, buildTask, guessed := project.GuessGradleTasks(projectDir)

	if !guessed {
		fmt.Printf(i18n.Get(i18n.KtorGradlePluginNotFound, project.DevModeSincePluginVersion, projectDir))
		os.Exit(1)
	}

	_, jdkPath, err := cli.ObtainJdk(client, verboseLogger, homeDir)

	if err != nil {
		cli.ExitWithError(err, hasGlobalLog, homeDir)
	}

	env := os.Environ()
	env = append(env, fmt.Sprintf("JAVA_HOME=%s", jdkPath))

	buildCmd := exec.Command(wrapper, buildTask, "--continuous")
	buildCmd.Env = env
	buildCmd.Dir = projectDir
	buildCmd.Stderr = os.Stderr

	verboseLogger.Printf(i18n.Get(i18n.StartingCommandMsg, "build", jdkPath, buildCmd.String()))

	err = buildCmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, i18n.Get(i18n.ErrorExecutingCommandMsg, "build", buildCmd))

		var pe *os.PathError
		if errors.As(err, &pe) {
			if errors.Is(pe.Err, os.ErrPermission) {
				err = &app.Error{Err: err, Kind: app.NoPermsForFile}
			}
		}

		cli.ExitWithError(err, hasGlobalLog, homeDir)
	}

	runCmd := exec.Command(wrapper, runTask, "-Pio.ktor.development=true")
	runCmd.Env = env
	runCmd.Dir = projectDir
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr

	verboseLogger.Printf(i18n.Get(i18n.StartingCommandMsg, "run", jdkPath, runCmd.String()))

	err = runCmd.Start()

	if err != nil {
		fmt.Fprintf(os.Stderr, i18n.Get(i18n.ErrorExecutingCommandMsg, "run", runCmd))

		var pe *os.PathError
		if errors.As(err, &pe) {
			if errors.Is(pe.Err, os.ErrPermission) {
				err = &app.Error{Err: err, Kind: app.NoPermsForFile}
			}
		}

		cli.ExitWithError(err, hasGlobalLog, homeDir)
	}

	doneChan := make(chan error)

	go func() {
		err = buildCmd.Wait()
		doneChan <- err
	}()

	go func() {
		err = runCmd.Wait()
		doneChan <- err
	}()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	go func() {
		for range interruptChan {
			// Send SIGINT to both child processes to exit them properly
			_ = buildCmd.Process.Signal(os.Interrupt)
			_ = runCmd.Process.Signal(os.Interrupt)
			os.Exit(1)
		}
	}()

	err = <-doneChan

	// One of the processes can be still alive
	_ = buildCmd.Process.Signal(os.Interrupt)
	_ = runCmd.Process.Signal(os.Interrupt)

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	} else if err != nil {
		os.Exit(1)
	}
}
