package handler

import (
	"net/http"
	"net/url"
)

type HTTPRequest[body any] interface {
	GetJSONBody() (body, error)
	GetHeaders() http.Header
	GetURL() url.URL
}
