package psychoclient

import (
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

func noRedirects(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func newDefaultClient(
	timeout time.Duration,
	redirectCallback func(req *http.Request, via []*http.Request) error,
	proxyURL ...string) (http.Client, error) {

	var proxy *url.URL
	if len(proxyURL) > 0 && len(proxyURL[0]) > 0 {
		proxy, _ = url.Parse(proxyURL[0])
	} else {
		return http.Client{
			Transport:     &http.Transport{},
			CheckRedirect: redirectCallback,
			Timeout:       timeout,
		}, nil
	}

	return http.Client{
		Timeout:       timeout,
		CheckRedirect: redirectCallback,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}, nil
}

// newClient : creates a new base psycho http client
func newClient(useDefaultClient bool, followRedirects bool, timeout time.Duration, proxyURL ...string) (http.Client, error) {

	redirectCallback := noRedirects
	if followRedirects {
		redirectCallback = nil // Reset to default policy (redirect up to 10 times)
	}

	if useDefaultClient {
		return newDefaultClient(timeout, redirectCallback, proxyURL...)
	}

	var dialer proxy.ContextDialer = proxy.Direct
	if len(proxyURL) > 0 && len(proxyURL[0]) > 0 {
		d, err := newConnectDialer(proxyURL[0])
		if err != nil {
			return http.Client{
				CheckRedirect: redirectCallback,
				Transport: &UHTTPTransport{
					DialContext:          dialer.DialContext,
					UTLSClientHelloSpecs: getChromeClientHelloSpecs(),
				},
				Timeout: timeout,
			}, err
		}
		dialer = d
	}

	return http.Client{
		CheckRedirect: redirectCallback,
		Transport: &UHTTPTransport{
			DialContext:          dialer.DialContext,
			UTLSClientHelloSpecs: getChromeClientHelloSpecs(),
		},
	}, nil
}
