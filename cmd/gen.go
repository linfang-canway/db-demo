package main

import (
	"github.com/linfang-canway/db-demo/dal/method"
	"github.com/linfang-canway/db-demo/dal/table"
	"gorm.io/gen"
)

const MySQLDSN = "root:123456@(localhost:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dal/query",
		ModelPkgPath: "./dal/table",
		Mode:         gen.WithDefaultQuery,
	})

	g.ApplyBasic(table.User{})

	g.ApplyInterface(func(method.UserMethod) {}, table.User{})

	g.Execute()

}

// 根据数据库反推出数据结构
/*func main()  {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath:      "./dal/dao",
		ModelPkgPath: "./dal/table",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithoutContext,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag

		WithUnitTest: false, // 是否生成测试文件
	})
	// 设置目标 db
	g.UseDB(db)

	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	users := g.GenerateModel("users")
	// 创建全部模型文件, 并覆盖前面创建的同名模型
	allModel := g.GenerateAllTable()

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	//基于 table 来生成基础 DAL 代码
	g.ApplyBasic(allModel...)

	//指明我们希望基于什么 model 和 interface 来生成自定义的接口实现。
	g.ApplyInterface(func(method method.UserMethod) {}, users)

	// 执行
	g.Execute()
}*/
