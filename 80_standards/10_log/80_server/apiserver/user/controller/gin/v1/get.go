package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (u *controller) Get(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: get")

	user, err := u.srv.NewUserService().Get(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, user)
}
