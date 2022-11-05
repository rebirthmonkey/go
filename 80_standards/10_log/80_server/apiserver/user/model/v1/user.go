package v1

import (
	"time"

	"github.com/rebirthmonkey/go/pkg/auth"
	"github.com/rebirthmonkey/go/pkg/metamodel"
	"github.com/rebirthmonkey/go/pkg/util"
	"gorm.io/gorm"
)

type User struct {
	metamodel.ObjectMeta `json:"metadata,omitempty"`

	Status      int64     `json:"status"              gorm:"column:status"    validate:"omitempty"`
	Nickname    string    `json:"nickname"            gorm:"column:nickname"  validate:"required,min=1,max=30"`
	Password    string    `json:"password,omitempty"  gorm:"column:password"  validate:"required"`
	Email       string    `json:"email"               gorm:"column:email"     validate:"required,email,min=1,max=100"`
	Phone       string    `json:"phone"               gorm:"column:phone"     validate:"omitempty"`
	IsAdmin     int64     `json:"isAdmin,omitempty"   gorm:"column:isAdmin"   validate:"omitempty"`
	TotalPolicy string    `json:"totalPolicy"         gorm:"-"                validate:"omitempty"`
	LoginedAt   time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

type UserList struct {
	// +optional
	metamodel.ListMeta `json:",inline"`

	Items []*User `json:"items"`
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "user"
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)

	return
}

// AfterCreate run after create database record.
func (u *User) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = util.GetInstanceID(u.ID, "user-")

	return tx.Save(u).Error
}
