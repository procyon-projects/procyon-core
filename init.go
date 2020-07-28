package core

func init() {
	/* Initialize Pool Manager */
	poolManager = newPoolManager()
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService)
}
