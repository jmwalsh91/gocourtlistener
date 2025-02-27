package gocourtlistener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Search performs a GET request to the search endpoint with the provided query string.
func (c *Client) Search(query string) (*SearchResponse, error) {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// Append the "search/" endpoint.
	newPath, err := url.JoinPath(base.Path, "search/")
	if err != nil {
		return nil, err
	}
	base.Path = newPath
	q := base.Query()
	q.Set("q", query)
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
		return nil, fmt.Errorf("Search: unexpected status code %d", resp.StatusCode)
	}

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
