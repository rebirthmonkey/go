package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
)

func main() {
	r := gin.Default()

	var policy = &ladon.DefaultPolicy{
		ID:          "888",
		Description: "Hair Design",
		Subjects:    []string{"<Tony|Kevin|Allen>"},
		Resources:   []string{"resources:hair"},
		Actions:     []string{"delete", "<create|update>"},
		Effect:      ladon.AllowAccess,
	}

	r.POST("/authz", func(c *gin.Context) {
		accessRequest := &ladon.Request{}
		var message string
		if err := c.BindJSON(accessRequest); err != nil {
			fmt.Println(err)
		} else {
			warden := &ladon.Ladon{
				Manager:     memory.NewMemoryManager(),
				AuditLogger: &ladon.AuditLoggerInfo{},
			}
			warden.Manager.Create(policy)
			if err := warden.IsAllowed(accessRequest); err != nil {
				message = "无操作权限"
			} else {
				message = "有操作权限"
			}

			c.JSON(200, gin.H{"message": message})
		}
	})

	r.Run(":8080")
}
