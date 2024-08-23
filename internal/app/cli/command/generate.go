package command

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/generate"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"log"
	"net/http"
	"os"
)

func Generate(client *http.Client, projectDir, projectName string, verboseLogger *log.Logger, hasGlobalLog bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to determine home directory.")
		os.Exit(1)
	}

	err = generate.Project(client, verboseLogger, projectDir, projectName)

	if err != nil {
		if _, err := os.Stat(projectDir); err == nil && utils.IsDirEmpty(projectDir) {
			_ = os.Remove(projectDir)
		}

		reportLog := cli.HandleAppError(projectDir, err)

		if hasGlobalLog && reportLog {
			fmt.Fprintf(os.Stderr, "You can find more information in the log: %s\n", config.LogPath(homeDir))
		}

		if hasGlobalLog {
			log.Fatal(err)
		}

		os.Exit(1)
	}

	fmt.Printf("Project \"%s\" has been created in the directory %s.\n", projectName, projectDir)

	if jh, ok := jdk.JavaHome(); ok {
		if v, err := jdk.GetJavaMajorVersion(jh); err == nil && v >= jdk.MinJavaVersion {
			fmt.Printf("JDK is detected in JAVA_HOME=%s\n", jh)
			cli.PrintCommands(projectName, true, "")
			os.Exit(0)
		}
	}

	if jdkPath, ok := config.GetValue("jdk"); ok {
		if st, err := os.Stat(jdkPath); err == nil && st.IsDir() {
			fmt.Printf("Detected JDK %s\n", jdkPath)
			cli.PrintCommands(projectName, false, jdkPath)
			os.Exit(0)
		}
	}

	if jdkPath, ok := jdk.FindLocally(jdk.MinJavaVersion); ok {
		config.SetValue("jdk", jdkPath)
		_ = config.Commit()
		fmt.Printf("JDK found locally %s\n", jdkPath)
		cli.PrintCommands(projectName, false, jdkPath)
		os.Exit(0)
	}

	jdkPath, err := cli.DownloadJdk(homeDir, client, verboseLogger, 0)
	if err != nil {
		reportLog := cli.HandleAppError(projectDir, err)

		if hasGlobalLog && reportLog {
			fmt.Fprintf(os.Stderr, "You can find more information in the log: %s\n", config.LogPath(homeDir))
		}

		if hasGlobalLog {
			log.Fatal(err)
		}

		os.Exit(1)
	}

	fmt.Printf("JDK has been downloaded to %s\n", jdkPath)
	cli.PrintCommands(projectName, false, jdkPath)
}
