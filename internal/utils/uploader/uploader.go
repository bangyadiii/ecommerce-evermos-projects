package uploader

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(fileHeader *multipart.FileHeader, dst string) error {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new file
	newFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Copy the file data to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		return err
	}

	return nil
}
