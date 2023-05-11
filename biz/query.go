package biz

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/linfang-canway/db-demo/dal"
	"github.com/linfang-canway/db-demo/dal/dao"
	"github.com/linfang-canway/db-demo/dal/table"
)

var q = dao.Use(dal.DB.Debug())

func Create(ctx context.Context) {

	var err error

	ud := q.User.WithContext(ctx)

	name := "modi"
	userData := &table.User{ID: 1, Name: &name}
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`id`) VALUES ('2021-09-13 20:05:51.389','2021-09-13 20:05:51.389',NULL,'modi',0,'',1)

	// 创建一个新的校验器
	validate := validator.New()

	// 使用Validator的Struct方法校验User结构体的字段
	err = validate.Struct(userData)
	if err != nil {
		// 校验失败，输出错误信息
		fmt.Println(err)
		return
	}

	err = ud.Create(userData)

	fmt.Println(err)
}
