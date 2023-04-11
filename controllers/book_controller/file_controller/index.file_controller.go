package filecontroller

import (
	"fmt"
	"gin-goinc-api/constants"
	"gin-goinc-api/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SendStatus(ctx *gin.Context) {
	filename := ctx.MustGet("filename").(string)
	 
	ctx.JSON(200, gin.H{
		"message": "file uploades",
		"file_name": filename,
	})
}

func HandleUploadFile(ctx *gin.Context) {

	claimsData := ctx.MustGet("claimsData").(jwt.MapClaims)
	fmt.Println("claimsData => email =>", claimsData["email"])

	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userId =>", userId)
	

	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message" : "file is required",
		})
		return
	}


	//validasi by ektensi
	fileExtension := []string{
		".jpg", ".png", ".pdf", ".jpeg",
	}

	isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtension)

	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H {
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

	ctx.JSON(200, gin.H{
		"message": "file uploaded",
	})
}

func HandleRemoveFile(ctx *gin.Context) {
	filename := ctx.Param("filename")

	if filename == "" {
		ctx.JSON(400, gin.H{
			"message" : "filename is required",
		})
	}
	
	err := utils.RemoveFile(constants.DIR_FILE + filename)
	if err != nil {
		ctx.JSON(500, gin.H {
			"message" : err.Error(),
		})
		return
	}
	
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "file removed",
	})
}