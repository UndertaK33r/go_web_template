// internal/dao/mysql/mysql.go
package mysql

import (
	"file/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLManager struct {
	master *gorm.DB
	slave  *gorm.DB
}

// NewMySQLManager 创建主从数据库管理器
func NewMySQLManager(cfg *config.MySQLConfig) (*MySQLManager, error) {
	manager := &MySQLManager{}

	// 连接主库
	masterDB, err := connectToDB(&cfg.Master)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to master database: %w", err)
	}
	manager.master = masterDB

	// 连接从库
	slaveDB, err := connectToDB(&cfg.Slaves)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to slave database: %w", err)
	}
	manager.slave = slaveDB

	return manager, nil
}

// connectToDB 连接到单个数据库实例
func connectToDB(cfg *config.MySQLInstanceConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
		cfg.ParseTime,
		cfg.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 可以根据需要配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

// Master 获取主库连接（用于写操作）
func (m *MySQLManager) Master() *gorm.DB {
	return m.master
}

// Slave 获取从库连接（用于读操作）
func (m *MySQLManager) Slave() *gorm.DB {
	return m.slave
}

// Close 关闭所有数据库连接
func (m *MySQLManager) Close() error {
	// 关闭主库连接
	if m.master != nil {
		sqlDB, err := m.master.DB()
		if err != nil {
			return err
		}
		sqlDB.Close()
	}

	// 关闭从库连接
	if m.slave != nil && m.slave != m.master {
		sqlDB, err := m.slave.DB()
		if err != nil {
			return err
		}
		sqlDB.Close()
	}

	return nil
}
