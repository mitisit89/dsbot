package yt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const apiURL = "https://www.googleapis.com/youtube/v3/search"

type YtResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}

func SearchFilmTrailer(film string) (*YtResponse, error) {
	var ytResponse YtResponse
	query := url.Values{}
	query.Set("part", "snippet")
	query.Set("q", film+" trailer")
	query.Set("type", "video")
	query.Set("key", os.Getenv("YT_API_KEY"))

	resp, err := http.Get(apiURL + "?" + query.Encode())
	if err != nil {
		return nil, fmt.Errorf("filmTrailersSearch: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("filmTrailersSearch: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&ytResponse); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}
	return &ytResponse, nil
}
