package gocourtlistener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Results  []SearchResult `json:"results"`
}

// SearchResult represents a single search result.
type SearchResult struct {
	AbsoluteURL              string    `json:"absolute_url"`
	Attorney                 string    `json:"attorney"`
	CaseName                 string    `json:"caseName"`
	CaseNameFull             string    `json:"caseNameFull"`
	Citation                 []string  `json:"citation"`
	CiteCount                int       `json:"citeCount"`
	ClusterID                int       `json:"cluster_id"`
	Court                    string    `json:"court"`
	CourtCitationString      string    `json:"court_citation_string"`
	CourtID                  string    `json:"court_id"`
	DateArgued               *string   `json:"dateArgued"`
	DateFiled                string    `json:"dateFiled"`
	DateReargued             *string   `json:"dateReargued"`
	DateReargumentDenied     *string   `json:"dateReargumentDenied"`
	DocketNumber             string    `json:"docketNumber"`
	DocketID                 int       `json:"docket_id"`
	Judge                    string    `json:"judge"`
	LexisCite                string    `json:"lexisCite"`
	Meta                     Meta      `json:"meta"`
	NeutralCite              string    `json:"neutralCite"`
	NonParticipatingJudgeIDs []int     `json:"non_participating_judge_ids"`
	Opinions                 []Opinion `json:"opinions"`
	PanelIDs                 []int     `json:"panel_ids"`
	PanelNames               []string  `json:"panel_names"`
	Posture                  string    `json:"posture"`
	ProceduralHistory        string    `json:"procedural_history"`
	ScdbID                   string    `json:"scdb_id"`
	SiblingIDs               []int     `json:"sibling_ids"`
	Source                   string    `json:"source"`
	Status                   string    `json:"status"`
	SuitNature               string    `json:"suitNature"`
	Syllabus                 string    `json:"syllabus"`
}

// Opinion represents an opinion object within a search result.
type Opinion struct {
	AuthorID    interface{} `json:"author_id"`
	Cites       []int       `json:"cites"`
	DownloadURL string      `json:"download_url"`
	ID          int         `json:"id"`
	JoinedByIDs []int       `json:"joined_by_ids"`
	LocalPath   string      `json:"local_path"`
	Meta        Meta        `json:"meta"`
	OrderingKey *string     `json:"ordering_key"`
	PerCuriam   bool        `json:"per_curiam"`
	Sha1        string      `json:"sha1"`
	Snippet     string      `json:"snippet"`
	Type        string      `json:"type"`
}

// Meta contains metadata for search results.
type Meta struct {
	Timestamp   string `json:"timestamp"`
	DateCreated string `json:"date_created"`
	Score       Score  `json:"score"`
}

// Score contains the BM25 relevance score.
type Score struct {
	BM25 float64 `json:"bm25"`
}

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
