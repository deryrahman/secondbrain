package server

import (
	"fmt"
	"net/http"

	codegenHandler "github.com/deryrahman/secondbrain/codegen/handler"
	handler "github.com/deryrahman/secondbrain/handler/http"
	"github.com/deryrahman/secondbrain/service"
)

var _ codegenHandler.ServerInterface = (*httpServer)(nil)
var _ http.Handler = (*httpServer)(nil)

type httpServer struct {
	baseURL       string
	recordService service.RecordService
}

func NewHTTPServer(baseURL string, recordService service.RecordService) (*httpServer, error) {
	if recordService == nil {
		return nil, fmt.Errorf("recordService is nil")
	}
	return &httpServer{
		baseURL:       baseURL,
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
	handler.HandleHTTPPostRecords(h.recordService)(w, r)
}

func (h *httpServer) GetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`pong`))
}
