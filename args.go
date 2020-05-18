package core

import (
	"reflect"
)

type CommandLineArgs struct {
	optionArgs    map[string][]string
	nonOptionArgs []string
}

func NewCommandLineArgs() CommandLineArgs {
	return CommandLineArgs{
		optionArgs:    make(map[string][]string),
		nonOptionArgs: make([]string, 0),
	}
}

func (args *CommandLineArgs) addOptionArgs(name string, value string) {
	if args.optionArgs[name] == nil {
		args.optionArgs[name] = make([]string, 0)
	}
	args.optionArgs[name] = append(args.optionArgs[name], value)
}

func (args CommandLineArgs) getOptionNames() []string {
	argMapKeys := reflect.ValueOf(args.optionArgs).MapKeys()
	optionNames := make([]string, len(argMapKeys))
	for i := 0; i < len(argMapKeys); i++ {
		optionNames[i] = argMapKeys[i].String()
	}
	return optionNames
}

func (args CommandLineArgs) containsOption(name string) bool {
	return args.optionArgs[name] != nil
}

func (args CommandLineArgs) getOptionValues(name string) []string {
	return args.optionArgs[name]
}

func (args *CommandLineArgs) addNonOptionArgs(value string) {
	args.nonOptionArgs = append(args.nonOptionArgs, value)
}

func (args CommandLineArgs) getNonOptionArgs() []string {
	return args.nonOptionArgs
}
