package cli

type Command string

const (
	NewCommand Command = "new"
)

// A map of a command to a number of args
var allCommands = map[Command]int{
	NewCommand: 1,
}
