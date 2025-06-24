package mysql

import (
	"Web_template/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

// Init 初始化MySQL连接
// 参数cfg是*settings.MySQLConfig结构体指针，用于传递MySQL数据库的配置信息
func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	// 也可以使用MustConnect连接不成功就panic
	// 注意：Go的MySQL驱动有一个限制，就是数据库必须支持多语句执行，否则使用sqlx的Batch方法会报错
	// 所以，这里我们使用了sqlx.Connect来连接MySQL，而不是sqlx.MustConnect
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// 由于db小写，所以只能在本包中使用
// 使用Close()方法关闭数据库连接,给外部调用
func Close() {
	db.Close()
}
