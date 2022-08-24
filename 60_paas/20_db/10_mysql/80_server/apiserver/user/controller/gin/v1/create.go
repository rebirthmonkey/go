package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/model/v1"
)

func (u *controller) Create(c *gin.Context) {
	fmt.Println("[GINServer] userController: create")

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
