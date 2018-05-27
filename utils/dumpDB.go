package main

import (
	"fmt"

	"github.com/dlockamy/videotogo"
)

func main() {
	var videoDb = videotogo.LoadDB("var/data/videos.json")

	fmt.Printf("%q\n", videoDb.GetAvailableVideos())
}
