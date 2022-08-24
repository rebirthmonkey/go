package mysql

import (
	"github.com/pkg/errors"
	userRepoInterface "github.com/rebirthmonkey/iam/internal/auth/user/repo"
	"github.com/rebirthmonkey/iam/pkg/mysql"
	"gorm.io/gorm"
)

type repo struct {
	dbEngine *gorm.DB
}

var _ userRepoInterface.Repo = (*repo)(nil)

func NewRepo(opts *mysql.Options) (userRepoInterface.Repo, error) {
	mysqlConfig, err := mysql.NewConfigFromOptions(opts)
	if err != nil {
		return nil, errors.Wrap(err, "create repo failed")
	}

	dbEngine, err2 := mysql.NewDB(mysqlConfig)
	if err2 != nil {
		return nil, errors.Wrap(err2, "create repo failed")
	}

	return &repo{dbEngine}, nil
}

func (r *repo) GetUserRepo() userRepoInterface.UserRepo {
	return newUserRepo(r.dbEngine)
}

func (r *repo) Close() error {
	dbEngine, err := r.dbEngine.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm instance failed")
	}

	return dbEngine.Close()
}
