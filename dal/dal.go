package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
	"sync"
)

var DB *gorm.DB
var once sync.Once

const MySQLDSN = "root:123456@(localhost:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	once.Do(func() {
		DB = ConnectDB().Debug()
	})
}

func ConnectDB() (conn *gorm.DB) {
	conn, err := gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{})

	//GORM 提供了 Prometheus 插件来收集 DBStats 和用户自定义指标
	conn.Use(prometheus.New(prometheus.Config{
		DBName:          "db1",                       // 使用 `DBName` 作为指标 label
		RefreshInterval: 15,                          // 指标刷新频率（默认为 15 秒）
		PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
		StartServer:     true,                        // 启用一个 http 服务来暴露指标
		HTTPServerPort:  8080,                        // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // 用户自定义指标
	}))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
