package errors

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/rules"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func gRPCErrorToHTTPCode(err error) int {
	switch {
	case errors.Is(err, rules.ErrNotFound):
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

func InterceptorGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err == nil {
		return resp, err
	}
	switch {
	case errors.Is(err, rules.ErrNotFound):
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return nil, status.Error(codes.Internal, err.Error())
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// TODO ...
func CustomHTTPErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	httpCode := gRPCErrorToHTTPCode(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	resp := errorResponse{
		Code:    httpCode,
		Message: status.Convert(err).Message(), // Текст gRPC ошибки
	}

	_ = json.NewEncoder(w).Encode(resp)
}
