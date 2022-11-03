// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/rebirthmonkey/go/pkg/metamodel"
	"github.com/rebirthmonkey/go/pkg/util"
	"gorm.io/gorm"
)

// Secret represents a secret restful resource.
// It is also used as gorm model.
type Secret struct {
	// Standard object's metadata.
	metamodel.ObjectMeta `       json:"metadata,omitempty"`
	Username             string `json:"username"           gorm:"column:username"  validate:"omitempty"`
	SecretID             string `json:"secretID"           gorm:"column:secretID"  validate:"omitempty"`
	SecretKey            string `json:"secretKey"          gorm:"column:secretKey" validate:"omitempty"`

	// Required: true
	Expires     int64  `json:"expires"     gorm:"column:expires"     validate:"omitempty"`
	Description string `json:"description" gorm:"column:description" validate:"description"`
}

type SecretList struct {
	metamodel.ListMeta `          json:",inline"`
	Items              []*Secret `json:"items"`
}

// TableName maps to mysql table name.
func (s *Secret) TableName() string {
	return "secret"
}

// AfterCreate run after create database record.
func (s *Secret) AfterCreate(tx *gorm.DB) error {
	s.InstanceID = util.GetInstanceID(s.ID, "secret-")

	return tx.Save(s).Error
}
