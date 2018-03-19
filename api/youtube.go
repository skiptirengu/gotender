package api

import (
	"net/http"
	"github.com/skiptirengu/gotender/config"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type YoutubeSearch struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

func SearchYoutubeVideos(q string) ([]YoutubeSearch, error) {
	apiKey := config.Get().YoutubeApiKey
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}

	// No way to get the duration without a second request
	call := service.Search.List("id,snippet").Type("video").Q(q).MaxResults(30)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	videos := make([]YoutubeSearch, len(response.Items))

	for i, item := range response.Items {
		videos[i] = YoutubeSearch{
			Id:          item.Id.VideoId,
			Title:       item.Snippet.Title,
			Thumbnail:   item.Snippet.Thumbnails.Medium.Url,
			Description: item.Snippet.Description,
		}
	}

	return videos, nil
}
