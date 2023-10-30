package db

import (
	"encoding/json"
	"io"
	"os"
)

// Files
var (
	engines = []Engine{}
)

func init() {
	// Load JSON Data
	file, err := os.Open("db/versions.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &engines)
	if err != nil {
		panic(err)
	}
}
