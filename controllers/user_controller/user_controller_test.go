package user_controller

import (

	// "log"

	"gin-goinc-api/configs/db_config"
	"gin-goinc-api/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	db_config.DB_DRIVER = "test"
	database.ConnectDatabase()

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/user", GetAllUser)

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/user", nil)
	r.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"data":[]}` // Seu corpo de resposta esperado
	assert.Equal(t, expectedBody, w.Body.String())

}
