package api

import (
	"encoding/json"
	"net/http"

	"github.com/anuar45/media/internal/service/media"
	"github.com/gorilla/mux"
)

const (
	webPrefix = "/public/"
	webDir    = "./front/public/"
)

type Server struct {
	config *media.Config
	svc    *media.Service
}

func NewServer(svc *media.Service, config *media.Config) *Server {
	return &Server{
		svc:    svc,
		config: config,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", s.ServeUI)
	router.HandleFunc("/api/v1/items", s.ItemsHandler)
	// router.HandleFunc(thumbsPrefix+"{path:.*}", s.ThumbsHandler)
	router.PathPrefix(webPrefix).Handler(http.StripPrefix(webPrefix, http.FileServer(http.Dir(webDir))))

	for virtualDir, fsDir := range s.config.MediaDirs {
		router.PathPrefix(s.config.MediaPrefix + virtualDir + "/").Handler(http.StripPrefix(s.config.MediaPrefix+virtualDir+"/", http.FileServer(http.Dir(fsDir))))
	}

	return http.ListenAndServe(s.config.ListenAddr, router)
}

func (s *Server) ItemsHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Query().Get("path")
	base := r.URL.Query().Get("base")

	mediaItems := s.svc.FilesToItems(base, path)

	json.NewEncoder(w).Encode(mediaItems)
}

func (s *Server) ServeUI(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "front/public/index.html")
}
