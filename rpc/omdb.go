package rpc

import (
	"context"
	"errors"

	con "s16-tech-test/context"
	"s16-tech-test/generated"
	"s16-tech-test/services"

	"google.golang.org/grpc"
)

type OmdbServer struct {
	generated.UnimplementedOMDBServiceServer
}

func NewOMDBService(srv *grpc.Server) {
	generated.RegisterOMDBServiceServer(srv, &OmdbServer{})
}

func (o *OmdbServer) GetMovieByID(ctx context.Context, req *generated.GetMovieByIDRequest) (*generated.GetMovieByIDResponse, error) {
	if !con.IsAuthenticated(ctx) {
		return nil, errors.New("unauthenticated")
	}
	return services.GetMovieByID(ctx, req)
}

func (o *OmdbServer) SearchMovies(ctx context.Context, req *generated.SearchMoviesRequest) (*generated.SearchMoviesResponse, error) {
	if !con.IsAuthenticated(ctx) {
		return nil, errors.New("unauthenticated")
	}
	return services.SearchMovies(ctx, req)
}
