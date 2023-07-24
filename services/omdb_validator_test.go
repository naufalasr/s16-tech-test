package services

import (
	"errors"
	"testing"

	"s16-tech-test/generated"

	"github.com/stretchr/testify/assert"
)

func TestValidateGetMovieByIDRequest(t *testing.T) {
	type args struct {
		req *generated.GetMovieByIDRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "positive case",
			args: args{
				req: &generated.GetMovieByIDRequest{
					Id: "tt0294568",
				},
			},
			wantErr: false,
		},
		{
			name: "negative case: empty id",
			args: args{
				req: &generated.GetMovieByIDRequest{
					Id: "",
				},
			},
			wantErr: true,
			err:     errors.New("id is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGetMovieByIDRequest(tt.args.req)
			if tt.wantErr {
				assert.Equal(t, err, tt.err)
			} else {
				assert.Empty(t, err)
			}
		})
	}
}

func TestValidateSearchMovieRequest(t *testing.T) {
	type args struct {
		req *generated.SearchMoviesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "positive case",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "movie a",
					Type:  "movie",
					Page:  1,
				},
			},
			wantErr: false,
		},
		{
			name: "negative case: query is empty",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "",
					Type:  "series",
					Page:  1,
				},
			},
			err:     errors.New("query cannot be empty"),
			wantErr: true,
		},
		{
			name: "negative case: page is invalid",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "movie a",
					Type:  "movie",
					Page:  0,
				},
			},
			wantErr: true,
			err:     errors.New("invalid page. valid page is 1 - 100"),
		},
		{
			name: "negative case: type is invalid",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "documentary a",
					Type:  "documentary",
					Page:  1,
				},
			},
			wantErr: true,
			err:     errors.New("invalid types. valid types: movie, series, episode"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSearchMovieRequest(tt.args.req)
			if tt.wantErr {
				assert.Equal(t, err, tt.err)
			} else {
				assert.Empty(t, err)
			}
		})
	}
}
