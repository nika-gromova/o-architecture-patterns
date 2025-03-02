package logging

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Infof("method: %v, req: %v\n", info.FullMethod, req)
	resp, err = handler(ctx, req)
	log.Infof("resp: %v, err: %v\n", resp, err)
	return resp, err
}
