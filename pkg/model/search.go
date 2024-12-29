package model

type SearchResults struct {
	Results    []EntityRef `json:"results"`
	Limit      int         `json:"limit"`
	NextOffset int         `json:"nextOffset"`
}
