package psychoclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var defaultTimeout time.Duration = 10 * time.Second

type Client struct {
	client http.Client
	err    error
}

func NewClient(s *SessionBuilder) *Client {
	client, err := newClient(s.UseDefaultClient, s.FollowRedirects, defaultTimeout, s.Proxy)
	return &Client{client: client, err: err}
}

func (c *Client) DoNewRequest(b *RequestBuilder) (res *http.Response, respBody []byte, err error) {
	defer c.close()

	if c.err != nil {
		return nil, nil, err
	}

	// setup a new http request, setup request method & endpoint & payload'
	req, _ := http.NewRequest(b.Method, b.Endpoint, b.Payload)

	if len(b.Headers) > 0 { // set the request headers
		for key, value := range b.Headers {
			req.Header.Add(key, value)
		}
	}

	// starts a new http request
	channel := make(chan *response, 1)
	defer close(channel)

	go func() {
		res, err := c.client.Do(req)
		if err != nil || res == nil { // checking if empty response or err occured during the request
			channel <- &response{
				httpResponse: nil,
				err:          err,
			}
			return
		}

		channel <- &response{ // puts the response back to channel
			httpResponse: res,
			err:          nil,
		}
	}()

	response := <-channel // waiting for channel to receive response
	res = response.httpResponse
	err = response.err

	if res != nil {
		// set response body
		if res.Body != nil {
			defer response.httpResponse.Body.Close()
			body, e := io.ReadAll(res.Body)
			respBody = body
			if e != nil {
				err = fmt.Errorf("error reading response body: %v", e)
			}
		}
	}

	return res, respBody, err
}

func (c *Client) RoundTripNewRequest(b *RequestBuilder) (res *http.Response, respBody []byte, err error) {
	defer c.close()

	if c.err != nil {
		return nil, nil, err
	}

	// setup a new http request, setup request method & endpoint & payload'
	req, _ := http.NewRequest(b.Method, b.Endpoint, b.Payload)

	if len(b.Headers) > 0 { // set the request headers
		for key, value := range b.Headers {
			req.Header.Add(key, value)
		}
	}

	// starts a new http request
	channel := make(chan *response, 1)
	defer close(channel)

	go func() {
		res, err := c.client.Transport.RoundTrip(req)
		if err != nil || res == nil { // checking if empty response or err occured during the request
			channel <- &response{
				httpResponse: nil,
				err:          err,
			}
			return
		}

		channel <- &response{ // puts the response back to channel
			httpResponse: res,
			err:          nil,
		}
	}()

	response := <-channel // waiting for channel to receive response

	res = response.httpResponse
	err = response.err

	if res != nil {
		// set response body
		if res.Body != nil {
			defer response.httpResponse.Body.Close()
			body, e := io.ReadAll(res.Body)
			respBody = body
			if e != nil {
				err = fmt.Errorf("error reading response body: %v", e)
			}
		}
	}

	return res, respBody, err
}

func (c *Client) close() {
	c.client.CloseIdleConnections()
}
