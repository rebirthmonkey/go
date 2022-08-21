package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *controller) Delete(c *gin.Context) {
	fmt.Println("[GINServer] userController: delete")

	if err := u.srv.NewUserService().Delete(c.Param("name")); err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, "delete user")
}
