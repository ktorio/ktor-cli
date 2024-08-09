package archive

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExtractZip(archive []byte, outDir string, logger *log.Logger) error {
	zr, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))

	if err != nil {
		return err
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

			if filepath.Dir(outPath) != outDir {
				logger.Printf("Creating directory %s\n", filepath.Dir(outPath))
				err := os.MkdirAll(filepath.Dir(outPath), 0755)

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

	return errors.Join(zipErrors...)
}
