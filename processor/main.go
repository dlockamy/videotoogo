package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/dlockamy/videotogo"
	"github.com/fsnotify/fsnotify"

	"gopkg.in/h2non/filetype.v1"
)

var videoDb videotogo.VideoDB
var dbPath = "./var/data/videos.json"

func main() {
	if _, err := os.Stat(dbPath); err != nil {
		log.Println("Error loading Videos DB, unable to continue")
		return
	}

	videoDb = *videotogo.LoadDB(dbPath)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					processNewFile(event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./var/uploads")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func generateHashFileName(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func processNewFile(path string) {
	file, _ := os.Open(path)

	head := make([]byte, 261)
	file.Read(head)

	if filetype.IsVideo(head) {
		var hashID = generateHashFileName(path)
		newFileName := "./var/blocks/" + hashID
		os.Rename(path, newFileName)

		videoDb.AddVideo(videotogo.Video{Id: hashID, Name: filepath.Base(path)})
	} else {
		os.Remove(path)
	}
}
