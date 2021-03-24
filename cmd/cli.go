package main

import "github.com/urfave/cli/v2"

var (
	cmddemo1, cmddemo2, cmddemo3 cli.Command
)

func main() {
	cmddemo1 = cli.Command{
		Name:                   "",
		Aliases:                nil,
		Usage:                  "",
		UsageText:              "",
		Description:            "",
		ArgsUsage:              "",
		Category:               "",
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
