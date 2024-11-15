package jdk

import (
	"bytes"
	"errors"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/archive"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/progress"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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
		logger.Printf(i18n.Get(i18n.ExtractingJdkFiles, outDir))

		if d.Platform == "windows" {
			reader, progressBar := progress.NewReaderAt(
				bytes.NewReader(jdkBytes),
				i18n.Get(i18n.ExtractingJdkProgress),
				len(jdkBytes),
				logger.Writer() == io.Discard,
			)
			defer progressBar.Done()

			extractedDirs, extractErr = archive.ExtractZip(reader, int64(len(jdkBytes)), outDir, logger)
		} else {
			reader, progressBar := progress.NewReader(
				bytes.NewReader(jdkBytes),
				i18n.Get(i18n.ExtractingJdkProgress),
				len(jdkBytes),
				logger.Writer() == io.Discard,
			)
			defer progressBar.Done()

			extractedDirs, extractErr = archive.ExtractTarGz(reader, outDir, logger)
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

	if runtime.GOOS == "darwin" {
		extractDir = filepath.Join(extractDir, "Contents", "Home")
	}

	return extractDir, nil
}
