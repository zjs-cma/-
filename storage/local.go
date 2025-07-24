package storage

import (
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	os.MkdirAll(basePath, 0755)
	return &LocalStorage{BasePath: basePath}
}

func (s *LocalStorage) Save(filename string, data []byte) (string, error) {
	path := filepath.Join(s.BasePath, filename)
	return path, os.WriteFile(path, data, 0644)
}

func (s *LocalStorage) Delete(filename string) error {
	targetPath := filepath.Join(s.BasePath, filename)
	return os.Remove(targetPath)
}
