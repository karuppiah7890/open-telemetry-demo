package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/open-telemetry/opentelemetry-collector/translator/conventions"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/plugin/othttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
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

func initializeTracer() {
	exp, err := otlp.NewExporter(otlp.WithInsecure(),
		otlp.WithAddress("localhost:55680"),
		otlp.WithGRPCDialOption(grpc.WithBlock()))
	if err != nil {
		log.Fatalf("Failed to create the collector exporter: %v", err)
	}

	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithResource(resource.New(
			// the service name used to display traces in Jaeger
			kv.Key(conventions.AttributeServiceName).String("fibonacci-service"),
		)),
		sdktrace.WithSyncer(exp))
	if err != nil {
		log.Fatalf("error creating trace provider: %v\n", err)
	}
	global.SetTraceProvider(tp)
}

func main() {
	initializeTracer()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/fibonacci", othttp.NewHandler(http.HandlerFunc(fibonacciHandler), "fibonacci-endpoint").ServeHTTP)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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

	span := trace.SpanFromContext(ctx)
	span.SetAttribute("sequenceNumber", fibonacciReq.SequenceNumber)

	result, err := fibonacciNumber(ctx, fibonacciReq.SequenceNumber)
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
}

// fibonacciNumber returns the fibonacci number
// given the sequence number n
func fibonacciNumber(ctx context.Context, sequenceNumber int) (int, error) {
	if sequenceNumber == 0 || sequenceNumber == 1 {
		return sequenceNumber, nil
	}

	fibonacciClient := NewFibonacciClient()

	value1, err := fibonacciClient.FibonacciNumber(ctx, sequenceNumber-1)
	if err != nil {
		return 0, err
	}

	value2, err := fibonacciClient.FibonacciNumber(ctx, sequenceNumber-2)
	if err != nil {
		return 0, err
	}

	result := value1 + value2

	return result, nil
}
