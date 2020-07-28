package core

var (
	proxyLoggerType = GetType((*ProxyLogger)(nil))
)

func init() {
	/* Initialize Pool Manager */
	poolManager = newPoolManager()
	/* Register Pool Types */
	RegisterPool(proxyLoggerType, newProxyLogger)
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService)
}
