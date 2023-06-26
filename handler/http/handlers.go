package handler

import (
	"net/http"

	"github.com/deryrahman/secondbrain/model"
	"github.com/deryrahman/secondbrain/service"
)

func HandleHTTPPostRecords(s service.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, _ := NewHTTPRequest[model.PostRecordsJSONRequestBody](r)
		reqBody, _ := req.GetJSONBody()

		id, _ := s.CreateRecord(r.Context(), *reqBody.Content, *reqBody.Tags...)

		resp, _ := NewHTTPResponse(w)
		resp.WriteJSON(model.PostRecordsJSONResponseBody{Id: id})
	}
}

func HandleHTTPGetRecords(s service.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.GetRecords(r.Context())

		resp, _ := NewHTTPResponse(w)
		resp.WriteJSON(model.GetRecordsJSONResponseBody{})
	}
}
