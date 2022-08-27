package mysql

import (
	"fmt"
	"regexp"

	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/rebirthmonkey/go/pkg/mysql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "github.com/rebirthmonkey/go/80_standards/20_error/80_server/apiserver/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/go/80_standards/20_error/80_server/apiserver/user/repo"
)

type userRepo struct {
	dbEngine *gorm.DB
}

var _ userRepoInterface.UserRepo = (*userRepo)(nil)

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

func (u *userRepo) close() error {
	dbEngine, err := u.dbEngine.DB()
	if err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return dbEngine.Close()
}

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

func (u *userRepo) Delete(username string) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", username).Find(&tmpUser)
	if tmpUser.Name == "" {
		err := errors.WithCode(errcode.ErrRecordNotFound, "the delete user not found")
		log.Errorf("%s\n", err)
		return err
	}

	if err := u.dbEngine.Where("name = ?", username).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}

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
