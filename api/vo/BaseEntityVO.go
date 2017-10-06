package vo

import "myracloud.com/myra-upload/listing"

//
// BaseEntityVO ...
//
type BaseEntityVO struct {
	Modified *listing.DateTime `json:"modified,omitempty"`
	Created  *listing.DateTime `json:"created,omitempty"`
	ID       *uint64           `json:"id,omitempty"`
}
