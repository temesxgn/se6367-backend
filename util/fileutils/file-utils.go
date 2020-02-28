package fileutils

import (
	"io/ioutil"
	"path/filepath"
)

// LoadFileAsString - loads the file as the given path and returns as string
func LoadFileAsString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", nil
	}

	return string(data), nil
}

// LoadEventsTodayEmailTemplate - loads the events today email template file
func LoadEventsTodayEmailTemplate() (string, error) {
	path, _ := filepath.Abs("./sendgrid/templates/my-events-today.html")
	return LoadFileAsString(path)
}

