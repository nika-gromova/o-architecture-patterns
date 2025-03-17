package models

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type ErrNotFoundS struct {
	message string
}

func (e ErrNotFoundS) Error() string {
	return e.message
}

func (e ErrNotFoundS) ToHttpCode() int {
	return http.StatusNotFound
}

func (e ErrNotFoundS) ToGrpcCode() codes.Code {
	return codes.NotFound
}

type ErrInvalidArgumentS struct {
	message string
}

func (e ErrInvalidArgumentS) Error() string {
	return e.message
}

func (e ErrInvalidArgumentS) ToHttpCode() int {
	return http.StatusBadRequest
}

func (e ErrInvalidArgumentS) ToGrpcCode() codes.Code {
	return codes.InvalidArgument
}

type ErrAlreadyExistsS struct {
	message string
}

func (e ErrAlreadyExistsS) Error() string {
	return e.message
}

func (e ErrAlreadyExistsS) ToHttpCode() int {
	return http.StatusConflict
}

func (e ErrAlreadyExistsS) ToGrpcCode() codes.Code {
	return codes.AlreadyExists
}

var (
	ErrNotFound = ErrNotFoundS{
		message: "not found",
	}
	ErrInvalidArgument = ErrInvalidArgumentS{
		message: "invalid argument",
	}
	ErrAlreadyExists = ErrAlreadyExistsS{
		message: "already exists",
	}
)
