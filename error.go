package core

type CommandLineArgsParseError struct {
	message string
}

func NewCommandLineArgsParseError(errorMessage string) *CommandLineArgsParseError {
	return &CommandLineArgsParseError{
		message: errorMessage,
	}
}

func (e CommandLineArgsParseError) Error() string {
	return e.message
}
