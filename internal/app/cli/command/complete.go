package command

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"strings"
)

func Complete(allModules []string, shell string) (string, error) {
	switch shell {
	case "zsh":
		err, s := completeZsh(allModules)
		return err, s
	case "bash":
		err, s := completeBash(allModules)
		return err, s
	case "fish":
		err, s := completeFish(allModules)
		return err, s
	default:
		return "", &app.Error{Err: &cli.ShellError{Shell: shell}, Kind: app.UnrecognizedShellError}
	}
}

func completeFish(allModules []string) (string, error) {
	var commands []string
	for c, s := range cli.AllCommandsSpec {
		if c == cli.CompletionCommand {
			continue
		}

		commands = append(commands, fmt.Sprintf("complete -c ktor -n \"__fish_use_subcommand\" -f -a \"%s\" -d \"%s\"", c, s.Description))
	}

	newCommand := fmt.Sprintf("complete -c ktor -n \"__fish_seen_subcommand_from new\" -a \"(commandline -ot)\" -d \"[project-name]\"")

	var modules []string
	for _, m := range allModules {
		modules = append(modules, fmt.Sprintf("complete -c ktor -n \"__fish_seen_subcommand_from add\" -f -a \"%s\"", m))
	}

	s := fmt.Sprintf(`
%s
%s
%s
`, strings.Join(commands, "\n"), newCommand, strings.Join(modules, "\n"))

	return s, nil
}

func completeBash(allModules []string) (string, error) {
	var commands []string
	for c := range cli.AllCommandsSpec {
		if c == cli.CompletionCommand {
			continue
		}

		commands = append(commands, string(c))
	}

	for _, s := range cli.AllFlagsSpec {
		for _, a := range s.Aliases {
			commands = append(commands, a)
		}
	}

	var modules []string
	for _, m := range allModules {
		modules = append(modules, m)
	}

	s := fmt.Sprintf(`
_ktor_complete() {
    local cur prev subcommands modules
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    subcommands="%s"
    modules="%s"

    if [[ $COMP_CWORD -eq 1 ]]; then
        COMPREPLY=($(compgen -W "$subcommands" -- "$cur"))
    else
        case "${COMP_WORDS[1]}" in
            new)
                if [[ $COMP_CWORD -eq 2 ]]; then
                    COMPREPLY=()
                fi
                ;;
            add)
                COMPREPLY=($(compgen -W "$modules" -- "$cur"))
                ;;
            *)
                COMPREPLY=()
                ;;
        esac
    fi
}

complete -F _ktor_complete ktor
`, strings.Join(commands, " "), strings.Join(modules, " "))

	return s, nil
}

func completeZsh(allModules []string) (string, error) {
	var commands []string
	for c, s := range cli.AllCommandsSpec {
		if c == cli.CompletionCommand {
			continue
		}

		commands = append(commands, fmt.Sprintf("\"%s:%s\"", string(c), s.Description))
	}

	for _, s := range cli.AllFlagsSpec {
		for _, a := range s.Aliases {
			commands = append(commands, fmt.Sprintf("\"%s:%s\"", a, s.Description))
		}
	}

	var modules []string
	for _, m := range allModules {
		modules = append(modules, `"`+m+`"`)
	}

	s := fmt.Sprintf(`
_ktor_subcommands=(
%s
)

_ktor_add_modules=(
%s
)

# Completion function for the ktor command
_ktor() {
  local context state line
  _arguments -C \
    '1: :->subcommands' \
    '*::arg:->args'

  case $state in
    subcommands)
      _describe -t commands "Ktor commands" _ktor_subcommands
      ;;
    args)
      case ${line[1]} in
        new)
          _arguments '*::[project-name]:_files'
          ;;
        add)
          _describe -t modules "Ktor modules" _ktor_add_modules
          ;;
      esac
  esac
}

# Load the completion function
compdef _ktor ktor
`, strings.Join(commands, "\n"), strings.Join(modules, "\n"))

	return s, nil
}
