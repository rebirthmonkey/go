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

	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/model/v1"
	policyRepoInterface "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo"
)

// policyRepo stores the secret's info.
type policyRepo struct {
	dbEngine *gorm.DB
}

var _ policyRepoInterface.PolicyRepo = (*policyRepo)(nil)

// newPolicyRepo creates and returns a user storage.
func newPolicyRepo(cfg *mysql.CompletedConfig) policyRepoInterface.PolicyRepo {
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

	return &policyRepo{dbEngine: db}
}

// close closes the repo's DB engine.
func (p *policyRepo) close() error {
	dbEngine, err := p.dbEngine.DB()
	if err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return dbEngine.Close()
}

// Create creates a new secret.
func (p *policyRepo) Create(policy *model.Policy) error {
	if err := p.dbEngine.Create(&policy).Error; err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyRepo) Delete(username, policyName string) error {
	err := p.dbEngine.Where("username = ? and name = ?", username, policyName).Delete(&model.Policy{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithCode(errcode.ErrUnknown, err.Error())
		}

		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyRepo) DeleteByUser(username string) error {
	err := p.dbEngine.Where("username = ?", username).Delete(&model.Policy{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithCode(errcode.ErrUnknown, err.Error())
		}

		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyRepo) Update(secret *model.Policy) error {
	if err := p.dbEngine.Save(secret).Error; err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyRepo) Get(username, name string) (*model.Policy, error) {
	policy := &model.Policy{}
	err := p.dbEngine.Where("username = ? and name= ?", username, name).First(&policy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrUnknown, err.Error())
		}

		return nil, errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return policy, nil
}

func (p *policyRepo) List(username string) (*model.PolicyList, error) {
	ret := &model.PolicyList{}

	//if username != "" {
	//	s.dbEngine = s.dbEngine.Where("username = ?", username)
	//}

	d := p.dbEngine.
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
