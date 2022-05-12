package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model   // predefine DB metadata, including ID, CreateAt, UpdateAT, DeleteAt, 其时间是自动更新
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置num为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func (u *User) TableName() string {
	return "yyy"
}

func main() {
	var (
		host     = pflag.StringP("host", "H", "127.0.0.1:3306", "MySQL service host address")
		username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
		password = pflag.StringP("password", "p", "root", "Password for access to mysql, should be used pair with password")
		database = pflag.StringP("database", "d", "test", "Database name to use")
		help     = pflag.BoolP("help", "h", false, "Print this help message")
	)

	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 1. Auto migration (binding) for given models, and create corresponding table/schema
	db.AutoMigrate(&User{})
}
