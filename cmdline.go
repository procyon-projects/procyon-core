package core

import "strings"

const ProcyonApplicationCommandLinePropertySource = "ProcyonApplicationCommandLinePropertySource"
const NonOptionArgsPropertyName = "nonOptionArgs"

type CommandLinePropertySource interface {
	ContainsOption(name string) bool
	GetOptionValues(name string) []string
	GetNonOptionArgs() []string
}

type SimpleCommandLinePropertySource struct {
	name   string
	source CommandLineArgs
}

func NewSimpleCommandLinePropertySource(args []string) SimpleCommandLinePropertySource {
	cmdLineArgs, err := NewCommandLineArgsParser().Parse(args)
	if err != nil {
		panic(err)
	}
	cmdlinePropertySource := SimpleCommandLinePropertySource{
		name:   ProcyonApplicationCommandLinePropertySource,
		source: cmdLineArgs,
	}
	return cmdlinePropertySource
}

func SimpleCommandLinePropertySourceWithName(name string, args []string) SimpleCommandLinePropertySource {
	cmdLineArgs, err := NewCommandLineArgsParser().Parse(args)
	if err != nil {
		panic(err)
	}
	cmdlinePropertySource := SimpleCommandLinePropertySource{
		name:   name,
		source: cmdLineArgs,
	}
	return cmdlinePropertySource
}

func (cmdLineSource SimpleCommandLinePropertySource) GetName() string {
	return cmdLineSource.name
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
