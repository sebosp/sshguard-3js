package main

import "context"
import "github.com/go-kit/kit/endpoint"

// Request and Responses for GetIPDetails
type getIPDetailsRequest struct {
	S string `json:"s"`
}
type getIPDetailsResponse struct {
	V   string `json: "v"`
	Err string `json: "err, omitempty"` // errors don't JSON-marshall, so we use a string (copied from the tutoral, XXX: Investigate further)
}

// Request and Responses for GetIPCount
type getIPCountRequest struct {
	S string `json:"s"`
}
type getIPCountResponse struct {
	V string `json: "v"`
}

// Request and Responses for GetIPsActiveSince
type getIPsActiveSinceRequest struct {
	S string `json:"s"`
}
type getIPsActiveSinceResponse struct {
	V string `json: "v"`
}

func makeGetIPDetailsEndpoint(svc BlacklistService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getIPDetailsRequest)
		v, err := svc.GetIPDetails(req.S)
		if err != nil {
			return getIPDetailsResponse{v, err.Error()}, nil
		}
		return getIPDetailsResponse{v, ""}, nil
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
		v := svc.GetIPsActiveSince(req.S)
		return getIPsActiveSinceResponse{v}, nil
	}
}
