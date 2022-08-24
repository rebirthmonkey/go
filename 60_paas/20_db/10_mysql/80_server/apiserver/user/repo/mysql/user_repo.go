package mysql

import (
	"github.com/rebirthmonkey/pkg/errors"
	"github.com/rebirthmonkey/pkg/log"
	"gorm.io/gorm"
	"regexp"

	policyRepo "github.com/rebirthmonkey/iam/internal/auth/policy/repo"
	secretRepo "github.com/rebirthmonkey/iam/internal/auth/secret/repo"
	model "github.com/rebirthmonkey/iam/internal/auth/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/iam/internal/auth/user/repo"
	"github.com/rebirthmonkey/iam/internal/pkg/errcode"
)

type userRepo struct {
	dbEngine *gorm.DB
}

var _ userRepoInterface.UserRepo = (*userRepo)(nil)

func newUserRepo(dbEngine *gorm.DB) userRepoInterface.UserRepo {
	return &userRepo{dbEngine}
}

func (u *userRepo) Create(user *model.User) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name != "" {
		log.Warn("the created user already exists.")
		return nil
	}

	err := u.dbEngine.Create(&user).Error
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry", err.Error()); match {
			return errors.WithCode(errcode.ErrUserAlreadyExist, err.Error())
		}

		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (u *userRepo) Delete(username string) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", username).Find(&tmpUser)
	if tmpUser.Name == "" {
		log.Warn("the deleted user does not exist.")
		return nil
	}

	_ = policyRepo.GetRepo().GetPolicyRepo().DeleteByUser(username)

	_ = secretRepo.GetRepo().GetSecretRepo().DeleteByUser(username)

	if err := u.dbEngine.Where("name = ?", username).Delete(&model.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithCode(errcode.ErrUserNotFound, err.Error())
		}

		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (u *userRepo) Update(user *model.User) error {
	if err := u.dbEngine.Save(user).Error; err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return nil
}

func (u *userRepo) Get(username string) (*model.User, error) {
	user := &model.User{}
	err := u.dbEngine.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrUserNotFound, err.Error())
		}
		return nil, errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return user, nil
}

func (u *userRepo) List() (*model.UserList, error) {
	ret := &model.UserList{}

	d := u.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	if d.Error != nil {
		return nil, errors.WithCode(errcode.ErrDatabase, d.Error.Error())
	}

	return ret, nil
}
