package model

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const maxFilenameLen = 255

type IdSet map[string]struct{}

type State struct {
	Running            bool
	ErrorLine          string
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
	ProjectName string
	ProjectDir  string
	Plugins     []string
	Quit        bool
}

func NewState() *State {
	return &State{
		Running:            true,
		PluginsByGroup:     make(map[string][]network.Plugin),
		IndirectPlugins:    make(map[string]IdSet),
		AddedPlugins:       make(IdSet),
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

func CheckProjectSettings(mdl *State) (errs []string) {
	if !IsDirEmptyOrAbsent(mdl.ProjectDir) {
		errs = append(errs, fmt.Sprintf(i18n.Get(i18n.DirNotEmptyError, mdl.ProjectDir)))
	}

	if ok, p := HasNonExistentDirsInPath(mdl.ProjectDir); ok {
		errs = append(errs, fmt.Sprintf(i18n.Get(i18n.DirNotExist, p)))
	}

	if len(filepath.Base(mdl.ProjectDir)) > maxFilenameLen {
		errs = append(errs, i18n.Get(i18n.ProjectDirLong))
	}

	for _, r := range mdl.ProjectName {
		isLatin := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
		if !isLatin && !unicode.IsDigit(r) && r != '.' && r != '_' && r != '-' {
			errs = append(errs, "Only Latin characters, digits, '_', '-' and '.' are allowed for the project name")
			break
		}
	}

	return
}

func CheckProjectSettingsAndUpdateError(mdl *State) {
	errs := CheckProjectSettings(mdl)
	mdl.ErrorLine = strings.Join(errs, "; ")
}

func InitProjectDir(mdl *State) {
	wd, err := os.Getwd()

	if err != nil {
		mdl.ProjectDir = filepath.Join(".", mdl.ProjectName)
		return
	}

	mdl.ProjectDir = filepath.Join(wd, mdl.ProjectName)
}
