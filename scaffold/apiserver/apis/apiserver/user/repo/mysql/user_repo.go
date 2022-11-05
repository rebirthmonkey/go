// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"fmt"
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo"
	"regexp"

	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/rebirthmonkey/go/pkg/mysql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// userRepo stores the user's info.
type userRepo struct {
	dbEngine *gorm.DB
}

var _ userRepoInterface.UserRepo = (*userRepo)(nil)

// newUserRepo creates and returns a user storage.
func newUserRepo(cfg *mysql.CompletedConfig) userRepoInterface.UserRepo {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Database,
		true,
		"Local")

	db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Mysql connection fails %+v\n", err)
		return nil
	}

	return &userRepo{dbEngine: db}
}

// close closes the repo's DB engine.
func (u *userRepo) close() error {
	dbEngine, err := u.dbEngine.DB()
	if err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return dbEngine.Close()
}

// Create creates a new user account.
func (u *userRepo) Create(user *model.User) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name != "" {
		err := errors.WithCode(errcode.ErrRecordAlreadyExist, "the created user already exit")

		log.Errorf("%+v", err)
		return err
	}

	err := u.dbEngine.Create(&user).Error
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry", err.Error()); match {
			return errors.WrapC(err, errcode.ErrRecordAlreadyExist, "duplicate entry.")
		}

		return err
	}

	return nil
}

// Delete deletes the user by the user identifier.
func (u *userRepo) Delete(username string) error {
	//tmpUser := model.User{}
	//u.dbEngine.Where("name = ?", username).Find(&tmpUser)
	//if tmpUser.Name == "" {
	//	err := errors.WithCode(errcode.ErrRecordNotFound, "the delete user not found")
	//	log.Errorf("%s\n", err)
	//	return err
	//}

	if err := u.dbEngine.Where("name = ?", username).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}

// Update updates a user account information.
func (u *userRepo) Update(user *model.User) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name == "" {
		err := errors.WithCode(errcode.ErrRecordNotFound, "the update user not found")
		log.Errorf("%s\n", err)
		return err
	}

	if err := u.dbEngine.Save(user).Error; err != nil {
		return err
	}

	return nil
}

// Get returns a user's info by the user identifier.
func (u *userRepo) Get(username string) (*model.User, error) {
	user := &model.User{}
	err := u.dbEngine.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrRecordNotFound, "the get user not found.")
		}

		return nil, errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return user, nil
}

// List returns all the related users.
func (u *userRepo) List() (*model.UserList, error) {
	ret := &model.UserList{}

	d := u.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
