package validator

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
)

func ValidateFile(file *multipart.FileHeader, maxSize int64, allowedExts []string) error {
	// Get the file extension
	ext := filepath.Ext(file.Filename)
	// Check if the extension is allowed
	allowed := false
	for _, e := range allowedExts {
		if e == ext {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Errorf("invalid file extension, only %v are allowed", allowedExts)
	}
	// Get the file size
	size := file.Size
	// Check if the file size is within the allowed limit
	if size > maxSize {
		return fmt.Errorf("file size exceeds limit of %d bytes", maxSize)
	}
	return nil
}
