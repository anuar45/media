package main

import (
	"encoding/json"
	"errors"
	"flag"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

const (
	mediaPrefix  = "/media/"
	thumbsPrefix = "/thumbs/"
	webPrefix    = "/public/"
	webDir       = "./svelte/public/"
)

// http://mediaserver/api/v1/items?path=/
// http://mediaserver/media/

// http://mediaserver/api/v1/items?path=/video
// http://mediaserver/media/video
type Config struct {
	mediaDir   string
	listenAddr string
}

func (c *Config) Parse() error {
	flag.StringVar(&c.mediaDir, "mediaDir", "", "path to local directory to serve")
	flag.StringVar(&c.listenAddr, "listenAddr", ":8080", "http listen addr")
	flag.Parse()

	return c.Validate()
}

func (c *Config) Validate() error {
	if c.mediaDir == "" {
		return errors.New("missing required flag --mediaDir")
	}

	return nil
}

type Server struct {
	config *Config
}

func main() {
	cfg := &Config{}

	err := cfg.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	srv := &Server{config: cfg}

	log.Println("Starting server on port", cfg.listenAddr)
	log.Fatalln(srv.Run())
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", s.HomeHandler)
	router.HandleFunc("/api/v1/items", s.ItemsHandler)
	router.HandleFunc(thumbsPrefix+"{path:.*}", s.ThumbsHandler)
	router.PathPrefix(mediaPrefix).Handler(http.StripPrefix(mediaPrefix, http.FileServer(http.Dir(s.config.mediaDir))))
	router.PathPrefix(webPrefix).Handler(http.StripPrefix(webPrefix, http.FileServer(http.Dir(webDir))))

	return http.ListenAndServe(s.config.listenAddr, router)
}

type MediaItem struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Type      string `json:"type"`
	ThumbPath string `json:"thumbpath"`
}

func (s *Server) ItemsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	mediaItems := s.filesToItems(vars["path"])

	json.NewEncoder(w).Encode(mediaItems)
}

func (s *Server) ThumbsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	file, err := os.Open(s.config.mediaDir + "/" + vars["path"])
	if err != nil {
		log.Printf("error getting thum for %s\n", s.config.mediaDir+vars["path"])
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Printf("error getting thumb for %s\n", s.config.mediaDir+vars["path"])
		return
	}

	m := resize.Resize(320, 0, img, resize.Lanczos3)

	jpeg.Encode(w, m, nil)
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "svelte/public/index.html")
}

func (s *Server) filesToItems(path string) []MediaItem {
	var mediaItems []MediaItem

	targetDir := s.config.mediaDir + path
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		log.Printf("error reading folder %s: %v\n", targetDir, err)
	}

	for _, file := range files {

		fileType := getFileType(file)

		var mediaPath, thumbPath string
		switch fileType {
		case "folder":
			mediaPath = "./svelte/public/folder.png"
			thumbPath = "./svelte/public/folder.png"
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
		return "folder"
	}

	return "media"
}
