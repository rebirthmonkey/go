package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (u *controller) List(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: list")

	users, err := u.srv.NewUserService().List()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, users)
}
