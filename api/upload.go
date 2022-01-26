package api

import (
	"io"
	"net/http"

	"github.com/Myra-Security-GmbH/go-upload-client/api/vo"
)

//
// UploadFile ...
//
func (a *API) UploadFile(domain string, bucket string, path string, mimeType string, reader io.Reader) error {
	ret, err := a.rawRequest(http.MethodPut, "/upload/"+domain+"/"+bucket+path, mimeType, reader)

	if err != nil {
		return err
	}

	resultVO := vo.ResultVO{}

	err = a.unmarshalResponse(ret, &resultVO)

	if err != nil {
		return err
	}

	return nil
}
