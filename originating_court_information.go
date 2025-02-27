// Package gocourtlistener provides a client for interacting with the Courtlistener REST API.
package gocourtlistener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// OriginatingCourtInformation retrieves records from the originating-court-information endpoint.
// The params argument can be used to supply query parameters (e.g. "cursor", "count").
func (c *Client) OriginatingCourtInformation(params map[string]string) (*OriginatingCourtInformationResponse, error) {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// Append the "originating-court-information/" endpoint.
	newPath, err := url.JoinPath(base.Path, "originating-court-information/")
	if err != nil {
		return nil, err
	}
	base.Path = newPath

	q := base.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	base.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.Email != "" {
		req.Header.Set("X-User-Email", c.Email)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OriginatingCourtInformation: unexpected status code %d", resp.StatusCode)
	}

	var result OriginatingCourtInformationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
