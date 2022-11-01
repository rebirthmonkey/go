// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"fmt"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/rebirthmonkey/go/pkg/mysql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/secret/model/v1"
	secretRepoInterface "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/secret/repo"
)

// secretRepo stores the secret's info.
type secretRepo struct {
	dbEngine *gorm.DB
}

var _ secretRepoInterface.SecretRepo = (*secretRepo)(nil)

// newSecretRepo creates and returns a user storage.
func newSecretRepo(cfg *mysql.CompletedConfig) secretRepoInterface.SecretRepo {
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

	return &secretRepo{dbEngine: db}
}

// close closes the repo's DB engine.
func (s *secretRepo) close() error {
	dbEngine, err := s.dbEngine.DB()
	if err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return dbEngine.Close()
}

// Create creates a new secret.
func (s *secretRepo) Create(secret *model.Secret) error {
	if err := s.dbEngine.Create(&secret).Error; err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secretRepo) Delete(username, secretName string) error {
	err := s.dbEngine.Where("username = ? and name = ?", username, secretName).Delete(&model.Secret{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithCode(errcode.ErrUnknown, err.Error())
		}

		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secretRepo) DeleteByUser(username string) error {
	err := s.dbEngine.Where("username = ?", username).Delete(&model.Secret{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithCode(errcode.ErrUnknown, err.Error())
		}

		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secretRepo) Update(secret *model.Secret) error {
	if err := s.dbEngine.Save(secret).Error; err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secretRepo) Get(username, name string) (*model.Secret, error) {
	secret := &model.Secret{}
	err := s.dbEngine.Where("username = ? and name= ?", username, name).First(&secret).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrUnknown, err.Error())
		}

		return nil, errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return secret, nil
}

func (s *secretRepo) List(username string) (*model.SecretList, error) {
	ret := &model.SecretList{}

	//if username != "" {
	//	s.dbEngine = s.dbEngine.Where("username = ?", username)
	//}

	d := s.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	if d.Error != nil {
		//return nil, errors.WithCode(errcode.ErrDatabase, d.Error.Error())
	}

	return ret, nil
}
