package filecontroller

import (
	"fmt"
	// "path/filepath"

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

	// file, errFile := fileHeader.Open()

	// if errFile != nil {
	// 	panic(errFile)
	// }

	// extensionFile := filepath.Ext(fileHeader.Filename)
	// fileName := 

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", fileHeader.Filename))

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