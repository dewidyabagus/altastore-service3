package business

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInternalServer = errors.New("Internal_Server_Error")

	ErrHasBeenModified = gorm.ErrInvalidData

	ErrNotFound = gorm.ErrRecordNotFound

	ErrInvalidSpec = errors.New("Given_Spec_Is_Not_Valid")

	ErrDataExists = errors.New("Data_Exists")

	ErrPasswordMisMatch = errors.New("Wrong_Password")

	ErrLoginFailed = errors.New("Login_Failed")

	ErrNotHavePermission = errors.New("Not_Have_Permission")

	ErrUnAuthorized = errors.New("UnAuthorized")

	ErrInvalidData = errors.New("Invalid_Data")
)
