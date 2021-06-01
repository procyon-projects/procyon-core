package core

import (
	"errors"
	"flag"
	"strings"
)

const ProcyonApplicationCommandLinePropertySource = "ProcyonApplicationCommandLinePropertySource"
const NonOptionArgsPropertyName = "nonOptionArgs"

type CommandLinePropertySource interface {
	PropertySource
	ContainsOption(name string) bool
	GetOptionValues(name string) []string
	GetNonOptionArgs() []string
}

type SimpleCommandLinePropertySource struct {
	source CommandLineArgs
}

func NewSimpleCommandLinePropertySource(args []string) SimpleCommandLinePropertySource {
	cmdLineArgs, err := NewCommandLineArgsParser().Parse(args)

	if err != nil {
		panic(err)
	}

	cmdlinePropertySource := SimpleCommandLinePropertySource{
		source: cmdLineArgs,
	}

	return cmdlinePropertySource
}

func (cmdLineSource SimpleCommandLinePropertySource) GetName() string {
	return ProcyonApplicationCommandLinePropertySource
}

func (cmdLineSource SimpleCommandLinePropertySource) GetSource() interface{} {
	return cmdLineSource.source
}

func (cmdLineSource SimpleCommandLinePropertySource) GetProperty(name string) interface{} {
	if NonOptionArgsPropertyName == name {
		nonOptValues := cmdLineSource.GetNonOptionArgs()

		if nonOptValues != nil {
			return strings.Join(nonOptValues, ",")
		}

		return nil
	}

	optValues := cmdLineSource.GetOptionValues(name)

	if optValues != nil {
		return strings.Join(optValues, ",")
	}

	return nil
}

func (cmdLineSource SimpleCommandLinePropertySource) ContainsProperty(name string) bool {
	if NonOptionArgsPropertyName == name {
		return len(cmdLineSource.GetNonOptionArgs()) != 0
	}

	return cmdLineSource.ContainsOption(name)
}

func (cmdLineSource SimpleCommandLinePropertySource) ContainsOption(name string) bool {
	return cmdLineSource.source.containsOption(name)
}

func (cmdLineSource SimpleCommandLinePropertySource) GetOptionValues(name string) []string {
	return cmdLineSource.source.getOptionValues(name)
}

func (cmdLineSource SimpleCommandLinePropertySource) GetNonOptionArgs() []string {
	return cmdLineSource.source.getNonOptionArgs()
}

func (cmdLineSource SimpleCommandLinePropertySource) GetPropertyNames() []string {
	return cmdLineSource.source.getOptionNames()
}

type SimpleCommandLineArgsParser struct {
}

func NewCommandLineArgsParser() SimpleCommandLineArgsParser {
	return SimpleCommandLineArgsParser{}
}

func (parser SimpleCommandLineArgsParser) Parse(args []string) (CommandLineArgs, error) {
	cmdLineArgs := NewCommandLineArgs()
	appArgumentFlagSet := flag.NewFlagSet("ProcyonApplicationArguments", flag.ContinueOnError)

	err := appArgumentFlagSet.Parse(args)

	if err != nil {
		return cmdLineArgs, err
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
				return cmdLineArgs, errors.New("Invalid argument syntax : " + arg)
			}

			cmdLineArgs.addOptionArgs(optionName, optionValue)
		} else {
			cmdLineArgs.addNonOptionArgs(arg)
		}

	}

	return cmdLineArgs, nil
}
