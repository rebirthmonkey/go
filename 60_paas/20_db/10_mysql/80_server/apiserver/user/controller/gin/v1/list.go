package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *controller) List(c *gin.Context) {
	fmt.Println("[GinServer] userController: list")

	users, err := u.srv.NewUserService().List()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, users)
}
