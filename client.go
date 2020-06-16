package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/plugin/httptrace"
)

// FibonacciClient represents a client
// for the fibonacci service
type FibonacciClient struct {
	url        string
	httpClient *http.Client
}

// NewFibonacciClient returns the a new
// client for connecting to the fibonacci serivce
func NewFibonacciClient() FibonacciClient {
	url := "http://localhost:8080/fibonacci"
	client := &http.Client{Timeout: time.Second * 10}

	return FibonacciClient{
		url:        url,
		httpClient: client,
	}
}

// FibonacciNumber returns the fibonacci number
// given the sequence number n by contacting the
// fibonacci service
func (f FibonacciClient) FibonacciNumber(ctx context.Context, sequenceNumber int) (int, error) {
	fibonacciReq := FibonacciRequest{SequenceNumber: sequenceNumber}
	data, err := json.Marshal(fibonacciReq)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", f.url, bytes.NewBuffer(data))
	if err != nil {
		return 0, fmt.Errorf("error while creating fibonacci request for sequence number %d: %v",
			sequenceNumber, err)
	}

	req.Header.Set("Content-Type", "application/json")

	ctx, req = httptrace.W3C(ctx, req)
	httptrace.Inject(ctx, req)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error while sending fibonacci request for sequence number %d: %v",
			sequenceNumber, err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error while reading fibonacci response for sequence number %d: %v",
			sequenceNumber, err)
	}

	var fibonacciRes FibonacciResponse
	err = json.Unmarshal(respBody, &fibonacciRes)
	if err != nil {
		return 0, fmt.Errorf("error while parsing fibonacci response for sequence number %d: %v",
			sequenceNumber, err)
	}

	return fibonacciRes.FibonacciNumber, nil
}
