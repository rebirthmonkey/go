package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *controller) Get(c *gin.Context) {
	fmt.Println("[GINServer] userController: get")

	user, err := u.srv.NewUserService().Get(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, user)
}
