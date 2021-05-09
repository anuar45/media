package main

import (
	"encoding/json"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

var mediaDir = "/home/anuar/media"
var mediaPrefix = "/media/"
var thumbsPrefix = "/thumbs/"
var assetsPrefix = "/assets/"
var assetsDir = "./assets"

// http://mediaserver/api/v/1/items?path=/
// http://mediaserver/media/

// http://mediaserver/api/v/1/items?path=/video
// http://mediaserver/media/video

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/items", ItemsHandler)
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc(thumbsPrefix+"{path:.*}", ThumbsHandler)
	router.PathPrefix(mediaPrefix).Handler(http.StripPrefix(mediaPrefix, http.FileServer(http.Dir(mediaDir))))
	router.PathPrefix(assetsPrefix).Handler(http.StripPrefix(assetsPrefix, http.FileServer(http.Dir(assetsDir))))

	log.Fatal(http.ListenAndServe(":8080", router))
}

type MediaItem struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Type      string `json:"type"`
	ThumbPath string `json:"thumbpath"`
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	mediaItems := filesToItems(vars["path"])

	json.NewEncoder(w).Encode(mediaItems)
}

func ThumbsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	file, err := os.Open(mediaDir + "/" + vars["path"])
	if err != nil {
		log.Printf("error getting thum for %s\n", mediaDir+vars["path"])
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Printf("error getting thumb for %s\n", mediaDir+vars["path"])
		return
	}

	m := resize.Resize(320, 0, img, resize.Lanczos3)

	jpeg.Encode(w, m, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func filesToItems(path string) []MediaItem {
	var mediaItems []MediaItem

	targetDir := mediaDir + path
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		log.Printf("error reading folder %s: %v\n", targetDir, err)
	}

	for _, file := range files {

		fileType := getFileType(file)

		var mediaPath, thumbPath string
		switch fileType {
		case "dir":
			mediaPath = "./assets/dir.png"
			thumbPath = "./assets/dir.png"
		default:
			mediaPath = mediaPrefix + path + file.Name()
			thumbPath = thumbsPrefix + path + file.Name()
		}

		mediaItem := MediaItem{
			file.Name(),
			mediaPath,
			fileType,
			thumbPath,
		}

		mediaItems = append(mediaItems, mediaItem)
	}

	return mediaItems
}

func getFileType(fileInfo os.FileInfo) string {
	if fileInfo.IsDir() {
		return "dir"
	}

	return "media"
}
