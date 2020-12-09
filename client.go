package psychoclient

import (
	"net/http"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/proxy"
)

// ClientConfig : Client hello config
type ClientConfig struct {
	ClientHelloID   utls.ClientHelloID
	ClientHelloSpec utls.ClientHelloSpec
}

// NewClient : NewClient
func NewClient(config *ClientConfig, proxyURL ...string) (http.Client, error) {
	if len(proxyURL) > 0 && len(proxyURL) > 0 {
		dialer, err := newConnectDialer(proxyURL[0])
		if err != nil {
			return http.Client{}, err
		}
		return http.Client{
			Transport: newRoundTripper(config.ClientHelloID, config.ClientHelloSpec, dialer),
		}, nil
	} else {
		return http.Client{
			Transport: newRoundTripper(config.ClientHelloID, config.ClientHelloSpec, proxy.Direct),
		}, nil
	}
}
