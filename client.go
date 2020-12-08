package psychoclient

import (
	"net/http"

	"golang.org/x/net/proxy"

	utls "github.com/refraction-networking/utls"
)

type clientConfig struct {
	clientHello     utls.ClientHelloID
	clientHelloSpec utls.ClientHelloSpec
	proxyURL        []string
}

// NewClient : NewClient
func NewClient(config *clientConfig) (http.Client, error) {
	if len(config.proxyURL) > 0 && len(config.proxyURL) > 0 {
		dialer, err := newConnectDialer(config.proxyURL[0])
		if err != nil {
			return http.Client{}, err
		}
		return http.Client{
			Transport: newRoundTripper(config.clientHello, config.clientHelloSpec, dialer),
		}, nil
	} else {
		return http.Client{
			Transport: newRoundTripper(config.clientHello, config.clientHelloSpec, proxy.Direct),
		}, nil
	}
}
