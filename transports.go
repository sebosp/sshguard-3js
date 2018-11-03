package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := blacklistService{}
	getIPDetailsHandler := httptransport.NewServer(
		makeGetIPDetailsEndpoint(svc),
		decodeGetIPDetailsRequest,
		encodeResponse,
	)
	getIPCountHandler := httptransport.NewServer(
		makeGetIPCountEndpoint(svc),
		decodeGetIPCountRequest,
		encodeResponse,
	)
	getIPsActiveSinceHandler := httptransport.NewServer(
		makeGetIPsActiveSinceEndpoint(svc),
		decodeGetIPsActiveSinceRequest,
		encodeResponse,
	)
	http.Handle("/getIPDetails", getIPDetailsHandler)
	http.Handle("/getIPCount", getIPCountHandler)
	http.Handle("/getIPsActiveSince", getIPsActiveSinceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeGetIPDetailsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getIPDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeGetIPCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getIPCountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeGetIPsActiveSinceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getIPsActiveSinceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
