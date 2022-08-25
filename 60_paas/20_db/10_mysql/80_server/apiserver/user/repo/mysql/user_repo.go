package mysql

import (
	"fmt"
	"regexp"

	"github.com/rebirthmonkey/go/pkg/mysql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/repo"
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
		fmt.Printf("%+v", err)
		return nil
	}

	return &userRepo{dbEngine: db}
}

func (u *userRepo) close() error {
	dbEngine, err := u.dbEngine.DB()
	if err != nil {
		return err
	}

	return dbEngine.Close()
}

func (u *userRepo) Create(user *model.User) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", user.Name).Find(&tmpUser)
	if tmpUser.Name != "" {
		fmt.Println("the created user already exists.")
		return nil
	}

	err := u.dbEngine.Create(&user).Error
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry", err.Error()); match {
			return err
		}

		return err
	}

	return nil
}

func (u *userRepo) Delete(username string) error {
	tmpUser := model.User{}
	u.dbEngine.Where("name = ?", username).Find(&tmpUser)
	if tmpUser.Name == "" {
		fmt.Println("the deleted user does not exist.")
		return nil
	}

	if err := u.dbEngine.Where("name = ?", username).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Update(user *model.User) error {
	if err := u.dbEngine.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Get(username string) (*model.User, error) {
	user := &model.User{}
	err := u.dbEngine.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
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
		return nil, d.Error
	}

	return ret, nil
}
