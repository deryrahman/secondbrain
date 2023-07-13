package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deryrahman/secondbrain/handler"
	"github.com/deryrahman/secondbrain/pkg/errors"
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
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		r.Body.Close()
		return value, errors.RootCause(err)
	}

	if err := json.Unmarshal(raw, &value); err != nil {
		return value, errors.RootCause(err)
	}

	return value, nil
}

func (r *request[bodyStruct]) GetQueryParams() map[string][]string {
	return r.URL.Query()
}
