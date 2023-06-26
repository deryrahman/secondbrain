package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deryrahman/secondbrain/handler"
)

var _ handler.HTTPResponse = (*response)(nil)

type response struct {
	http.ResponseWriter
}

func NewHTTPResponse(httpWriter http.ResponseWriter) (*response, error) {
	if httpWriter == nil {
		return nil, fmt.Errorf("httpWriter is nil")
	}
	return &response{httpWriter}, nil
}

func (r *response) WriteJSON(content any) error {
	raw, err := json.Marshal(content)
	if err != nil {
		return err
	}
	_, err = r.Write(raw)
	return err
}
