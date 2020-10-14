package apis

import (
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

// User manages
type User struct {
}

// GetUserByID godoc
// @Summary Get a user
// @Description Get a user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Router /user [get]
func (u *User) GetUser(ctx *gin.Context) {
	var user model.User
	user.Name = "test"
	user.Email = "test@test.com"
	ctx.JSON(http.StatusOK, user)
}
