package server

import (
	"fmt"
	"net/http"
	"net/url"
)

func ForwardRequests(r *http.Request, port int) (*http.Response, error) {
	// Parse the target URL
	target := fmt.Sprintf("http://localhost:%d", port)
	targetURL, err := url.Parse(target)
	if err != nil {
		return &http.Response{}, err
	}

	// Create a new request based on the original request
	req, err := http.NewRequest(r.Method, targetURL.ResolveReference(r.URL).String(), r.Body)
	if err != nil {
		return &http.Response{}, err
	}

	// Copy the headers from the original request
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
