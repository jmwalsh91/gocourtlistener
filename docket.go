// Package gocourtlistener provides a client for interacting with the Courtlistener REST API.
package gocourtlistener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)


// Dockets retrieves docket records from the Courtlistener API.
// The params argument can be used to supply query parameters (e.g. "cursor", "count").
func (c *Client) Dockets(params map[string]string) (*DocketsResponse, error) {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// Append the "dockets/" endpoint.
	newPath, err := url.JoinPath(base.Path, "dockets/")
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
		return nil, fmt.Errorf("Dockets: unexpected status code %d", resp.StatusCode)
	}

	var result DocketsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
