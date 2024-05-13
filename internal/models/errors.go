package models

import (
	"net/http"

	"github.com/doxanocap/pkg/errs"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// default http error responses
var (
	HttpBadRequest          = errs.NewHttp(http.StatusBadRequest, "bad request")
	HttpNotFound            = errs.NewHttp(http.StatusNotFound, "not found")
	HttpInternalServerError = errs.NewHttp(http.StatusInternalServerError, "internal server error")
	HttpConflict            = errs.NewHttp(http.StatusConflict, "conflict")
	HttpUnauthorized        = errs.NewHttp(http.StatusUnauthorized, "unauthorized")
)

// custom errors for special cases
var (
	ErrInvalidEmail = errs.NewHttp(http.StatusBadRequest, "invalid email")
)
