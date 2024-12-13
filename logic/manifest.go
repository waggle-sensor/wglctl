package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

// FetchVSNs fetches the VSNs from the given URL and returns them in JSON.
// 	url: Manifest API url.
func FetchVSNs(url string) ([]byte, error) {
	fmt.Printf("Retrieving Vsn list...\n")
	// Make an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Extract and collect VSNs
	vsnSet := make(map[string]struct{})
	for _, item := range data {
		if vsn, ok := item["vsn"].(string); ok {
			vsnSet[vsn] = struct{}{}
		}
	}

	// Convert the set to a sorted slice
	vsns := make([]string, 0, len(vsnSet))
	for vsn := range vsnSet {
		vsns = append(vsns, vsn)
	}
	sort.Strings(vsns)

	// Encode the sorted slice into JSON
	vsnsJSON, err := json.Marshal(vsns)
	if err != nil {
		return nil, fmt.Errorf("failed to encode VSNs to JSON: %w", err)
	}

	return vsnsJSON, nil
}
