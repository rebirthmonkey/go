package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/log"

	model "github.com/rebirthmonkey/go/80_standards/10_log/80_server/apiserver/user/model/v1"
)

func (u *controller) Create(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: create")

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := u.srv.NewUserService().Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}