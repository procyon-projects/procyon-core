package core

func init() {
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService())
}
