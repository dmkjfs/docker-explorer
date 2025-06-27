package service

import (
	"os"
	"path/filepath"
)

type FileService struct{}

func (fm *FileService) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fm *FileService) Write(path string, content []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, content, 0644)
}

func (fm *FileService) Delete(path string) error {
	return os.Remove(path)
}

func (fm *FileService) List(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var items []string
	for _, entry := range entries {
		items = append(items, entry.Name())
	}
	return items, nil
}
