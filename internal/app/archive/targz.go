package archive

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ExtractTarGz(r io.Reader, outDir string, logger *log.Logger) (rootDirs utils.StringSet, err error) {
	zr, err := gzip.NewReader(r)
	rootDirs = utils.NewStringSet()

	if err != nil {
		return rootDirs, err
	}

	defer zr.Close()

	tr := tar.NewReader(zr)

	var extractErrors []error
	for {
		header, err := tr.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			extractErrors = append(extractErrors, err)
			continue
		}

		err = func() error {
			switch header.Typeflag {
			case tar.TypeDir:
				isRootDir := strings.Count(header.Name, "/") == 1
				if isRootDir {
					rootDirs.Add(filepath.Join(outDir, header.Name))
				}
				logger.Printf(i18n.Get(i18n.CreatingDir, path.Join(outDir, header.Name)))
				if err := os.Mkdir(path.Join(outDir, header.Name), os.FileMode(header.Mode&0xfff)); err != nil {
					if os.IsExist(err) && isRootDir {
						return &app.Error{Err: err, Kind: app.ExtractRootDirExistError}
					}

					return err
				}
			case tar.TypeReg:
				fp := path.Join(outDir, header.Name)
				outFile, err := os.Create(fp)

				if err != nil {
					return err
				}

				defer outFile.Close()
				logger.Printf(i18n.Get(i18n.Extracting, header.Name, fp))
				if _, err := io.Copy(outFile, tr); err != nil {
					return err
				}

				err = outFile.Sync()
				if err != nil {
					return err
				}

				err = os.Chmod(fp, os.FileMode(header.Mode&0xfff))
				if err != nil {
					return err
				}
			}

			return nil
		}()

		if err != nil {
			extractErrors = append(extractErrors, err)
		}

		var appError *app.Error
		if errors.As(err, &appError) && appError.Kind == app.ExtractRootDirExistError {
			break
		}
	}

	err = errors.Join(extractErrors...)
	return
}
