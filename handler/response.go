package handler

type HTTPResponse interface {
	WriteJSON(content any) error
}
