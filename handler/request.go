package handler

type HTTPRequest[body any] interface {
	GetJSONBody() (body, error)
	GetQueryParams() map[string][]string
}
