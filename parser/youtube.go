package parser

import (
	"encoding/json"
	"errors"
	"io"
	"os/exec"
	"strings"
)

type YoutubeQuality string

const (
	QualityBest  YoutubeQuality = "bestaudio"
	QualityWorst YoutubeQuality = "worstaudio"
)

var supportedCodecs = [...]string{
	"vorbis",
	"opus",
	"mp4a",
}
var es ExtractorFormatSelector

func init() {
	es = &youtubeParser{}
}

// youtubeVideoInfo: Contains info about a Youtube video
type youtubeVideoInfo struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	youtubeFormat
}

// Extractor: Extracts info of a given video
type Extractor interface {
	Extract(string) (*youtubeResponse)
}

// FormatSelector: Selects a format based on an given YoutubeQuality
type FormatSelector interface {
	Select(*youtubeVideo, YoutubeQuality)
}

// ExtractorFormatSelector: Composes extractor and video selector
type ExtractorFormatSelector interface {
	Extractor
	FormatSelector
}

type youtubeResponse struct {
	Data  chan *youtubeVideo
	Error chan error
}

type youtubeVideo struct {
	youtubeVideoInfo
	Formats []youtubeFormat
}

type youtubeFormat struct {
	Format   string  `json:"format"`
	Ext      string  `json:"ext"`
	Acodec   string  `json:"acodec"`
	FormatId string  `json:"format_id"`
	Tbr      float64 `json:"tbr"`
	Vcodec   string  `json:"vcodec"`
}

type youtubeParser struct{}

func (youtubeParser) Extract(url string) (*youtubeResponse) {
	resp := &youtubeResponse{
		Data:  make(chan *youtubeVideo),
		Error: make(chan error),
	}

	args := []string{"--print-json", "--quiet", "-J", url}

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
			resp.Data <- s
		}
	}()

	return resp
}

func (youtubeParser) Select(video *youtubeVideo, quality YoutubeQuality) {
	var match *youtubeFormat
formatLoop:
	for index, format := range video.Formats {
		if format.Vcodec != "none" {
			continue formatLoop
		}
		for _, sup := range supportedCodecs {
			if strings.HasPrefix(format.Acodec, sup) {
				if match == nil {
					match = &video.Formats[index]
					continue formatLoop
				} else if quality == QualityBest && format.Tbr > match.Tbr {
					match = &video.Formats[index]
					continue formatLoop
				} else if quality == QualityWorst && format.Tbr < match.Tbr {
					match = &video.Formats[index]
					continue formatLoop
				}
			}
		}
	}
	video.youtubeFormat = *match
}

func Parse(url string, quality YoutubeQuality) (*youtubeVideoInfo, error) {
	if strings.TrimSpace(url) == "" {
		return nil, errors.New("url can't be empty")
	}

	res := es.Extract(url)
	select {
	case err := <-res.Error:
		return nil, err
	case d := <-res.Data:
		es.Select(d, quality)
		return &d.youtubeVideoInfo, nil
	}

	return nil, errors.New("error parsing video")
}
