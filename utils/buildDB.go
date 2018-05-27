package main

import (
	"github.com/dlockamy/videotogo"
)

func main() {
	var videoDb = *videotogo.NewDB()

	videoDb.AddVideo(videotogo.Video{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Do cool things on your hover board and laser cannon"})

	videoDb.SaveAs("var/data/videos.json")
}
