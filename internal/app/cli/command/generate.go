package command

import (
	"context"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/generate"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"log"
	"net/http"
	"os"
)

func Generate(client *http.Client, projectDir, projectName string, plugins []string, verboseLogger *log.Logger, hasGlobalLog bool, ctx context.Context) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, i18n.Get(i18n.CannotDetermineHomeDir))
		os.Exit(1)
	}

	err = generate.Project(client, verboseLogger, projectDir, projectName, plugins, ctx)

	if err != nil {
		if _, err := os.Stat(projectDir); err == nil && utils.IsDirEmpty(projectDir) {
			_ = os.Remove(projectDir)
		}

		reportLog := cli.HandleAppError(projectDir, err)

		if hasGlobalLog && reportLog {
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.LogHint, config.LogPath(homeDir)))
		}

		if hasGlobalLog {
			log.Fatal(err)
		}

		os.Exit(1)
	}

	fmt.Printf(i18n.Get(i18n.ProjectCreated, projectName, projectDir))

	if jh, ok := jdk.JavaHome(); ok {
		if v, err := jdk.GetJavaMajorVersion(jh, homeDir); err == nil && v >= jdk.MinJavaVersion {
			fmt.Printf(i18n.Get(i18n.JDKDetectedJavaHome, jh))
			cli.PrintCommands(projectDir, true, "")
			os.Exit(0)
		}
	}

	if jdkPath, ok := config.GetValue("jdk"); ok {
		if st, err := os.Stat(jdkPath); err == nil && st.IsDir() {
			fmt.Printf(i18n.Get(i18n.JdkDetected, jdkPath))
			cli.PrintCommands(projectDir, false, jdkPath)
			os.Exit(0)
		}
	}

	if jdkPath, ok := jdk.FindLocally(jdk.MinJavaVersion); ok {
		config.SetValue("jdk", jdkPath)
		_ = config.Commit()
		fmt.Printf(i18n.Get(i18n.JdkFoundLocally, jdkPath))
		cli.PrintCommands(projectDir, false, jdkPath)
		os.Exit(0)
	}

	jdkPath, err := cli.DownloadJdk(homeDir, client, verboseLogger, 0)
	if err != nil {
		reportLog := cli.HandleAppError(projectDir, err)

		if hasGlobalLog && reportLog {
			fmt.Fprintf(os.Stderr, i18n.Get(i18n.LogHint, config.LogPath(homeDir)))
		}

		if hasGlobalLog {
			log.Fatal(err)
		}

		os.Exit(1)
	}

	fmt.Printf(i18n.Get(i18n.JdkDownloaded, jdkPath))
	cli.PrintCommands(projectDir, false, jdkPath)
}
