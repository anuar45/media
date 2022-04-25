package media

import (
	"errors"
	"flag"
	"path/filepath"
)

type mediaDirs map[string]string

const (
	mediaPrefix  = "/media/"
	thumbsPrefix = "/thumbs/"
	webPrefix    = "/public/"
	webDir       = "./front/public/"
)

type Config struct {
	MediaDirs    mediaDirs
	WebPrefix    string
	WebDir       string
	MediaPrefix  string
	ThumbsPrefix string
	ListenAddr   string
}

func ConfigFromFlags() (*Config, error) {
	config := Config{
		WebPrefix:    webPrefix,
		WebDir:       webDir,
		MediaPrefix:  mediaPrefix,
		ThumbsPrefix: thumbsPrefix,
	}

	err := config.Parse()

	return config, err
}

func (m mediaDirs) String() string {
	return "path to media dir"
}

func (m mediaDirs) Set(value string) error {
	if m == nil {
		m = make(map[string]string)
	}

	mediaDirBase := filepath.Base(value)

	if _, ok := m[mediaDirBase]; ok {
		for i, ok := 1, true; ok; i++ {
			mediaDirBaseNew := mediaDirBase + string(i)
			if _, ok = m[mediaDirBaseNew]; !ok {
				m[mediaDirBase] = value
				break
			}
		}
	} else {
		m[mediaDirBase] = value
	}

	return nil
}

func (c *Config) Parse() error {
	flag.Var(&c.MediaDirs, "mediaDir", "paths to media directory to serve")
	flag.StringVar(&c.ListenAddr, "listenAddr", ":8080", "http listen addr")
	flag.Parse()

	return c.Validate()
}

func (c *Config) Validate() error {
	if len(c.MediaDirs) == 0 {
		return errors.New("missing required flag --mediaDir")
	}

	return nil
}
