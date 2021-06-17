package core

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSimpleCommandLinePropertySource_ContainsOption(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	assert.True(t, commandLinePropertySource.ContainsOption("procyon.application.name"))
	assert.True(t, commandLinePropertySource.ContainsOption("procyon.server.port"))
}

func TestSimpleCommandLinePropertySource_ContainsProperty(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	assert.True(t, commandLinePropertySource.ContainsProperty("procyon.application.name"))
	assert.True(t, commandLinePropertySource.ContainsProperty("procyon.server.port"))
}

func TestSimpleCommandLinePropertySource_GetName(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	assert.Equal(t, ProcyonApplicationCommandLinePropertySource, commandLinePropertySource.GetName())
}

func TestSimpleCommandLinePropertySource_GetNonOptionArgs(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	nonOptionArgs := commandLinePropertySource.GetNonOptionArgs()
	assert.Contains(t, nonOptionArgs, "-debug")
}

func TestSimpleCommandLinePropertySource_GetProperty(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	result := commandLinePropertySource.GetProperty("procyon.application.name")
	assert.Equal(t, "\"Test Application\"", result)

	result = commandLinePropertySource.GetProperty("procyon.server.port")
	assert.Equal(t, "8080,8090", result)
}

func TestSimpleCommandLinePropertySource_GetPropertyNames(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertyNames := commandLinePropertySource.GetPropertyNames()
	assert.Contains(t, propertyNames, "procyon.application.name")
	assert.Contains(t, propertyNames, "procyon.server.port")
}

func TestSimpleCommandLinePropertySource_GetOptionValues(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	values := commandLinePropertySource.GetOptionValues("procyon.server.port")
	assert.Contains(t, values, "8080")
	assert.Contains(t, values, "8090")
}

func TestSimpleCommandLinePropertySource_GetSource(t *testing.T) {
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	assert.NotNil(t, commandLinePropertySource.GetSource())
}

func getTestApplicationArguments() []string {
	var args = make([]string, 0)
	args = append(args, os.Args...)
	args = append(args, "--procyon.application.name=\"Test Application\"")
	args = append(args, "--procyon.server.port=8080")
	args = append(args, "--procyon.server.port=8090")
	args = append(args, "-debug")
	return args
}

func TestSimpleCommandLineArgsParser_Parse(t *testing.T) {
	commandLineParser := NewCommandLineArgsParser()

	args, err := commandLineParser.Parse(getTestApplicationArguments())

	assert.Nil(t, err)
	assert.NotNil(t, args)

	assert.Equal(t, 2, len(args.optionArgs))

	assert.Equal(t, 1, len(args.getOptionValues("procyon.application.name")))
	assert.Equal(t, 2, len(args.getOptionValues("procyon.server.port")))

	assert.Contains(t, args.getOptionNames(), "procyon.application.name")
	assert.Contains(t, args.getOptionNames(), "procyon.server.port")

	assert.True(t, args.containsOption("procyon.application.name"))
	assert.True(t, args.containsOption("procyon.server.port"))

	assert.Contains(t, args.getOptionValues("procyon.application.name"), "\"Test Application\"")
	assert.Contains(t, args.getOptionValues("procyon.server.port"), "8080")
	assert.Contains(t, args.getOptionValues("procyon.server.port"), "8090")
}
