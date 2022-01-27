package vo

import "github.com/Myra-Security-GmbH/go-upload-client/listing"

//
// ListingItemVO ...
//
type ListingItemVO struct {
	Type        byte             `json:"type"`
	Path        string           `json:"path"`
	Basename    string           `json:"basename"`
	Size        uint64           `json:"size"`
	Hash        string           `json:"hash"`
	Modified    listing.DateTime `json:"modified"`
	ContentType string           `json:"contentType"`
}
