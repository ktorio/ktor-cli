package model

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"os"
	"path/filepath"
	"unicode"
)

const maxFilenameLen = 255

type IdSet map[string]struct{}

type State struct {
	Running            bool
	errorMap           map[ErrorKind]string
	StatusLine         string
	Search             string
	PluginsFetched     bool
	Groups             []string
	PluginsByGroup     map[string][]network.Plugin
	AddedPlugins       IdSet
	IndirectPlugins    map[string]IdSet
	PluginDeps         map[string][]string
	AllPluginsByGroup  map[string][]network.Plugin
	AllSortedGroups    []string
	ShouldFetchPlugins bool
	Result
}

type Result struct {
	ProjectName    string
	ProjectDir     string
	Plugins        []string
	VersionCatalog bool
	Quit           bool
}

func NewState() *State {
	return &State{
		Running:            true,
		PluginsByGroup:     make(map[string][]network.Plugin),
		IndirectPlugins:    make(map[string]IdSet),
		AddedPlugins:       make(IdSet),
		errorMap:           make(map[ErrorKind]string),
		ShouldFetchPlugins: true,
	}
}

func InsertRune(input string, pos int, r rune) string {
	if pos < 0 {
		return input
	}

	if input == "" {
		return fmt.Sprintf("%c", r)
	}

	runes := []rune(input)
	if pos >= len(runes) {
		return string(append([]rune(input), r))
	}

	var result []rune
	for _, run := range runes[:pos] {
		result = append(result, run)
	}
	result = append(result, r)
	for _, run := range runes[pos:] {
		result = append(result, run)
	}

	return string(result)
}

func DeleteChar(input string, pos int) string {
	if pos >= len(input) || pos < 0 {
		return input
	}

	runes := []rune(input)

	var result []rune
	for _, r := range runes[:pos] {
		result = append(result, r)
	}

	for _, r := range runes[pos+1:] {
		result = append(result, r)
	}

	return string(result)
}

func CheckProjectSettings(mdl *State) bool {
	mdl.RemoveErrors(ProjectDirNotEmptyError, DirNotExistError, ProjectDirTooLongError, ProjectNameEmptyError, ProjectNameAllowedCharsError)
	hasError := false

	if len(mdl.ProjectName) == 0 {
		hasError = true
		mdl.SetError(ProjectNameEmptyError, i18n.Get(i18n.ProjectNameRequired))
	}

	if !IsDirEmptyOrAbsent(mdl.GetProjectPath()) {
		hasError = true
		mdl.SetError(ProjectDirNotEmptyError, fmt.Sprintf(i18n.Get(i18n.DirNotEmptyError, mdl.GetProjectPath())))
	}

	if ok, p := HasNonExistentDirsInPath(mdl.GetProjectPath()); ok {
		hasError = true
		mdl.SetError(DirNotExistError, fmt.Sprintf(i18n.Get(i18n.DirNotExist, p)))
	}

	if len(filepath.Base(mdl.GetProjectPath())) > maxFilenameLen {
		hasError = true
		mdl.SetError(ProjectDirTooLongError, i18n.Get(i18n.ProjectDirLong))
	}

	for _, r := range mdl.ProjectName {
		isLatin := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
		if !isLatin && !unicode.IsDigit(r) && r != '.' && r != '_' && r != '-' {
			hasError = true
			mdl.SetError(ProjectNameAllowedCharsError, i18n.Get(i18n.ProjectNameAllowedChars))
			break
		}
	}

	return hasError
}

func InitProjectDir(mdl *State) {
	wd, err := os.Getwd()

	if err != nil {
		mdl.ProjectDir = "."
		return
	}

	mdl.ProjectDir = wd
}

func (mdl *State) GetProjectPath() string {
	return filepath.Join(mdl.ProjectDir, mdl.ProjectName)
}

func FindVacantProjectName(mdl *State) (string, error) {
	for i := 1; i < 1000; i++ {
		newName := fmt.Sprintf("%s%d", mdl.ProjectName, i)
		if _, err := os.Stat(filepath.Join(mdl.ProjectDir, newName)); errors.Is(err, os.ErrNotExist) {
			return newName, nil
		}
	}

	return "", &app.Error{Err: errors.New("cannot find vacant project name"), Kind: app.UnknownError}
}
