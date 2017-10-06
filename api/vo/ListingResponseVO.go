package vo

//
// ListingResponseVO ...
//
type ListingResponseVO struct {
	Error        bool            `json:"error"`
	ErrorMessage string          `json:"errorMessage"`
	Count        uint64          `json:"count"`
	CursorNext   string          `json:"cursorNext"`
	CursorPrev   string          `json:"cursorPrev"`
	Result       []ListingItemVO `json:"result"`
}
