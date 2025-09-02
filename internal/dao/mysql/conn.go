// internal/dao/mysql/conn.go
package mysql

import (
	"database/sql"
	"gorm.io/gorm"
)

type Conn struct {
	manager *MySQLManager
}

func NewConn(manager *MySQLManager) *Conn {
	return &Conn{manager: manager}
}

// Master 获取主库连接（用于写操作）
func (c *Conn) Master() *gorm.DB {
	return c.manager.Master()
}

// Slave 获取从库连接（用于读操作）
func (c *Conn) Slave() *gorm.DB {
	return c.manager.Slave()
}

// GetSQLDB 获取底层的 sql.DB 连接
func (c *Conn) GetSQLDB(isMaster bool) *sql.DB {
	if isMaster {
		sqlDB, _ := c.manager.Master().DB()
		return sqlDB
	}

	sqlDB, _ := c.manager.Slave().DB()
	return sqlDB
}

// Close 关闭所有数据库连接
func (c *Conn) Close() error {
	return c.manager.Close()
}
