package errors

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomError interface {
	Error() string
	ToHttpCode() int
	ToGrpcCode() codes.Code
}

type defaultCustomError struct {
}

func (e *defaultCustomError) Error() string {
	return ""
}

func (e *defaultCustomError) ToHttpCode() int {
	return http.StatusInternalServerError
}

func (e *defaultCustomError) ToGrpcCode() codes.Code {
	return codes.Internal
}

func InterceptorGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err == nil {
		return resp, err
	}

	custom := errorToCustom(err)
	return nil, status.Error(custom.ToGrpcCode(), err.Error())
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CustomHTTPErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	custom := errorToCustom(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(custom.ToHttpCode())

	resp := errorResponse{
		Code:    custom.ToHttpCode(),
		Message: status.Convert(err).Message(),
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func errorToCustom(err error) CustomError {
	var custom CustomError
	ok := errors.As(err, &custom)
	if !ok {
		return &defaultCustomError{}
	}

	return custom
}
