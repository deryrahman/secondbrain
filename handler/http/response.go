package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deryrahman/secondbrain/handler"
	"github.com/deryrahman/secondbrain/pkg/errors"
	"github.com/deryrahman/secondbrain/pkg/log"
)

var _ handler.HTTPResponse = (*response)(nil)

type response struct {
	writer http.ResponseWriter
	logger log.Logger
}

func NewHTTPResponse(logger log.Logger, httpWriter http.ResponseWriter) (*response, error) {
	var err error
	if logger == nil {
		err = errors.Join(err, fmt.Errorf("logger is nil"))
	}
	if httpWriter == nil {
		err = errors.Join(err, fmt.Errorf("httpWriter is nil"))
	}
	if err != nil {
		return nil, errors.RootCause(err)
	}
	return &response{logger: logger, writer: httpWriter}, nil
}

func (r *response) WriteJSON(content any) {
	raw, err := json.Marshal(content)
	if err != nil {
		r.logger.Error(errors.RootCause(err))
	}
	_, err = r.writer.Write(raw)
	if err != nil {
		r.logger.Error(errors.RootCause(err))
	}
}

func (r *response) WriteError(err error) {
	var causeErr errors.RootCauseError
	statusCode := http.StatusInternalServerError
	if errors.As(err, &causeErr) {
		statusCode = causeErr.StatusCode()
	}
	r.writer.WriteHeader(statusCode)
}
