package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (u *controller) Delete(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: delete")

	if err := u.srv.NewUserService().Delete(c.Param("name")); err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, "delete user")
}
