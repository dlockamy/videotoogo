package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/dlockamy/videotogo"
)

var videoDb videotogo.VideoDB
var dbPath = "./var/data/videos.json"

func main() {
	if _, err := os.Stat(dbPath); err != nil {
		log.Println("Error loading Videos DB, unable to continue")
		return
	}

	videoDb = *videotogo.LoadDB(dbPath)

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./www/")))

	r.Handle("/block", http.FileServer(http.Dir("./data/")))

	r.HandleFunc("/list", listStreams)

	r.HandleFunc("/stream/{video-id}", streamVideo)
	r.HandleFunc("/stream/{video-id}/ctl", streamCtl)

	http.ListenAndServe(":3002", handlers.LoggingHandler(os.Stdout, r))
}

func invalidVideoID(videoID string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	//TODO: define a standard error object and return it as JSON
	//	payload, _ := json.Marshal(videoError)
	//w.Write([]byte(payload))

	w.Write([]byte("Video " + videoID + "Item Not Found"))
}

func listStreams(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(videoDb.GetAvailableVideos())
	w.Write([]byte(payload))
}

func streamVideo(w http.ResponseWriter, r *http.Request) {
	var videoItem videotogo.Video
	vars := mux.Vars(r)
	slug := vars["video-id"]

	var videoCatalog = videoDb.GetAvailableVideos()

	for _, p := range videoCatalog {
		if p.Slug == slug {
			videoItem = p
		}
	}

	if videoItem.Slug != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Product Not Found"))
		payload, _ := json.Marshal(videoItem)
		w.Write([]byte(payload))
	} else {
		invalidVideoID(videoItem.Slug, w)
	}
}

func streamCtl(w http.ResponseWriter, r *http.Request) {
	var videoItem videotogo.Video
	vars := mux.Vars(r)
	slug := vars["video-id"]

	var videoCatalog = videoDb.GetAvailableVideos()

	for _, p := range videoCatalog {
		if p.Slug == slug {
			videoItem = p
		}
	}

	if videoItem.Slug != "" {
		w.Header().Set("Content-Type", "application/json")
		payload, _ := json.Marshal(videoItem)
		w.Write([]byte(payload))
	} else {
		invalidVideoID(videoItem.Slug, w)
	}
}
