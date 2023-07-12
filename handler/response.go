package handler

type HTTPResponse interface {
	WriteJSON(content any)
	WriteError(err error)
}
