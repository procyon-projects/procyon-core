package core

import (
	"flag"
	"strings"
)

type CommandLineArgsParser interface {
	Parse(args []string) (CommandLineArgs, *CommandLineArgsParseError)
}

type SimpleCommandLineArgsParser struct {
}

func NewCommandLineArgsParser() SimpleCommandLineArgsParser {
	return SimpleCommandLineArgsParser{}
}

func (parser SimpleCommandLineArgsParser) Parse(args []string) (CommandLineArgs, *CommandLineArgsParseError) {
	cmdLineArgs := NewCommandLineArgs()

	appArgumentFlagSet := flag.NewFlagSet("ProcyonApplicationArguments", flag.ExitOnError)
	err := appArgumentFlagSet.Parse(args)
	if err != nil {
		return cmdLineArgs, NewCommandLineArgsParseError(err.Error())
	}

	for _, arg := range appArgumentFlagSet.Args() {
		if strings.HasPrefix(arg, "--") {
			optionText := arg[2:]
			indexOfEqualSign := strings.Index(optionText, "=")
			optionName := ""
			optionValue := ""
			if indexOfEqualSign > -1 {
				optionName = optionText[0:indexOfEqualSign]
				optionValue = optionText[indexOfEqualSign+1:]
			} else {
				optionName = optionText
			}
			optionName = strings.TrimSpace(optionName)
			optionValue = strings.TrimSpace(optionValue)
			if optionName == "" {
				return cmdLineArgs, NewCommandLineArgsParseError("Invalid argument syntax : " + arg)
			}
			cmdLineArgs.addOptionArgs(optionName, optionValue)
		} else {
			cmdLineArgs.addNonOptionArgs(arg)
		}
	}

	return cmdLineArgs, nil
}
