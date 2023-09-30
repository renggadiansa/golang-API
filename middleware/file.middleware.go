package middleware

import (
	"gin-goinc-api/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is required",
		})
		return
	}

	//validasi by ektensi
	fileExtension := []string{
		".jpg", ".png", ".pdf", ".jpeg",
	}

	isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtension)

	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file type is not valid",
		})
		return
	}

	//validasi by content type
	// fileType := []string{
	// 	"image/jpg",
	// }

	// isFileValidated := utils.FileValidation(fileHeader, fileType)

	// if !isFileValidated {
	// 	ctx.AbortWithStatusJSON(400, gin.H {
	// 		"message": "file type is not valid",
	// 	})
	// 	return
	// }

	extensionFile := filepath.Ext(fileHeader.Filename)

	fileName := utils.RandomFileName(extensionFile)

	isSaved := utils.SaveFile(ctx, fileHeader, fileName)

	if !isSaved {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.Set("filename", fileName)

	ctx.Next()
}
