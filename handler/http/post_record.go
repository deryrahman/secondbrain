package handler

import (
	"net/http"

	"github.com/deryrahman/secondbrain/model"
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

		req, err := NewHTTPRequest[model.PostRecordsJSONRequestBody](r)
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

		resp.WriteJSON(model.PostRecordsJSONResponseBody{Id: id})
	}
}
