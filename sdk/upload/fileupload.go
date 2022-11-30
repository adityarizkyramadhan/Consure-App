package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	storage_go "github.com/supabase-community/storage-go"
)

const (
	MaxFileSize = 2 * 1024 * 1024
)

func UploadImage(fileInput *multipart.FileHeader) (string, error) {
	if fileInput.Size > MaxFileSize {
		return "", errors.New("file size is too big")
	}
	file, err := fileInput.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileName := fmt.Sprintf("data%s%s", time.Now().Format("20060102150405"), fileInput.Filename)
	client := storage_go.NewClient("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1", os.Getenv("JWT_SUPABASE"), nil)
	client.UploadFile("foto-proker", fileName, file)
	linkImage := GenerateLinkImage(fileName)
	return linkImage, nil
}

func GenerateLinkImage(fileName string) string {
	return fmt.Sprintf("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1/object/public/foto-proker/%s", fileName)
}
