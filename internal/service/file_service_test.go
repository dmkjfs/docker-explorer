package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	fm := &FileService{}

	err := fm.Write("file.txt", []byte("helloworld"))
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	fm := &FileService{}

	content, err := fm.Read("file.txt")
	assert.NoError(t, err)
	assert.Equal(t, "helloworld", string(content))

	_, err = fm.Read(filepath.Join("file.txt", "nonexistent.txt"))
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	fm := &FileService{}

	err := fm.Delete("file.txt")
	assert.NoError(t, err)

	err = fm.Delete("nonexistent.txt")
	assert.Error(t, err)
}

func TestList(t *testing.T) {
	fm := &FileService{}
	testDir := t.TempDir()
	expectedFiles := []string{"test1.txt", "test2.txt", "test3.txt"}

	for _, fname := range expectedFiles {
		err := os.WriteFile(filepath.Join(testDir, fname), []byte("test"), 0644)
		assert.NoError(t, err, "Failed to create test file %s: %v", fname, err)
	}

	items, err := fm.List(testDir)
	assert.NoError(t, err)
	assert.Equal(t, len(expectedFiles), len(items))
}
