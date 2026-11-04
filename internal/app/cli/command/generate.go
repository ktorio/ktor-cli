package command

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/generate"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
)

func Generate(client *http.Client, projectDir, projectName string, plugins []string, versionCatalog bool, verboseLogger *log.Logger, hasGlobalLog bool, ctx context.Context) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, i18n.Get(i18n.CannotDetermineHomeDir))
		os.Exit(1)
	}

	err = generate.Project(client, verboseLogger, projectDir, projectName, plugins, versionCatalog, ctx)

	if err != nil {
		cli.ExitWithProjectError(err, projectDir, hasGlobalLog, homeDir)
	}

	fmt.Print(i18n.Get(i18n.ProjectCreated, projectName, projectDir))

	jdkSrc, jdkPath, err := cli.ObtainJdk(client, verboseLogger, homeDir)

	switch jdkSrc {
	case jdk.FromJavaHome:
		fmt.Print(i18n.Get(i18n.JDKDetectedJavaHome, jdkPath))
		cli.PrintCommands(projectDir, true, "")
	case jdk.FromConfig:
		fmt.Print(i18n.Get(i18n.JdkDetected, jdkPath))
		cli.PrintCommands(projectDir, false, jdkPath)
	case jdk.Locally:
		fmt.Print(i18n.Get(i18n.JdkFoundLocally, jdkPath))
		cli.PrintCommands(projectDir, false, jdkPath)
	case jdk.Downloaded:
		if err != nil {
			cli.ExitWithProjectError(err, projectDir, hasGlobalLog, homeDir)
		}
		fmt.Print(i18n.Get(i18n.JdkDownloaded, jdkPath))
		cli.PrintCommands(projectDir, false, jdkPath)
	}
}
