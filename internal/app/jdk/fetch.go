package jdk

import (
	"bytes"
	"errors"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/archive"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"log"
	"net/http"
	"os"
	"sync"
)

func fetch(client *http.Client, d *Descriptor, outDir string, logger *log.Logger) (string, error) {
	jdkBytes, err := DownloadJdk(client, d, logger)
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup

	var jdkValid bool
	var verifyErr error
	wg.Add(1)
	go func() {
		jdkValid, verifyErr = Verify(client, d, bytes.NewReader(jdkBytes), logger)
		wg.Done()
	}()

	wg.Add(1)
	var extractErr error
	var extractDir string
	go func() {
		defer wg.Done()
		extractedDirs := utils.NewStringSet()
		logger.Printf("Extracting JDK files to %s\n", outDir)
		if d.Platform == "windows" {
			extractedDirs, extractErr = archive.ExtractZip(bytes.NewReader(jdkBytes), int64(len(jdkBytes)), outDir, logger)
		} else {
			extractedDirs, extractErr = archive.ExtractTarGz(bytes.NewReader(jdkBytes), outDir, logger)
		}

		if extractErr != nil {
			return
		}

		dir, extractErr := extractedDirs.Single()

		if extractErr != nil {
			return
		}

		extractDir = dir
	}()

	wg.Wait()
	if verifyErr != nil {
		if extractErr == nil {
			os.RemoveAll(extractDir)
		}

		return "", &app.Error{Err: verifyErr, Kind: app.JdkVerificationFailed}
	}

	if extractErr != nil {
		return "", &app.Error{Err: extractErr, Kind: app.JdkExtractError}
	}

	if !jdkValid {
		if extractErr == nil {
			os.RemoveAll(extractDir)
		}

		return "", &app.Error{Err: errors.New("verify jdk: downloaded JDK checksum failed"), Kind: app.JdkVerificationFailed}
	}

	return extractDir, nil
}
