package media

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	config *Config
}

func NewService(config *Config) *Service {
	return &Service{config}
}

func (s *Service) FilesToItems(base, path string) []MediaItem {
	var mediaItems []MediaItem

	if base == "" {
		for virtualDir := range s.config.MediaDirs {

			mediaItems = append(
				mediaItems,
				MediaItem{
					Name: virtualDir,
					Path: "/",
					Type: "directory",
					Base: virtualDir,
				},
			)
		}

		return mediaItems
	}

	// /media/pics/file.jpg
	// /media/pics2/video/file.mkv

	log.Println(base, path)
	items, err := ioutil.ReadDir(s.config.MediaDirs[base] + path)
	if err != nil {
		log.Printf("error reading directory %s: %v\n", base, err)
	}

	for _, item := range items {

		itemType := getItemType(item)

		var mediaPath string
		switch itemType {
		case "directory":
			if path == "/" {
				mediaPath = "/" + item.Name()
			} else {
				mediaPath = strings.Join([]string{path, item.Name()}, "/")
			}
		default:
			if path == "/" {
				mediaPath = mediaPrefix + base + "/" + item.Name()
			} else {
				mediaPath = mediaPrefix + base + path + "/" + item.Name()
			}
		}

		mediaItem := MediaItem{
			item.Name(),
			mediaPath,
			itemType,
			base,
		}

		mediaItems = append(mediaItems, mediaItem)
	}

	return mediaItems
}

var fileCategory = map[string]string{
	".mp4": "video",
	".jpg": "image",
	".png": "image",
}

func getItemType(fileInfo os.FileInfo) string {
	if fileInfo.IsDir() {
		return "directory"
	}

	fileType, ok := fileCategory[filepath.Ext(fileInfo.Name())]
	if !ok {
		return "unknown"
	}

	return fileType
}
