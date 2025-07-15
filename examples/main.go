package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/sicko7947/psychoclient"
)

func main() {
	// Enable verbose logging for the uhttp transport
	os.Setenv("UHTTP_VERBOSE", "1")

	// Set up panic recovery to capture full stack trace
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC: %v\n", r)
			fmt.Printf("STACK TRACE:\n%s", debug.Stack())
		}
	}()

	fmt.Println("Creating client...")
	client := psychoclient.NewClient(&psychoclient.SessionBuilder{
		UseDefaultClient: false,
		FollowRedirects:  true,
		Timeout:          10 * time.Second,
	})

	fmt.Println("Making request to:", "https://tls.peet.ws/api/all")

	// Let's also try a simpler endpoint first to see if it's endpoint-specific
	testEndpoints := []string{
		"https://tls.peet.ws/api/all",
	}

	for _, endpoint := range testEndpoints {
		fmt.Printf("\n--- Testing endpoint: %s ---\n", endpoint)
		res, respBody, err := client.RoundTripNewRequest(&psychoclient.RequestBuilder{
			Endpoint: endpoint,
			Method:   "GET",
		})

		if err != nil {
			fmt.Printf("Error occurredaa: %v\n", err)
			fmt.Printf("Error type: %T\n", err)
			fmt.Printf("Error details: %+v\n", err)

			// Print stack trace manually
			fmt.Printf("STACK TRACE:\n%s", debug.Stack())

			// Don't panic immediately, try the next endpoint
			continue
		}

		if res != nil {
			fmt.Printf("Success! Status: %s\n", res.Status)
			defer res.Body.Close()
		} else {
			fmt.Println("Success! No response object")
		}

		if respBody != nil {
			fmt.Printf("Response length: %d bytes\n", len(respBody))
			if len(respBody) > 200 {
				fmt.Printf("Response preview: %s...\n", string(respBody[:200]))
			} else {
				fmt.Printf("Response: %s\n", string(respBody))
			}
		} else {
			fmt.Println("No response body")
		}

		// If we got here, the request succeeded
		return
	}

	fmt.Println("All endpoints failed")
}
