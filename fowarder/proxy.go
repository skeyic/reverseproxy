package fowarder

type ProxyForwarder struct {
	host string
}

func NewProxyForwarder(host string) *ProxyForwarder {
	return &ProxyForwarder{host: host}
}
