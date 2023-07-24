package services

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"s16-tech-test/generated"
)

func TestExtractGetMovieByIDResponse(t *testing.T) {
	successBody := `
	{"Title":"Guardians of the Galaxy Vol. 2","Year":"2017","Rated":"PG-13","Released":"05 May 2017","Runtime":"136 min","Genre":"Action, Adventure, Comedy","Director":"James Gunn","Writer":"James Gunn, Dan Abnett, Andy Lanning","Actors":"Chris Pratt, Zoe Saldana, Dave Bautista","Plot":"The Guardians struggle to keep together as a team while dealing with their personal family issues, notably Star-Lord's encounter with his father, the ambitious celestial being Ego.","Language":"English","Country":"United States","Awards":"Nominated for 1 Oscar. 15 wins & 60 nominations total","Poster":"https://m.media-amazon.com/images/M/MV5BNjM0NTc0NzItM2FlYS00YzEwLWE0YmUtNTA2ZWIzODc2OTgxXkEyXkFqcGdeQXVyNTgwNzIyNzg@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"7.6/10"},{"Source":"Rotten Tomatoes","Value":"85%"},{"Source":"Metacritic","Value":"67/100"}],"Metascore":"67","imdbRating":"7.6","imdbVotes":"719,971","imdbID":"tt3896198","Type":"movie","DVD":"22 Aug 2017","BoxOffice":"$389,813,101","Production":"N/A","Website":"N/A","Response":"True"}
	`
	type args struct {
		resp *http.Response
		req  *generated.GetMovieByIDRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *generated.GetMovieByIDResponse
		wantErr bool
	}{
		{
			name: "Success Case",
			args: args{
				resp: &http.Response{
					Body: io.NopCloser(strings.NewReader(successBody)),
				},
				req: &generated.GetMovieByIDRequest{
					Id: "tt3896198",
				},
			},
			want: &generated.GetMovieByIDResponse{
				Id:       "tt3896198",
				Title:    "Guardians of the Galaxy Vol. 2",
				Year:     "2017",
				Rated:    "PG-13",
				Genre:    "Action, Adventure, Comedy",
				Plot:     "The Guardians struggle to keep together as a team while dealing with their personal family issues, notably Star-Lord's encounter with his father, the ambitious celestial being Ego.",
				Director: "James Gunn",
				Actors: []string{
					"Chris Pratt", "Zoe Saldana", "Dave Bautista",
				},
				Language:  "English",
				Country:   "United States",
				Type:      "movie",
				PosterUrl: "https://m.media-amazon.com/images/M/MV5BNjM0NTc0NzItM2FlYS00YzEwLWE0YmUtNTA2ZWIzODc2OTgxXkEyXkFqcGdeQXVyNTgwNzIyNzg@._V1_SX300.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractGetMovieByIDResponse(tt.args.resp, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractGetMovieByIDResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractGetMovieByIDResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractSearchMovieResponse(t *testing.T) {
	successResp := `
	{"Search":[{"Title":"Minotaur","Year":"2007","imdbID":"tt1433153","Type":"movie"}],"totalResults":"1","Response":"True"}
	`
	incorrectResp := `
	{"result":[{"Title":"Minotaur","Year":"2007","imdbID":"tt1433153","Type":"movie"}],"totalResults":"1","Response":"True"}
	`

	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *generated.SearchMoviesResponse
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				resp: &http.Response{
					Body: io.NopCloser(strings.NewReader(successResp)),
				},
			},
			want: &generated.SearchMoviesResponse{
				Movies: []*generated.MovieResult{{
					Id:    "tt1433153",
					Title: "Minotaur",
					Year:  "2007",
					Type:  "movie",
				}},
				TotalResults: 1,
			},
			wantErr: false,
		},
		{
			name: "negative case: invalid response",
			args: args{
				resp: &http.Response{
					Body: io.NopCloser(strings.NewReader(incorrectResp)),
				},
			},
			want: &generated.SearchMoviesResponse{
				TotalResults: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractSearchMovieResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractSearchMovieResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractSearchMovieResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
