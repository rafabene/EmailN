package endpoints

import (
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type EndPointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndPointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)
		if err != nil {

			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
			} else {
				render.Status(r, http.StatusBadRequest)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
