package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadJSONFile(filePath string, v interface{}) error {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, v)
}

func WriteJSONFile(filePath string, v interface{}) error {
	fileData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, fileData, 0644)
}

func EnsureDir(dirPath string) error {
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		return os.MkdirAll(dirPath, 0755)
	}
	return nil
}

func GetFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}