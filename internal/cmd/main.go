package main

import (
	"file/internal/config"
	"file/internal/dao/mysql"
	"file/internal/logger"
	"file/internal/redis"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("load config error: %v", err)
	}

	err = logger.InitLogger(&cfg.Log)
	if err != nil {
		log.Fatal("init logger error: %v", err)
	}
	// 同步日志
	defer logger.Sync()

	err = redis.InitRedis(&cfg.Redis)
	if err != nil {
		log.Fatal("init redis error: %v", err)
	}
	defer redis.Close()
	log.Println("redis init success")

	//初始化mysql
	mannager, err := mysql.NewMySQLManager(&cfg.MySQL)
	if err != nil {
		log.Fatal("init mysql error: %v", err)
	}
	defer mannager.Close()
	log.Println("mysql init success")
	conn := mysql.NewConn(mannager)
	log.Println("mysql conn init success")
	defer conn.Close()

}
