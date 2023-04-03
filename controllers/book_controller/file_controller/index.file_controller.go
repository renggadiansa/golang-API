package filecontroller

import (
	"fmt"
	"gin-goinc-api/utils"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleUploadFile(ctx *gin.Context) {

	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message" : "file is required",
		})
		return
	}


	extensionFile := filepath.Ext(fileHeader.Filename)

	currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")

	fileName := fmt.Sprintf("%s-%s%s", currentTime , utils.RandomString(5), extensionFile)

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", fileName))

	if errUpload != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "file uploaded",
	})
}