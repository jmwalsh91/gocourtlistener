// Package gocourtlistener provides a client for interacting with the Courtlistener REST API.
package gocourtlistener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Opinions retrieves judicial opinions from the Courtlistener API.
// You may pass additional query parameters (e.g., "cursor", "fields") via the params map.
func (c *Client) Opinions(params map[string]string) (*OpinionsResponse, error) {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// Append the "opinions/" endpoint.
	newPath, err := url.JoinPath(base.Path, "opinions/")
	if err != nil {
		return nil, err
	}
	base.Path = newPath

	// Set any query parameters.
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
		return nil, fmt.Errorf("Opinions: unexpected status code %d", resp.StatusCode)
	}

	var result OpinionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
