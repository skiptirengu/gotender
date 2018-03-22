package parser

import (
	"errors"
	"os/exec"
	"encoding/json"
	"io"
)

type YoutubeQuality string

const (
	QualityBest  YoutubeQuality = "bestaudio"
	QualityWorst YoutubeQuality = "worstaudio"
)

var supportedCodecs = [...]string{
	"vorbis",
	"opus",
	"mp4a.40.2",
}

type YoutubeResponse struct {
	Data  chan *YoutubeVideoInfo
	Error chan error
}

type YoutubeVideoInfo struct {
	Id        string `json:"id"`
	FullTitle string `json:"fulltitle"`
	Duration  int    `json:"duration"`
	Url       string `json:"url"`
	FormatId  string `json:"format_id"`
}

type youtubeVideo struct {
	YoutubeVideoInfo
	Formats []youtubeFormat
}

type youtubeFormat struct {
	Format   string `json:"format"`
	Ext      string `json:"ext"`
	Acodec   string `json:"acodec"`
	FormatId string `json:"format_id"`
}

func Parse(url string, quality YoutubeQuality) (*YoutubeResponse, error) {
	if url == "" {
		return nil, errors.New("url can't be empty")
	}

	return invokeYoutbedl(url, quality), nil
}

func invokeYoutbedl(url string, quality YoutubeQuality) (*YoutubeResponse) {
	resp := &YoutubeResponse{
		Data:  make(chan *YoutubeVideoInfo),
		Error: make(chan error),
	}

	args := []string{"--print-json", "--format", string(quality), url}

	go func() {
		var (
			reader io.ReadCloser
			err    error
		)

		cmd := exec.Command("youtube-dl", args...)

		if reader, err = cmd.StdoutPipe(); err != nil {
			resp.Error <- err
			return
		}

		s := &youtubeVideo{}
		if err := json.NewDecoder(reader).Decode(&s); err != nil {
			resp.Error <- err
		} else {
			resp.Data <- &s.YoutubeVideoInfo
		}
	}()

	return resp
}
