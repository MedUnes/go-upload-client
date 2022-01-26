package vo

import "github.com/Myra-Security-GmbH/go-upload-client/listing"

//
// BaseEntityVO ...
//
type BaseEntityVO struct {
	Modified *listing.DateTime `json:"modified,omitempty"`
	Created  *listing.DateTime `json:"created,omitempty"`
	ID       *uint64           `json:"id,omitempty"`
}
