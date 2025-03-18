package logging

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func InterceptorGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Infof("method: %v, req: %v\n", info.FullMethod, req)
	resp, err = handler(ctx, req)
	log.Infof("resp: %v, err: %v\n", resp, err)
	return resp, err
}

func InterceptorHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("method: %v, url: %v\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
