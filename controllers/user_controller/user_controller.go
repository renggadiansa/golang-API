package user_controller

import (
	"gin-goinc-api/database"
	"gin-goinc-api/models"
	"gin-goinc-api/requests"
	"gin-goinc-api/responses"
	"strconv"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {

	users := new([]models.User)
	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(responses.UserResponse)

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errDb != nil || user.ID == nil {
		ctx.JSON(500, gin.H{
			"messange": "Inernal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messange": "Data user not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data transmitted",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	userEmailExist := new(models.User)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)

	if userEmailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already used",
		})
		return
	}

	user := new(models.User)
	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	ErrDB := database.DB.Table("users").Create(&user).Error
	if ErrDB != nil {
		ctx.JSON(500, gin.H{
			"message": "can't create data",
		})
		return
	}

	userResponse := responses.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}

	ctx.JSON(200, gin.H{
		"message": "Data created successfully",
		"data":    userResponse,
	})
}

func UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)
	userEmailExist := new(models.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDB := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDB != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	}

	//email exist
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmailExist != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(400, gin.H{
			"message": "Email already used",
		})
		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "can't update data",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data updated successfully",
		"data":    user,
	})
}

func DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)

	errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errFind != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	// log.Println("user", user)

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "data not found",
		})
		return
	}

	ErrDB := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.User{}).Error

	if ErrDB != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
			"error":   ErrDB.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data deleted successfully",
	})
}

func GetUserPaginate(ctx *gin.Context) {

	page := ctx.Query("page")

	if page == "" {
		page = ""
	}

	perPage := ctx.Query("perPage")

	if perPage == "" {
		perPage = "10"
	}

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]models.User)
	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}
