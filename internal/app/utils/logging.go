package utils

import (
	"errors"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"log"
	"os"
)

func PrepareGlobalLog(homeDir string) (*os.File, error) {
	err := os.Mkdir(config.KtorDir(homeDir), 0755)

	if err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}

	f, err := os.OpenFile(config.LogPath(homeDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		return nil, err
	}

	log.SetOutput(f)
	return f, nil
}
