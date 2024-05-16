package main

import (
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
)

//go:export run
func run() int32 {
	// Get the URL from the input
	WebURL := pdk.InputString()

	// Create the request to the screenshot API
	req := pdk.NewHTTPRequest(pdk.MethodGet, "https://restpack.io/api/screenshot/v7/capture?access_token=xWwNo0styczmArzOH2pDK1kQbIHb052Of8upxU0cRsI98ejB&url=" + WebURL + "&json=true&cache_ttl=604800")

	// Send the request
	res := req.Send()

	// Decode the JSON response
	var jsonResponse map[string]interface{}
	err := json.Unmarshal(res.Body(), &jsonResponse)
	if err != nil {
		pdk.SetErrorString(fmt.Sprintf("Failed to decode JSON: %v", err))
		return 1
	}

	// Extract the "image" URL
	imageURL, ok := jsonResponse["image"].(string)
	if !ok {
		pdk.SetErrorString("image URL not found in response")
		return 1
	}

	// Encode the image URL as JSON
	imageData, err := json.Marshal(map[string]string{"image": imageURL})
	if err != nil {
		pdk.SetErrorString(fmt.Sprintf("Failed to encode image URL: %v", err))
		return 1
	}

	// Allocate memory for the encoded image URL
	mem := pdk.AllocateBytes(imageData)

	// Pass the memory offset
	pdk.OutputMemory(mem)

	return 0
}

func main() {}
