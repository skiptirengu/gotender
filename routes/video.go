package routes

import (
	"net/http"
	"github.com/skiptirengu/gotender/util"
	"github.com/skiptirengu/gotender/api"
	"strings"
	"encoding/json"
	"github.com/gorilla/mux"
)

type VideoSearchRequest struct{}

func (VideoSearchRequest) HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := mux.Vars(r)["query"]
	if strings.TrimSpace(query) == "" {
		util.NewHttpErrorWithMessage(w, http.StatusBadRequest, "Query can't be empty")
		return
	}

	res, err := api.SearchYoutubeVideos(query)
	if err != nil {
		util.NewHttpError(w, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		util.NewHttpError(w, http.StatusInternalServerError)
	}
}
