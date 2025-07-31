# 博客系统

一个使用 Go、Gin 框架和 GORM 库构建的个人博客系统。该项目实现了基本的博客功能，包括用户认证、文章管理和评论系统。

## 功能特性

- 使用 JWT 认证的用户注册和登录
- 文章的创建、读取、更新和删除操作
- 文章评论系统
- 使用 bcrypt 进行密码加密
- 使用 SQLite 数据库便于设置和测试

## 前置要求

- Go 1.23+
- Git

## 依赖项

- Gin Web 框架
- GORM (ORM 库)
- SQLite 驱动
- 用于认证的 JWT
- 用于密码哈希的 bcrypt

## 安装步骤

1. 克隆仓库：
   ```bash
   git clone <repository-url>
   ```

2. 导航到项目目录：
   ```bash
   cd task4/blog
   ```

3. 初始化 Go 模块（如果尚未完成）：
   ```bash
   go mod init blog
   ```

4. 安装所需依赖：
   ```bash
   go get github.com/gin-gonic/gin
   go get gorm.io/gorm
   go get github.com/glebarez/sqlite
   go get github.com/golang-jwt/jwt/v5
   go get golang.org/x/crypto/bcrypt
   ```

5. 整理模块依赖：
   ```bash
   go mod tidy
   ```

## 项目结构

```
blog/
├── main.go                 # 应用程序入口点
├── go.mod                  # Go 模块依赖
├── go.sum                  # 依赖校验和
├── README.md               # 本文件
├── API_TESTING.md          # API 测试指南
├── controllers/            # 请求处理器
│   ├── user_controller.go  # 用户注册和登录
│   ├── post_controller.go  # 文章管理
│   └── comment_controller.go # 评论管理
├── database/               # 数据库初始化
│   └── database.go         # 数据库连接和迁移
├── middleware/             # 中间件函数
│   └── jwt.go              # JWT 认证中间件
├── models/                 # 数据库模型
│   └── models.go           # 用户、文章和评论模型
├── routes/                 # 路由定义
│   └── routes.go           # API 路由设置
└── utils/                  # 工具函数
    └── password.go         # 密码哈希函数
```

## 运行应用程序

1. 确保你在项目目录中：
   ```bash
   cd task4/blog
   ```

2. 运行应用程序（解决CGO问题）：
   ```bash
   # 方法1: 启用CGO（推荐）
   CGO_ENABLED=1 go run main.go
   
   # Windows PowerShell中的方法1:
   $env:CGO_ENABLED=1; go run main.go
   
   # Windows CMD中的方法1:
   set CGO_ENABLED=1 && go run main.go
   
   # 方法2: 如果无法启用CGO，则使用纯Go的SQLite驱动（需要修改代码）
   go run main.go
   ```

3. 服务器将在 `http://localhost:8080` 启动



## 认证

大多数 API 端点都需要使用 JWT 令牌进行认证。要进行认证：

1. 注册新用户或使用现有凭据登录
2. 在后续请求的 `Authorization` 头中使用返回的 JWT 令牌：
   ```
   Authorization: Bearer <your-jwt-token>
   ```

## 数据库

该项目使用 SQLite 进行数据持久化。在首次运行时，将在项目目录中自动创建数据库文件 `blog.db`。

## 生产环境构建

要为生产环境构建应用程序：

```bash
# 启用CGO构建
CGO_ENABLED=1 go build -o blog-server main.go

# Windows PowerShell:
$env:CGO_ENABLED=1; go build -o blog-server main.go

# Windows CMD:
set CGO_ENABLED=1 && go build -o blog-server main.go
```

这将创建一个名为 `blog-server` 的可执行文件，可以直接运行。