package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/dlockamy/videotogo"
)

var videoDb videotogo.VideoDB
var dbPath = "./var/data/videos.json"

const BUFSIZE = 1024 * 8

func main() {
	if _, err := os.Stat(dbPath); err != nil {
		log.Println("Error loading Videos DB, unable to continue")
		return
	}

	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./www/")))
	r.Handle("/block", http.FileServer(http.Dir("./data/")))
	r.HandleFunc("/list", listStreams)
	r.HandleFunc("/stream/{video-id}", streamVideo)
	r.HandleFunc("/stream/{session-id}/ctl", streamCtl)

	http.ListenAndServe(":3002", handlers.LoggingHandler(os.Stdout, r))
}

func invalidVideoID(videoID string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	//TODO: define a standard error object and return it as JSON
	w.Write([]byte("Video " + videoID + "Item Not Found"))
}

func listStreams(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	videoDb = *videotogo.LoadDB(dbPath)
	payload, _ := json.Marshal(videoDb.GetAvailableVideos())
	w.Write([]byte(payload))
}

func streamVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["video-id"]

	var filePath = "./var/blocks/" + slug

	log.Printf("loading " + filePath)

	if _, err := os.Stat(filePath); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Product Not Found"))
		payload, _ := json.Marshal(slug)
		w.Write([]byte(payload))
	}

	file, err := os.Open(filePath)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	defer file.Close()

	fi, err := file.Stat()

	if err != nil {
		w.WriteHeader(500)
		return
	}

	fileSize := int(fi.Size())

	contentLength := strconv.Itoa(fileSize)
	contentEnd := strconv.Itoa(fileSize - 1)

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Content-Length", contentLength)
	w.Header().Set("Content-Range", "bytes 0-"+contentEnd+"/"+contentLength)
	w.WriteHeader(200)

	buffer := make([]byte, BUFSIZE)

	for {
		n, err := file.Read(buffer)

		if n == 0 {
			break
		}

		if err != nil {
			break
		}

		data := buffer[:n]
		w.Write(data)
		w.(http.Flusher).Flush()
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
