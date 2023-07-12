package handler

import (
	"net/http"

	"github.com/deryrahman/secondbrain/service"
)

func HandleHTTPGetRecords(s service.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
