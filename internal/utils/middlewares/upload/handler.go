package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"slices"
)

func Ext(f *multipart.FileHeader) string {
	return filepath.Ext(f.Filename)
}

func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	// Check for file size
	if file.Size > MAX_FILE_SIZE {
		return nil, fmt.Errorf("file size exceeded %d bytes for %s", MAX_FILE_SIZE, file.Filename)
	}

	// Check for file extension
	if !slices.Contains(ALLOWED_EXTENSIONS, Ext(file)) {
		return nil, fmt.Errorf("file type not allowed for %s", file.Filename)
	}

	f, err := file.Open()
	if err != nil {
		fmt.Printf("open %s error: %s\n", file.Filename, err.Error())
		return nil, err
	}
	return io.ReadAll(f)
}
