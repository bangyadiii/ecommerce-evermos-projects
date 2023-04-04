package storage

import (
	"context"
	"ecommerce-evermos-projects/internal/helper"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"path/filepath"

	"cloud.google.com/go/storage"
	gcss "cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const currentfilepath = "internal/infrastructure/storage/storage.go"

type Storage interface {
	UploadFile(fh *multipart.FileHeader, folderName string) (string, error)
	DeleteFile(fileUrl string) error
}

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) Storage {
	return &LocalStorage{
		BasePath: basePath,
	}
}

func (ls *LocalStorage) UploadFile(fh *multipart.FileHeader, folderName string) (string, error) {
	// Generate unique filename
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), fh.Filename)

	// Create folder if it doesn't exist
	folderPath := filepath.Join(ls.BasePath, folderName)

	if err := createFolderIfNotExist(folderPath); err != nil {
		return "", err
	}

	// Create file
	filePath := filepath.Join(folderPath, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Open uploaded file
	uploadedFile, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	// Copy uploaded file to created file
	_, err = io.Copy(file, uploadedFile)
	if err != nil {
		return "", err
	}

	// Return file URL
	return filepath.Join(folderName, filename), nil
}

func createFolderIfNotExist(folderPath string) error {
	// Create folder if it doesn't exist
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (ls *LocalStorage) DeleteFile(fileUrl string) error {
	filePath := filepath.Join(ls.BasePath, fileUrl)
	return os.Remove(filePath)
}

type CloudStorage struct {
	Client    *gcss.Client
	Bucket    string
	BasePath  string
	PublicURL string
}

func NewCloudStorage(bq, basePath, publicUrl string) Storage {
	client, err := createClient()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprint("Error when instance new Storage", err.Error()))
	}

	return &CloudStorage{
		Client:    client,
		BasePath:  basePath,
		Bucket:    bq,
		PublicURL: publicUrl,
	}
}

func createClient() (client *gcss.Client, err error) {
	ctx := context.Background()
	client, err = storage.NewClient(ctx, option.WithCredentialsFile("/path/to/your/keyfile.json"))
	if err != nil {
		return client, err
	}
	return client, nil
}

func (cs *CloudStorage) UploadFile(fh *multipart.FileHeader, folderName string) (string, error) {
	// Generate unique filename
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), fh.Filename)

	// Create object handle
	object := cs.Client.Bucket(cs.Bucket).Object(filepath.Join(cs.BasePath, folderName, filename))

	// Open uploaded file
	uploadedFile, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	// Upload file to cloud storage
	wc := object.NewWriter(context.Background())
	if _, err := io.Copy(wc, uploadedFile); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	// Set public access to uploaded file
	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	// Return file URL
	return fmt.Sprintf("%s/%s/%s/%s", cs.PublicURL, cs.Bucket, cs.BasePath, filepath.Join(folderName, filename)), nil
}

func (cs *CloudStorage) DeleteFile(fileUrl string) error {
	object := cs.Client.Bucket(cs.Bucket).Object(filepath.Join(cs.BasePath, fileUrl))
	return object.Delete(context.Background())
}
