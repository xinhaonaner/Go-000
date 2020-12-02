package main

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/jinzhu/gorm"
	xErrors "github.com/pkg/errors"
	"time"
)

var DBEngine *gorm.DB

type User struct {
	ID   int64
	Name string
}

func (User) TableName() string {
	return "users"
}

func (u *User) NoDataError() error {
	return sql.ErrNoRows
}

func (u *User) GetUser(id int64) (*User, error) {
	var doc *User

	table := DBEngine.Table(u.TableName())
	if u.ID != 0 {
		table = table.Where("id = ?", id)
	}

	// 应该将 not found 翻译为明确指令给到上层，因为对于 DAO 层，不同 DB 的 lib 库可能有区别
	if err := table.First(doc).Error; err != nil {
		if xErrors.Is(err, sql.ErrNoRows) {
			return doc, xErrors.WithMessage(err, "no data found")
		}
		return nil, xErrors.Wrap(err, "db error")
	}
	return doc, nil
}

func main() {
	Go(func() {
		glog.Info("InfoLog")
		panic("I'm panic.")
	})

	time.Sleep(time.Second)

	// 作业
	Biz()
}

func Biz() {
	var u User
	user, err := u.GetUser(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("user info: %+v\n", user)
}

// Go 避免野生协程 panic
func Go(x func()) {
	go func() {
		// 延迟处理
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover:", err)
			}
		}()
		x()
	}()
}
