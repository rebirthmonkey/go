package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (u *controller) List(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: list")

	users, err := u.srv.NewUserService().List()
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, users)
}
