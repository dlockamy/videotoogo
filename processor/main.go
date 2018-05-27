package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"

	"gopkg.in/h2non/filetype.v1"
)

func main() {
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
					log.Println("modified file:", event.Name)
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

	return fmt.Sprintf("%x", h.Sum(nil)) //string(h.Sum(nil))
}

func processNewFile(path string) {
	file, _ := os.Open(path)

	head := make([]byte, 261)
	file.Read(head)

	if filetype.IsVideo(head) {

		newFileName := "./var/blocks/" + generateHashFileName(path)

		log.Printf("New File name = " + newFileName)

		os.Rename(path, newFileName)
	} else {
		log.Println("Not an video")
		os.Remove(path)
	}
}
