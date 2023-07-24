package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"s16-tech-test/generated"
)

func ExtractGetMovieByIDResponse(resp *http.Response, req *generated.GetMovieByIDRequest) (*generated.GetMovieByIDResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	m["id"] = req.Id

	actors, ok := m["Actors"].(string)
	if ok {
		m["actors"] = strings.Split(actors, ", ")
		delete(m, "Actors")
	}

	poster, ok := m["Poster"].(string)
	if ok {
		m["poster_url"] = poster
	}

	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var res generated.GetMovieByIDResponse
	err = json.Unmarshal(jsonData, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func ExtractSearchMovieResponse(resp *http.Response) (*generated.SearchMoviesResponse, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	movies, ok := m["Search"].([]interface{})
	if ok {
		movs := []interface{}{}
		for _, v := range movies {
			mov := v.(map[string]interface{})
			mov["id"] = mov["imdbID"]
			mov["poster_url"] = mov["Poster"]

			movs = append(movs, mov)
		}
		m["movies"] = movs
	}

	result, ok := m["totalResults"]
	if ok {
		tot := result.(string)
		m["total_results"], err = strconv.Atoi(tot)
		if err != nil {
			return nil, err
		}
	}

	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var res generated.SearchMoviesResponse
	err = json.Unmarshal(jsonData, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
