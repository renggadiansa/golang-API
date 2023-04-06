package utils

import (
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var carset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
func RandomString(n int) string {
	//rand.Seed(time.Now().UnixMilli())
	b := make([]byte, n)
	for i := range b {
		b[i] = carset[rand.Intn(len(carset))]
	}
	return string(b)
}

func FileValidationByExtension(fileHeader *multipart.FileHeader, fileExtension []string) bool {
	extension := filepath.Ext(fileHeader.Filename)
	log.Println( "extension",extension)
	result := false

	for _, typeFile := range fileExtension {
		if extension == typeFile {
			result = true
			break
		}
	}

	return result
}

func RandomFileName(extensionFile string, prefix ...string) string {

	currentPrefix := "file"

	if len(prefix) > 0 {

		if prefix[0] != "" {
			currentPrefix= prefix[0]
		}

	}

	currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")
	fileName := fmt.Sprintf("%s-%s-%s%s",currentPrefix ,currentTime , RandomString(5), extensionFile)

	return fileName

}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, fileName string) bool {
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", fileName))

	if errUpload != nil {
		log.Println("can't upload file")
		return false
	} else {
		return true
	}
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("error remove file")
		return err
	}
	return nil
}