package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (u *controller) Get(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: get")

	user, err := u.srv.NewUserService().Get(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, user)
}
