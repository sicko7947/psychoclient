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
	return http.Client{
		Transport: newRoundTripper(config.ClientHelloID, config.ClientHelloSpec, proxy.Direct),
	}, nil
}
