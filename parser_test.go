package core

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

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
