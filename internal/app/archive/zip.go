package archive

import (
	"archive/zip"
	"errors"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ExtractZip(rt io.ReaderAt, size int64, outDir string, logger *log.Logger) (rootDirs utils.StringSet, err error) {
	rootDirs = utils.NewStringSet()
	zr, err := zip.NewReader(rt, size)

	if err != nil {
		return
	}

	var zipErrors []error
	for _, zf := range zr.File {
		zipFile, err := zr.Open(zf.Name)

		if err != nil {
			zipErrors = append(zipErrors, err)
			continue
		}

		err = func() error {
			defer zipFile.Close()

			outPath := filepath.Join(outDir, zf.Name)

			if filepath.Dir(zf.Name) != "." {
				dir := filepath.Dir(outPath)
				logger.Printf("Creating directory %s\n", dir)

				if i := strings.Index(zf.Name, "/"); i != -1 {
					rootDirs.Add(filepath.Join(outDir, zf.Name[:i]))
				}

				err := os.MkdirAll(dir, 0755)

				if err != nil {
					return err
				}
			}

			out, err := os.Create(outPath)

			if err != nil {
				return err
			}

			logger.Printf("Extracting %s to %s\n", zf.Name, outPath)
			if _, err = io.Copy(out, zipFile); err != nil {
				return err
			}

			return out.Sync()
		}()

		if err != nil {
			zipErrors = append(zipErrors, err)
		}
	}

	err = errors.Join(zipErrors...)
	return
}
