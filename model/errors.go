package model

import (
	"github.com/go-chi/render"
	"net/http"
)

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}

func ErrorInvalidRequest(err error)render.Renderer{
	return &ErrResponse{
		Err:            err,
		Code:			400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),

	}
}

func ErrorInternalServer(err error)render.Renderer{
	return &ErrResponse{
		Err:            err,
		Code:			500,
		StatusText:     "Internal server error",
		ErrorText:      err.Error(),

	}
}
