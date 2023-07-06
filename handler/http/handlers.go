package handler

import (
	"net/http"

	"github.com/deryrahman/secondbrain/model"
	"github.com/deryrahman/secondbrain/pkg/log"
	"github.com/deryrahman/secondbrain/service"
)

func HandleHTTPPostRecords(logger log.Logger, s service.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := NewHTTPRequest[model.PostRecordsJSONRequestBody](r)
		if err != nil {
			logger.Error(err)
			return
		}
		reqBody, err := req.GetJSONBody()
		if err != nil {
			logger.Error(err)
			return
		}

		id, err := s.CreateRecord(r.Context(), *reqBody.Content, *reqBody.Tags...)
		if err != nil {
			logger.Error(err)
			return
		}

		resp, err := NewHTTPResponse(w)
		if err != nil {
			logger.Error(err)
			return
		}
		if err := resp.WriteJSON(model.PostRecordsJSONResponseBody{Id: id}); err != nil {
			logger.Error(err)
			return
		}
		logger.Infof("success")
	}
}

func HandleHTTPGetRecords(s service.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := s.GetRecords(r.Context())
		if err != nil {
			return
		}

		resp, _ := NewHTTPResponse(w)
		if err := resp.WriteJSON(model.GetRecordsJSONResponseBody{}); err != nil {
			return
		}
	}
}
