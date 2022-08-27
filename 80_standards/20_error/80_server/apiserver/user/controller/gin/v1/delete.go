package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (u *controller) Delete(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: delete")

	if err := u.srv.NewUserService().Delete(c.Param("name")); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	var msg string = "deleted user " + c.Param("name")
	util.WriteResponse(c, nil, msg)
}
