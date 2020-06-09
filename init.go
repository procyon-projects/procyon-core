package core

func init() {
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService)
	/* Register Property Resolver */
	Register(NewSimplePropertyResolver)
}
