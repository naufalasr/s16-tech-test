package interceptors

import (
	"context"

	con "s16-tech-test/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// acts as a middleware in unary call
func AuthMwUnary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if md["api_key"] != nil {
				key := md["api_key"][0]
				ctx = context.WithValue(ctx, con.APIKey, key)
			}
		}
		return handler(ctx, req)
	}
}
