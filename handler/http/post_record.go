package handler

import (
	"net/http"

	model "github.com/deryrahman/secondbrain/model/handler"
	"github.com/deryrahman/secondbrain/pkg/log"
	"github.com/deryrahman/secondbrain/service"
)

func HandleHTTPPostRecords(logger log.Logger, s service.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := NewHTTPResponse(logger, w)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		req, err := NewHTTPRequest[model.PostRecordsRequest](r)
		if err != nil {
			logger.Error(err)
			resp.WriteError(err)
			return
		}

		reqBody, err := req.GetJSONBody()
		if err != nil {
			logger.Error(err)
			resp.WriteError(err)
			return
		}

		id, err := s.CreateRecord(r.Context(), *reqBody.Content, *reqBody.Tags...)
		if err != nil {
			logger.Error(err)
			resp.WriteError(err)
			return
		}

		resp.WriteJSON(model.PostRecordsResponse{ID: id})
	}
}
