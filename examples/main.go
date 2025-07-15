package main

import (
	"fmt"
	"os"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sicko7947/psychoclient"
)

func main() {
	// Enable verbose logging for the uhttp transport
	os.Setenv("UHTTP_VERBOSE", "1")
	client := psychoclient.NewClient(&psychoclient.SessionBuilder{
		UseDefaultClient: true,
	})

	res, respBody, err := client.DoNewRequest(&psychoclient.RequestBuilder{
		Endpoint: "https://tls.peet.ws/api/all",
		Method:   "GET",
		Headers: map[string]string{
			"application": "json",
			"User-Agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
		},
	})

	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}

	fmt.Println(res.Header)
	if respBody != nil {
		fmt.Println(gconv.String(respBody))
	}

}
