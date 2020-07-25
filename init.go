package core

func init() {
	/* Initialize Pool Manager */
	PoolManager = newPoolManager()
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService)
}
