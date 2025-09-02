# Go Web 项目脚手架

一个功能完整的 Go 语言项目脚手架，集成了常用的开发组件和最佳实践，帮助您快速启动 Go Web 项目开发。

## 功能特性

- 🚀 **快速启动**: 预配置的项目结构，开箱即用
- 🗄️ **MySQL 主从支持**: 内置读写分离数据库访问层
- 🔥 **Redis 集成**: 高性能缓存支持
- 📊 **结构化日志**: 使用 Uber Zap 实现高性能日志记录
- ⚙️ **配置管理**: 基于 Viper 的配置文件管理
- 🛡️ **数据库 ORM**: GORM 支持，简化数据库操作
- 📦 **模块化设计**: 清晰的代码结构，易于维护和扩展

## 技术栈

- **语言**: Go 1.19+
- **数据库**: MySQL 8.0 (支持主从复制)
- **缓存**: Redis
- **ORM**: GORM
- **日志**: Zap
- **配置**: Viper

## 项目结构

```
.
├── config
│   └── config.yaml          # 配置文件
├── internal
│   ├── cmd
│   │   └── main.go          # 程序入口
│   ├── config
│   │   └── config.go        # 配置解析模块
│   ├── dao
│   │   └── mysql
│   │       ├── conn.go      # 数据库连接管理
│   │       └── mysql.go     # MySQL 主从连接实现
│   ├── logger
│   │   └── logger.go        # 日志模块
│   └── redis
│       └── redis.go         # Redis 连接管理
├── create_table.sql         # 数据库表结构创建脚本
└── docker-compose.yml       # Docker 部署配置
```

## 快速开始

### 1. 克隆项目

```bash
git clone <your-repo-url>
cd <project-name>
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置环境

复制并修改配置文件：

```bash
cp config/config.yaml.example config/config.yaml
```

修改 `config/config.yaml` 中的数据库和 Redis 配置以匹配您的环境。

### 4. 数据库设置

确保您已经配置好 MySQL 主从复制环境，或者修改配置文件使用单一数据库实例。

执行数据库脚本创建表结构：

```bash
mysql -u username -p < create_table.sql
```

### 5. 运行项目

```bash
go run main.go
```

或者编译后运行：

```bash
go build -o webtemplate
./webtemplate
```

## 核心组件说明

### MySQL 主从复制支持

项目内置了 MySQL 主从复制支持，自动实现读写分离：

```go
// 写操作使用主库
err := dbConn.Master().Create(&user).Error

// 读操作使用从库
var user User
err := dbConn.Slave().First(&user, id).Error
```

### 配置管理

使用 Viper 管理配置，支持多种格式（YAML、JSON、TOML 等）：

```yaml
mysql:
  master:
    host: "localhost"
    port: 3308
    username: "root"
    password: "password"
    dbname: "test"
  slaves:
    - host: "localhost"
      port: 3307
      username: "root"
      password: "password"
      dbname: "test"
```

### 日志系统

使用 Uber Zap 实现高性能结构化日志：

```go
logger.Logger.Info("User created successfully", 
    zap.String("username", user.Name),
    zap.Int("user_id", user.ID))
```

### Redis 集成

预配置的 Redis 客户端，支持连接池和自动重连：

```go
err := redis.RedisClient.Set(ctx, key, value, expiration).Err()
```

## Docker 支持

项目支持 Docker 部署，可以使用 Docker Compose 启动完整的开发环境：

```bash
docker-compose up -d
```

## 扩展开发

### 添加新的业务模块

1. 在 `internal/` 目录下创建新的业务模块目录
2. 实现业务逻辑
3. 在 `main.go` 中初始化模块

### 添加新的配置项

1. 在 `config/config.yaml` 中添加配置项
2. 在 `internal/config/config.go` 中定义对应的结构体
3. 在业务代码中使用配置

### 数据库操作

使用 GORM 进行数据库操作：

```go
// 创建记录
db.Create(&user)

// 查询记录
db.First(&user, 1)
db.Find(&users)

// 更新记录
db.Model(&user).Update("name", "new name")

// 删除记录
db.Delete(&user)
```

## 最佳实践

1. **错误处理**: 统一的错误处理机制
2. **日志记录**: 完整的操作日志和错误日志
3. **配置管理**: 环境隔离的配置管理
4. **数据库访问**: 读写分离优化数据库性能
5. **缓存策略**: 合理使用 Redis 缓存提升性能
