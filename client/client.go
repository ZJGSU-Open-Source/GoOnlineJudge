package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ApiClient struct {
	Client      *http.Client
	APIPath     string
	AccessToken string
}

func New(apiPath, accessToken string) *ApiClient {
	return &ApiClient{
		Client:      &http.Client{},
		APIPath:     apiPath,
		AccessToken: accessToken,
	}
}

func (api *ApiClient) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c := resp.StatusCode; !(200 <= c && c <= 299) {
		return nil, errors.New(fmt.Sprintf("response code %d", c))
	}

	if v != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(data, v)
		// err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (api *ApiClient) NewRequest(method string, rawurl string, body interface{}) (*http.Request, error) {
	requrl := api.APIPath + rawurl

	if len(api.AccessToken) > 0 {
		parsed, _ := url.Parse(requrl)
		qry := parsed.Query()
		qry.Add("access_token", api.AccessToken)
		parsed.RawQuery = qry.Encode()
		requrl = parsed.String()
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method,
		requrl,
		buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}
