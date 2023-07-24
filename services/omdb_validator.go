package services

import (
	"errors"
	"strings"

	"s16-tech-test/generated"
)

func ValidateGetMovieByIDRequest(req *generated.GetMovieByIDRequest) error {
	if len(req.Id) < 1 {
		return errors.New("id is required")
	}

	return nil
}

func ValidateSearchMovieRequest(req *generated.SearchMoviesRequest) error {
	if len(req.Query) < 1 {
		return errors.New("query cannot be empty")
	}

	if req.Page < 1 || req.Page > 100 {
		return errors.New("invalid page. valid page is 1 - 100")
	}

	switch strings.ToLower(req.Type) {
	case "movie", "series", "episode":
	default:
		return errors.New("invalid types. valid types: movie, series, episode")
	}

	return nil
}
