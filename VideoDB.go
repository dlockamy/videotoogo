package videotogo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type VideoDB struct {
	dbFilePath   string
	videoCatalog []Video
}

func NewDB() *VideoDB {
	return &VideoDB{}
}

func LoadDB(path string) *VideoDB {
	var v = &VideoDB{dbFilePath: path}
	v.loadFromDisk()
	return v
}

func (v *VideoDB) AddVideo(video Video) {
	v.videoCatalog = append(v.videoCatalog, video)
}

func (v *VideoDB) GetAvailableVideos() []Video {
	return v.videoCatalog
}
func (v *VideoDB) SaveAs(path string) {
	v.dbFilePath = path
	v.saveToDisk()
}

func (v *VideoDB) loadFromDisk() {
	data, err := ioutil.ReadFile(v.dbFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &v.videoCatalog)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (v *VideoDB) saveToDisk() {
	videoCatalogJson, _ := json.Marshal(&v.videoCatalog)
	ioutil.WriteFile(v.dbFilePath, videoCatalogJson, 0644)
}
