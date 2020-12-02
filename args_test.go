package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandLineArgs_AddOptionArgs(t *testing.T) {
	commandLineArgs := NewCommandLineArgs()
	commandLineArgs.addOptionArgs("test-arg-1", "test-arg-value-1")
	commandLineArgs.addOptionArgs("test-arg-2", "test-arg-value-2")

	assert.Equal(t, 2, len(commandLineArgs.optionArgs))

	assert.Contains(t, commandLineArgs.optionArgs, "test-arg-1")
	assert.Contains(t, commandLineArgs.optionArgs, "test-arg-2")
	assert.NotContains(t, commandLineArgs.optionArgs, "test-arg-3")
}

func TestCommandLineArgs_GetOptionNames(t *testing.T) {
	commandLineArgs := NewCommandLineArgs()
	commandLineArgs.addOptionArgs("test-arg-1", "test-arg-value-1")
	commandLineArgs.addOptionArgs("test-arg-2", "test-arg-value-2")

	assert.Equal(t, 2, len(commandLineArgs.getOptionNames()))
	assert.Contains(t, commandLineArgs.getOptionNames(), "test-arg-1")
	assert.Contains(t, commandLineArgs.getOptionNames(), "test-arg-2")
}

func TestCommandLineArgs_ContainsOption(t *testing.T) {
	commandLineArgs := NewCommandLineArgs()
	commandLineArgs.addOptionArgs("test-arg-1", "test-arg-value-1")
	commandLineArgs.addOptionArgs("test-arg-2", "test-arg-value-2")

	assert.True(t, commandLineArgs.containsOption("test-arg-1"))
	assert.True(t, commandLineArgs.containsOption("test-arg-2"))
	assert.False(t, commandLineArgs.containsOption("test-arg-3"))
}

func TestCommandLineArgs_GetOptionValues(t *testing.T) {
	commandLineArgs := NewCommandLineArgs()
	commandLineArgs.addOptionArgs("test-arg-1", "test-arg-value-1")
	commandLineArgs.addOptionArgs("test-arg-2", "test-arg-value-2")
	commandLineArgs.addOptionArgs("test-arg-2", "test-arg-value-3")

	assert.Equal(t, 1, len(commandLineArgs.getOptionValues("test-arg-1")))
	assert.Equal(t, 2, len(commandLineArgs.getOptionValues("test-arg-2")))

	assert.Contains(t, commandLineArgs.getOptionValues("test-arg-1"), "test-arg-value-1")
	assert.Contains(t, commandLineArgs.getOptionValues("test-arg-2"), "test-arg-value-2")
	assert.Contains(t, commandLineArgs.getOptionValues("test-arg-2"), "test-arg-value-3")
}

func TestCommandLineArgs_AddNonOptionArgs(t *testing.T) {
	commandLineArgs := NewCommandLineArgs()
	commandLineArgs.addNonOptionArgs("test-nonoption-arg-1")
	commandLineArgs.addNonOptionArgs("test-nonoption-arg-2")

	assert.Equal(t, 2, len(commandLineArgs.nonOptionArgs))

	assert.Contains(t, commandLineArgs.nonOptionArgs, "test-nonoption-arg-1")
	assert.Contains(t, commandLineArgs.nonOptionArgs, "test-nonoption-arg-2")
}

func TestCommandLineArgs_GetNonOptionArgs(t *testing.T) {
	commandLineArgs := NewCommandLineArgs()
	commandLineArgs.addNonOptionArgs("test-nonoption-arg-1")
	commandLineArgs.addNonOptionArgs("test-nonoption-arg-2")

	assert.Equal(t, 2, len(commandLineArgs.getNonOptionArgs()))

	assert.Contains(t, commandLineArgs.getNonOptionArgs(), "test-nonoption-arg-1")
	assert.Contains(t, commandLineArgs.getNonOptionArgs(), "test-nonoption-arg-2")
}
