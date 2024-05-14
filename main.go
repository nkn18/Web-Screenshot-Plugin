package main

import (
	"github.com/extism/go-pdk"
)

//go:export run
func run() int32 {
	WebURL := pdk.InputString()
	req := pdk.NewHTTPRequest(pdk.MethodGet, "https://restpack.io/api/screenshot/v7/capture?access_token={YOUR_RESTPACK_API_TOKEN}="+ WebURL +"&json=true&cache_ttl=604800")
	res := req.Send()
	pdk.OutputMemory(res.Memory())
	return 0
}

func main() {}

