package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

func makeGetIPDetailsEndpoint(svc BlacklistService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getIPDetailsRequest)
		v, err := svc.GetIPDetails(req.S)
		if err != nil {
			return getIPDetailsResponse{"", err.Error()}, nil
		}
		jsonv, _ := json.Marshal(v)
		return getIPDetailsResponse{string(jsonv), ""}, nil
	}
}
func makeGetIPCountEndpoint(svc BlacklistService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getIPCountRequest)
		v := svc.GetIPCount(req.S)
		return getIPCountResponse{v}, nil
	}
}
func makeGetIPsActiveSinceEndpoint(svc BlacklistService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getIPsActiveSinceRequest)
		inputEpoch, err := strconv.ParseInt(req.S, 0, 64)
		if err != nil {
			return getIPsActiveSinceResponse{""}, err
		}
		v := svc.GetIPsActiveSince(inputEpoch)
		jsonv, _ := json.Marshal(v)
		return getIPsActiveSinceResponse{string(jsonv)}, nil
	}
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
