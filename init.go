package core

func init() {
	/* Init Proxy Logger */
	initProxyLoggerPool()
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService)
}
