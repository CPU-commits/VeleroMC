package db

import "errors"

// Model
type Engine struct {
	Type     string   `json:"type"`
	URL      string   `json:"url"`
	Versions []string `json:"versions"`
	Name     string   `json:"name"`
}

type enginesRepository struct{}

func (*enginesRepository) FindOneByEngine(engine string) (*Engine, error) {
	for _, version := range engines {
		if version.Name == engine {
			return &version, nil
		}
	}

	return nil, errors.New("Engine not found")
}

func (*enginesRepository) FindAll() ([]Engine, error) {
	return engines, nil
}

func NewEnginesRepository() *enginesRepository {
	return &enginesRepository{}
}
