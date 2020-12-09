package psychoclient

import (
	"net/http"

	"golang.org/x/net/proxy"

	utls "github.com/refraction-networking/utls"
)

type ClientConfig struct {
	clientHelloID   utls.ClientHelloID
	clientHelloSpec utls.ClientHelloSpec
	proxyURL        []string
}

// NewClient : NewClient
func NewClient(config *ClientConfig) (http.Client, error) {
	if len(config.proxyURL) > 0 && len(config.proxyURL) > 0 {
		dialer, err := newConnectDialer(config.proxyURL[0])
		if err != nil {
			return http.Client{}, err
		}
		return http.Client{
			Transport: newRoundTripper(config.clientHelloID, config.clientHelloSpec, dialer),
		}, nil
	} else {
		return http.Client{
			Transport: newRoundTripper(config.clientHelloID, config.clientHelloSpec, proxy.Direct),
		}, nil
	}
}
