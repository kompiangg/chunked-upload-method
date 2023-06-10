package handler

import "github.com/kompiangg/shipper-fp/pkg/http"

// Default docs response
// Copas to make your work easier

type HTTPBaseResp struct {
	Error *http.HTTPErrorBaseResponse `json:"error"`
}

type HTTPErrResp struct {
	Error HTTPBaseResp
	Data  http.HTTPErrorBaseResponse `json:"data"`
}

type HTTPPingResp struct {
	Error HTTPBaseResp
	Data  string `json:"data"`
}
