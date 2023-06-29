package server

import (
	"errors"
	"fmt"
	"net/http"

	codegenHandler "github.com/deryrahman/secondbrain/codegen/handler"
	handler "github.com/deryrahman/secondbrain/handler/http"
	"github.com/deryrahman/secondbrain/pkg/log"
	"github.com/deryrahman/secondbrain/service"
)

var _ codegenHandler.ServerInterface = (*httpServer)(nil)
var _ http.Handler = (*httpServer)(nil)

type httpServer struct {
	baseURL       string
	logger        log.Logger
	recordService service.RecordService
}

func NewHTTPServer(baseURL string, logger log.Logger, recordService service.RecordService) (*httpServer, error) {
	var err error
	if recordService == nil {
		err = errors.Join(err, fmt.Errorf("recordService is nil"))
	}
	if logger == nil {
		err = errors.Join(err, fmt.Errorf("logger is nil"))
	}
	if err != nil {
		return nil, err
	}

	return &httpServer{
		baseURL:       baseURL,
		logger:        logger,
		recordService: recordService,
	}, nil
}

func (h *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler().ServeHTTP(w, r)
}

func (h *httpServer) Handler() http.Handler {
	return codegenHandler.HandlerWithOptions(h, codegenHandler.ChiServerOptions{
		BaseURL: h.baseURL,
	})
}

func (h *httpServer) GetRecords(w http.ResponseWriter, r *http.Request) {
	handler.HandleHTTPGetRecords(h.recordService)(w, r)
}

func (h *httpServer) PostRecords(w http.ResponseWriter, r *http.Request) {
	handler.HandleHTTPPostRecords(h.logger, h.recordService)(w, r)
}

func (h *httpServer) GetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`pong`))
}
