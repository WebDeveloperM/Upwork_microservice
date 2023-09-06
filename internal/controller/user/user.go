package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	user2 "upwork/internal/service/user"
	"upwork/internal/usecase/user"
)

type Controller struct {
	user user.UseCase
}

func ControllerUser(user user.UseCase) Controller {
	return Controller{user: user}
}

func (ct Controller) Register(c *gin.Context) {
	var createUser user2.Create
	ctx := context.Background()

	errBind := c.ShouldBind(&createUser)
	if errBind != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errBind.Error(),
		})
		return
	}

	token, err := ct.user.Create(ctx, createUser)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": false,
		"token":   token,
	})

}
