package services

import (
	"context"
	"fmt"
	"net/http"

	"s16-tech-test/config"
	"s16-tech-test/generated"
)

func GetMovieByID(ctx context.Context, req *generated.GetMovieByIDRequest) (*generated.GetMovieByIDResponse, error) {
	err := ValidateGetMovieByIDRequest(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&i=%s", config.ConfigInUse.Omdb.APIKey, req.Id)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ExtractGetMovieByIDResponse(resp, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func SearchMovies(ctx context.Context, req *generated.SearchMoviesRequest) (*generated.SearchMoviesResponse, error) {
	err := ValidateSearchMovieRequest(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&s=%s&page=%d", config.ConfigInUse.Omdb.APIKey, req.Query, req.Page)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ExtractSearchMovieResponse(resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}
