// types.go
// Package gocourtlistener provides a client for interacting with the Courtlistener REST API.
package gocourtlistener

import (
	"net/http"
	"time"
)

// HTTPClient abstracts the Do method so that any client (e.g., http.Client) can be used.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the base client for interacting with the Courtlistener API.
type Client struct {
	BaseURL    string
	Email      string
	HTTPClient HTTPClient
}

// NewClient creates a new Courtlistener API client.
// If client is nil, a default http.Client with a 10-second timeout is used.
func NewClient(baseURL, email string, client HTTPClient) *Client {
	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}
	return &Client{
		BaseURL:    baseURL,
		Email:      email,
		HTTPClient: client,
	}
}

// Pagination is a common structure for paginated responses.
// Note: Some endpoints may represent count as a URL string while others as a number.
type Pagination struct {
	// Count can be either an int or a string depending on the endpoint.
	Count    FlexibleCount `json:"count"`
	Next     string        `json:"next"`
	Previous *string       `json:"previous"`
}

// Meta contains common metadata for API responses.
type Meta struct {
	Timestamp   string `json:"timestamp"`
	DateCreated string `json:"date_created"`
	// Score is included when relevant.
	Score *Score `json:"score,omitempty"`
}

// Score represents the BM25 relevance score in some API responses.
type Score struct {
	BM25 float64 `json:"bm25"`
}

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

// OpinionsResponse represents the response from the opinions endpoint.
type OpinionsResponse struct {
	Count    FlexibleCount `json:"count"`
	Next     string        `json:"next"`
	Previous *string       `json:"previous"`
	Results  []Opinion     `json:"results"`
}

// DocketsResponse represents the response structure returned by the dockets endpoint.
type DocketsResponse struct {
	Count    FlexibleCount `json:"count"`
	Next     string        `json:"next"`
	Previous *string       `json:"previous"`
	Results  []Docket      `json:"results"`
}

// Docket represents a single docket record.
type Docket struct {
	AbsoluteURL                    string        `json:"absolute_url"`
	AppealFrom                     *string       `json:"appeal_from"`
	AppealFromStr                  string        `json:"appeal_from_str"`
	AppellateCaseTypeInformation   string        `json:"appellate_case_type_information"`
	AppellateFeeStatus             string        `json:"appellate_fee_status"`
	AssignedTo                     *string       `json:"assigned_to"`
	AssignedToStr                  string        `json:"assigned_to_str"`
	AudioFiles                     []string      `json:"audio_files"`
	Blocked                        bool          `json:"blocked"`
	CaseName                       string        `json:"case_name"`
	CaseNameFull                   string        `json:"case_name_full"`
	CaseNameShort                  string        `json:"case_name_short"`
	Cause                          string        `json:"cause"`
	Clusters                       []string      `json:"clusters"`
	Court                          string        `json:"court"`
	CourtID                        string        `json:"court_id"`
	DateArgued                     *string       `json:"date_argued"`
	DateBlocked                    *string       `json:"date_blocked"`
	DateCertDenied                 *string       `json:"date_cert_denied"`
	DateCertGranted                *string       `json:"date_cert_granted"`
	DateCreated                    string        `json:"date_created"`
	DateFiled                      *string       `json:"date_filed"`
	DateLastFiling                 string        `json:"date_last_filing"`
	DateLastIndex                  *string       `json:"date_last_index"`
	DateModified                   string        `json:"date_modified"`
	DateReargued                   *string       `json:"date_reargued"`
	DateReargumentDenied           *string       `json:"date_reargument_denied"`
	DateTerminated                 *string       `json:"date_terminated"`
	DocketNumber                   string        `json:"docket_number"`
	DocketNumberCore               string        `json:"docket_number_core"`
	FederalDefendantNumber         *int          `json:"federal_defendant_number"`
	FederalDNCaseType              string        `json:"federal_dn_case_type"`
	FederalDNJudgeInitialsAssigned string        `json:"federal_dn_judge_initials_assigned"`
	FederalDNJudgeInitialsReferred string        `json:"federal_dn_judge_initials_referred"`
	FederalDNOfficeCode            string        `json:"federal_dn_office_code"`
	FilepathIA                     string        `json:"filepath_ia"`
	FilepathIAJSON                 string        `json:"filepath_ia_json"`
	IADateFirstChange              string        `json:"ia_date_first_change"`
	IANeedsUpload                  bool          `json:"ia_needs_upload"`
	IAUploadFailureCount           *int          `json:"ia_upload_failure_count"`
	ID                             int           `json:"id"`
	IDBData                        interface{}   `json:"idb_data"`
	JurisdictionType               string        `json:"jurisdiction_type"`
	JuryDemand                     string        `json:"jury_demand"`
	MDLStatus                      string        `json:"mdl_status"`
	NatureOfSuit                   string        `json:"nature_of_suit"`
	OriginalCourtInfo              interface{}   `json:"original_court_info"`
	PacerCaseID                    string        `json:"pacer_case_id"`
	Panel                          []interface{} `json:"panel"`
	PanelStr                       string        `json:"panel_str"`
	ParentDocket                   interface{}   `json:"parent_docket"`
	ReferredTo                     interface{}   `json:"referred_to"`
	ReferredToStr                  string        `json:"referred_to_str"`
	ResourceURI                    string        `json:"resource_uri"`
	Slug                           string        `json:"slug"`
	Source                         int           `json:"source"`
	Tags                           []interface{} `json:"tags"`
}

// OriginatingCourtInformationResponse represents the response structure from the originating-court-information endpoint.
type OriginatingCourtInformationResponse struct {
	Count    FlexibleCount                 `json:"count"`
	Next     string                        `json:"next"`
	Previous *string                       `json:"previous"`
	Results  []OriginatingCourtInformation `json:"results"`
}

// OriginatingCourtInformation represents a single originating court record.
type OriginatingCourtInformation struct {
	AssignedTo       *string `json:"assigned_to"`
	AssignedToStr    string  `json:"assigned_to_str"`
	CourtReporter    string  `json:"court_reporter"`
	DateCreated      string  `json:"date_created"`
	DateDisposed     *string `json:"date_disposed"`
	DateFiled        *string `json:"date_filed"`
	DateFiledNOA     *string `json:"date_filed_noa"`
	DateJudgment     *string `json:"date_judgment"`
	DateJudgmentEOD  *string `json:"date_judgment_eod"`
	DateModified     string  `json:"date_modified"`
	DateReceivedCOA  *string `json:"date_received_coa"`
	DocketNumber     string  `json:"docket_number"`
	ID               int     `json:"id"`
	OrderingJudge    *string `json:"ordering_judge"`
	OrderingJudgeStr string  `json:"ordering_judge_str"`
	ResourceURI      string  `json:"resource_uri"`
}
