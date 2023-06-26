package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/deryrahman/secondbrain/handler"
)

var _ handler.HTTPRequest[any] = (*request[any])(nil)

type request[bodyStruct any] struct {
	*http.Request
}

func NewHTTPRequest[bodyStruct any](httpReq *http.Request) (*request[bodyStruct], error) {
	if httpReq == nil {
		return nil, fmt.Errorf("httpReq is nil")
	}
	return &request[bodyStruct]{httpReq}, nil
}

func (r *request[bodyStruct]) GetJSONBody() (bodyStruct, error) {
	var value bodyStruct
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		r.Body.Close()
		return value, err
	}

	if err := json.Unmarshal(raw, &value); err != nil {
		return value, err
	}

	return value, err
}

func (r *request[bodyStruct]) GetHeaders() http.Header {
	return r.GetHeaders()
}

func (r *request[bodyStruct]) GetURL() url.URL {
	return r.GetURL()
}
