package mediaunlocktest

import (
	"net/http"
)

func HoyTV(c http.Client) Result {
	resp, err := GET(c, "https://hoytv-live-stream.hoy.tv/ch78/index-fhd.m3u8")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == 403 {
		return Result{Status: StatusNo}
	}
	
	if resp.StatusCode == 200 {
		return Result{Status: StatusOK}
	}

	return Result{Status: StatusUnexpected}
}