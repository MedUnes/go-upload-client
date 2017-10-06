package vo

//
// ResultVO ...
//
type ResultVO struct {
	Error        bool     `json:"error"`
	ErrorMessage string   `json:"errorMessage"`
	Status       string   `json:"status"`
	StatusCode   uint16   `json:"statusCode"`
	BucketList   []string `json:"bucketList,omitempty"`
}
