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
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

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

func fibonacciNumber(sequenceNumber int) (int, error) {
	if sequenceNumber == 0 || sequenceNumber == 1 {
		return sequenceNumber, nil
	}

	return 0, fmt.Errorf("don't know the answer :p")
}
