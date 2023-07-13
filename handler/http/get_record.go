package handler

import (
	"net/http"

	codegenHandler "github.com/deryrahman/secondbrain/codegen/handler"
	model "github.com/deryrahman/secondbrain/model/handler"
	"github.com/deryrahman/secondbrain/pkg/log"
	"github.com/deryrahman/secondbrain/service"
)

func HandleHTTPGetRecords(logger log.Logger, s service.RecordService, params codegenHandler.GetRecordsParams) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := NewHTTPResponse(logger, w)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tags := []string{}
		if params.Tag != nil {
			tags = *params.Tag
		}

		records, err := s.GetRecords(r.Context(), tags...)
		if err != nil {
			logger.Error(err)
			resp.WriteError(err)
			return
		}

		recordResponse := model.GetRecordsResponse{}
		recordResponse.From(records)

		resp.WriteJSON(recordResponse)
	}
}
