package common

import (
	"AltaStore/business"
	"net/http"
)

const (
	errInternalServerError responseCode = "500"
	errNotFound            responseCode = "404"
	errHasBeenModified     responseCode = "400"
	errNotHavePermission   responseCode = "401"
	errPasswordMisMatch    responseCode = "403"
	errInvalidSpec         responseCode = "400"
	errDataExists          responseCode = "409"
	errUnAuthorized        responseCode = "401"
	errInvalidData         responseCode = "400"
)

// Mengembalikan respons status dari permintaan
func NewBusinessErrorResponse(err error) (int, ControllerResponse) {
	return errorMapping(err)
}

func errorMapping(err error) (int, ControllerResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()

	case business.ErrNotFound:
		return newNotFoundResponse()

	case business.ErrHasBeenModified:
		return newHasBeenModifiedResponse()

	case business.ErrNotHavePermission:
		return newNotHavePermission()

	case business.ErrPasswordMisMatch:
		return newErrPasswordMisMatch()

	case business.ErrInvalidSpec:
		return newErrInvalidSpec()

	case business.ErrDataExists:
		return newErrDataExists()

	case business.ErrUnAuthorized:
		return newErrUnAuthorized()
	case business.ErrInvalidData:
		return newErrInvalidData()
	}
}

func newInternalServerErrorResponse() (int, ControllerResponse) {
	return http.StatusInternalServerError,
		ControllerResponse{errInternalServerError, "Internal Server Error", map[string]interface{}{}}
}

func newNotFoundResponse() (int, ControllerResponse) {
	return http.StatusNotFound,
		ControllerResponse{errNotFound, "Data Not Found", map[string]interface{}{}}
}

func newHasBeenModifiedResponse() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errHasBeenModified, "Data Has Been Modified", map[string]interface{}{}}
}

func newNotHavePermission() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errHasBeenModified, "Not Have Permission", map[string]interface{}{}}
}

func newErrPasswordMisMatch() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errHasBeenModified, "Wrong Password", map[string]interface{}{}}
}

func newErrInvalidSpec() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errInvalidSpec, "Bad Request", map[string]interface{}{}}
}

func newErrDataExists() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errDataExists, "Data Exists", map[string]interface{}{}}
}

func newErrUnAuthorized() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errUnAuthorized, "Unauthorized", map[string]interface{}{}}
}

func newErrInvalidData() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errInvalidData, "Invalid Data", map[string]interface{}{}}
}
