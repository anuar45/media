package media

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Service struct {
	config *Config
}

func NewService(config *Config) *Service {
	return &Service{config}
}

func (s *Service) FilesToItems(path string) []MediaItem {
	var mediaItems []MediaItem

	if path == "/" {
		for virtualDir, fsDir := range s.config.MediaDirs {
			basePath := filepath.Base(fsDir)
			mediaItems = append(mediaItems, MediaItem{Name: basePath, Path: virtualDir, Type: "directory"})
		}

		return mediaItems
	}

	// /media/pics/file.jpg
	// /media/pics2/video/file.mkv

	targetDir := path
	log.Println(targetDir)
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		log.Printf("error reading directory %s: %v\n", targetDir, err)
	}

	for _, file := range files {

		fileType := getFileType(file)

		var mediaPath string
		switch fileType {
		case "directory":
			mediaPath = "/" + path + file.Name()
		default:
			mediaPath = mediaPrefix + path + file.Name()
		}

		mediaItem := MediaItem{
			file.Name(),
			mediaPath,
			fileType,
		}

		mediaItems = append(mediaItems, mediaItem)
	}

	return mediaItems
}

var fileCategory = map[string]string{
	".mp4": "video",
	".jpg": "image",
}

func getFileType(fileInfo os.FileInfo) string {
	if fileInfo.IsDir() {
		return "directory"
	}

	fileType, ok := fileCategory[filepath.Ext(fileInfo.Name())]
	if !ok {
		return "other"
	}

	return fileType
}
