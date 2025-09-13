package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Movie struct {
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	VoteAverage float64 `json:"vote_average"`
	PosterPath  string  `json:"poster_path"`
}

type SearchResponse struct {
	Results []Movie `json:"results"`
}

func SearchMovie(apiKey, query string) ([]Movie, error) {
	endpoint := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s", apiKey, url.QueryEscape(query))

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Results, nil
}
