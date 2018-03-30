package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/skiptirengu/gotender/parser"
	"github.com/skiptirengu/gotender/util"
)

type StreamRequest struct{}

func (StreamRequest) HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		vars    = mux.Vars(r)
		id      = vars["id"]
		quality = vars["quality"]
	)

	if strings.TrimSpace(id) == "" {
		util.NewHttpErrorWithMessage(w, http.StatusBadRequest, "Id can't be empty")
		return
	}

	if quality != string(parser.QualityBest) && quality != string(parser.QualityWorst) {
		util.NewHttpErrorWithMessage(w, http.StatusBadRequest, "Unknown quality")
		return
	}

	data, err := parser.Parse(id, parser.YoutubeQuality(quality))
	if err != nil {
		util.NewHttpError(w, http.StatusInternalServerError)
		return
	}

	if json.NewEncoder(w).Encode(data) != nil {
		util.NewHttpError(w, http.StatusInternalServerError)
	}
}
