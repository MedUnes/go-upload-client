package vo

//
// ListingQueryVO ...
//
type ListingQueryVO struct {
	Bucket string `json:"bucket"`
	Cursor string `json:"cursor"`
	Path   string `json:"path"`
	Limit  uint32 `json:"limit"`
	Type   byte   `json:"type"`
}
