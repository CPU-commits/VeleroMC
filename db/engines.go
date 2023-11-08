package db

import (
	"errors"
	"strings"
)

// Model
type Engine struct {
	Type     string   `json:"type"`
	URL      string   `json:"url"`
	Versions []string `json:"versions"`
	Name     string   `json:"name"`
}

func (engine *Engine) ExistsVersion(version string) (bool, error) {
	for _, eng := range engine.Versions {
		if eng == version {
			return true, nil
		}
	}
	return false, nil
}

func (engine *Engine) BuildDownloadVersionURL(version string) string {
	return strings.ReplaceAll(engine.URL, "{version}", version)
}

// Repository
type enginesRepository struct{}

func (*enginesRepository) FindOneByEngine(engine string) (*Engine, error) {
	for _, version := range engines {
		if version.Name == engine {
			return &version, nil
		}
	}

	return nil, errors.New("Engine not found")
}

func (*enginesRepository) ExistsEngine(engine string) (bool, error) {
	for _, eng := range engines {
		if eng.Name == engine {
			return true, nil
		}
	}
	return false, nil
}

func (*enginesRepository) FindAll() ([]Engine, error) {
	return engines, nil
}

func NewEnginesRepository() *enginesRepository {
	return &enginesRepository{}
}
