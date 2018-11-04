package main

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
	V int `json:"v"`
}

// Request and Responses for GetIPsActiveSince
type getIPsActiveSinceRequest struct {
	S string `json:"s"`
}
type getIPsActiveSinceResponse struct {
	V string `json:"v"`
}
