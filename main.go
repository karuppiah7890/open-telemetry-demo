package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FibonacciRequest represents a request to the
// Fibonacci service
type FibonacciRequest struct {
	SequenceNumber int
}

// FibonacciResponse represents a response from the
// Fibonacci service
type FibonacciResponse struct {
	FibonacciNumber int
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer r.Body.Close()

		var fibonacciReq FibonacciRequest
		err = json.Unmarshal(body, &fibonacciReq)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		result, err := fibonacciNumber(fibonacciReq.SequenceNumber)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fibonacciRes := FibonacciResponse{FibonacciNumber: result}
		jsonResponse, err := json.Marshal(fibonacciRes)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// fibonacciNumber returns the fibonacci number
// given the sequence number n
func fibonacciNumber(sequenceNumber int) (int, error) {
	if sequenceNumber == 0 || sequenceNumber == 1 {
		return sequenceNumber, nil
	}

	fibonacciClient := NewFibonacciClient()

	value1, err := fibonacciClient.FibonacciNumber(sequenceNumber - 1)
	if err != nil {
		return 0, err
	}

	value2, err := fibonacciClient.FibonacciNumber(sequenceNumber - 2)
	if err != nil {
		return 0, err
	}

	result := value1 + value2

	return result, nil
}
