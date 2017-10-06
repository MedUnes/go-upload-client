package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"

	auth "myracloud.com/myra-upload/api/authentication"
)

//
// API ...
//
type API struct {
	APIKey  string
	Secret  string
	BaseURI string

	signature *auth.MyraSignature
	client    *http.Client
}

//
// NewAPI creates a new instance of API
//
func NewAPI(APIKey string, Secret string, BaseURI string, ProxyURL string) (*API, error) {
	signature := &auth.MyraSignature{
		Secret: Secret,
	}

	transport := &http.Transport{}

	if ProxyURL != "" {
		purl, err := url.Parse(ProxyURL)

		if err != nil {
			return nil, err
		}

		transport.Proxy = http.ProxyURL(purl)
	}

	client := &http.Client{
		Transport: transport,
	}

	return &API{
		APIKey:    APIKey,
		Secret:    Secret,
		BaseURI:   BaseURI,
		signature: signature,
		client:    client,
	}, nil
}

func (a *API) rawRequest(method string, url string, mimeType string, reader io.Reader) (*http.Response, error) {
	content, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, a.BaseURI+"/v2"+url, bytes.NewReader(content))

	if err != nil {
		return nil, err
	}

	t := time.Now().Format(time.RFC3339)

	sig, err := a.signature.Build(
		string(content),
		method,
		"/v2"+url,
		mimeType,
		t,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "MYRA "+a.APIKey+":"+sig)
	request.Header.Add("Date", t)
	request.Header.Add("Content-Type", mimeType)

	ret, err := a.client.Do(request)

	if err != nil {
		return nil, err
	}

	if ret.StatusCode == 403 {
		return nil, errors.New("Permission denied.")
	}

	return ret, nil
}

func (a *API) request(method string, url string, obj interface{}) (*http.Response, error) {
	var content []byte
	var err error

	if obj != nil {
		content, err = json.Marshal(&obj)

		if err != nil {
			return nil, err
		}
	} else {
		content = []byte("")
	}

	return a.rawRequest(method, url, "application/json", bytes.NewReader(content))
}

func (a *API) unmarshalResponse(res *http.Response, data interface{}) error {
	tmp, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		err = json.Unmarshal(tmp, data)

		if err != nil {
			return err
		}

		t := reflect.ValueOf(data).Elem()

		if t.Kind() == reflect.Struct {
			if t.FieldByName("Error").Bool() {
				msg := t.FieldByName("ErrorMessage").String()

				return errors.New(msg)
			}
		}

	} else {
		return fmt.Errorf("API returned status code: %d\n%s\n", res.StatusCode, tmp)
	}

	return nil
}
