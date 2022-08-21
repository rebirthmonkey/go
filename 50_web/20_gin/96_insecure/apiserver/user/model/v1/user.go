package v1

import (
	"time"

	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/metamodel"
)

type User struct {
	metamodel.ObjectMeta `json:"metadata,omitempty"`

	Status      int       `json:"status"              gorm:"column:status"    validate:"omitempty"`
	Nickname    string    `json:"nickname"            gorm:"column:nickname"  validate:"required,min=1,max=30"`
	Password    string    `json:"password,omitempty"  gorm:"column:password"  validate:"required"`
	Email       string    `json:"email"               gorm:"column:email"     validate:"required,email,min=1,max=100"`
	Phone       string    `json:"phone"               gorm:"column:phone"     validate:"omitempty"`
	IsAdmin     int       `json:"isAdmin,omitempty"   gorm:"column:isAdmin"   validate:"omitempty"`
	TotalPolicy int64     `json:"totalPolicy"         gorm:"-"                validate:"omitempty"`
	LoginedAt   time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

type UserList struct {
	// Standard list metadata.
	// +optional
	metamodel.ListMeta `json:",inline"`

	Items []*User `json:"items"`
}
