# API 测试指南

本文档提供了使用 Postman 或其他 HTTP 客户端测试博客系统 API 的详细说明和测试用例。

## 准备工作

1. 确保博客系统正在运行（默认端口为 8080）
2. 安装 Postman 或其他 API 测试工具（如 curl、Insomnia 等）

## 测试用例

### 1. 用户注册

**请求：**
```
POST http://localhost:8080/users/register
Content-Type: application/json
```

**请求体：**
```json
{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

**预期响应：**
```
Status: 201 Created
```
```json
{
  "id": 1,
  "username": "testuser",
  "email": "test@example.com",
  "created_at": "2025-07-30T15:04:05Z"
}
```

### 2. 用户登录

**请求：**
```
POST http://localhost:8080/users/login
Content-Type: application/json
```

**请求体：**
```json
{
  "username": "testuser",
  "password": "password123"
}
```

**预期响应：**
```
Status: 200 OK
```
```json
{
  "user_id": 1,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

保存返回的 token，后续需要认证的请求都需要在请求头中添加：
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### 3. 获取所有文章（无需认证）

**请求：**
```
GET http://localhost:8080/posts/
```

**预期响应：**
```
Status: 200 OK
```
```json
[]
```

### 4. 创建文章（需要认证）

**请求：**
```
POST http://localhost:8080/posts/
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**请求体：**
```json
{
  "title": "我的第一篇博客",
  "content": "这是我的第一篇博客文章内容。"
}
```

**预期响应：**
```
Status: 201 Created
```
```json
{
  "id": 1,
  "title": "我的第一篇博客",
  "content": "这是我的第一篇博客文章内容。",
  "user_id": 1,
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-07-30T15:04:05Z",
    "updated_at": "2025-07-30T15:04:05Z"
  },
  "created_at": "2025-07-30T15:04:05Z",
  "updated_at": "2025-07-30T15:04:05Z",
  "comments": []
}
```

### 5. 获取特定文章（无需认证）

**请求：**
```
GET http://localhost:8080/posts/1
```

**预期响应：**
```
Status: 200 OK
```
```json
{
  "id": 1,
  "title": "我的第一篇博客",
  "content": "这是我的第一篇博客文章内容。",
  "user_id": 1,
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-07-30T15:04:05Z",
    "updated_at": "2025-07-30T15:04:05Z"
  },
  "created_at": "2025-07-30T15:04:05Z",
  "updated_at": "2025-07-30T15:04:05Z",
  "comments": []
}
```

### 6. 更新文章（需要认证，且只能由文章作者更新）

**请求：**
```
PUT http://localhost:8080/posts/1
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**请求体：**
```json
{
  "title": "我的第一篇博客（已更新）",
  "content": "这是更新后的博客文章内容。"
}
```

**预期响应：**
```
Status: 200 OK
```
```json
{
  "id": 1,
  "title": "我的第一篇博客（已更新）",
  "content": "这是更新后的博客文章内容。",
  "user_id": 1,
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-07-30T15:04:05Z",
    "updated_at": "2025-07-30T15:04:05Z"
  },
  "created_at": "2025-07-30T15:04:05Z",
  "updated_at": "2025-07-30T15:05:00Z",
  "comments": []
}
```

### 7. 创建评论（需要认证）

**请求：**
```
POST http://localhost:8080/comments/
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**请求体：**
```json
{
  "content": "这是一条很好的文章！",
  "post_id": 1
}
```

**预期响应：**
```
Status: 201 Created
```
```json
{
  "id": 1,
  "content": "这是一条很好的文章！",
  "user_id": 1,
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-07-30T15:04:05Z",
    "updated_at": "2025-07-30T15:04:05Z"
  },
  "post_id": 1,
  "post": {
    "id": 1,
    "title": "我的第一篇博客（已更新）",
    "content": "这是更新后的博客文章内容。",
    "user_id": 1,
    "created_at": "2025-07-30T15:04:05Z",
    "updated_at": "2025-07-30T15:05:00Z"
  },
  "created_at": "2025-07-30T15:06:00Z"
}
```

### 8. 获取文章的所有评论（无需认证）

**请求：**
```
GET http://localhost:8080/comments/post/1
```

**预期响应：**
```
Status: 200 OK
```
```json
[
  {
    "id": 1,
    "content": "这是一条很好的文章！",
    "user_id": 1,
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-07-30T15:04:05Z",
      "updated_at": "2025-07-30T15:04:05Z"
    },
    "post_id": 1,
    "post": {
      "id": 1,
      "title": "我的第一篇博客（已更新）",
      "content": "这是更新后的博客文章内容。",
      "user_id": 1,
      "created_at": "2025-07-30T15:04:05Z",
      "updated_at": "2025-07-30T15:05:00Z"
    },
    "created_at": "2025-07-30T15:06:00Z"
  }
]
```

### 9. 删除文章（需要认证，且只能由文章作者删除）

**请求：**
```
DELETE http://localhost:8080/posts/1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**预期响应：**
```
Status: 200 OK
```
```json
{
  "message": "Post deleted successfully"
}
```

## 错误情况测试

### 1. 注册已存在的用户名

**请求：**
```
POST http://localhost:8080/users/register
Content-Type: application/json
```

**请求体：**
```json
{
  "username": "testuser",
  "email": "another@example.com",
  "password": "password123"
}
```

**预期响应：**
```
Status: 409 Conflict
```
```json
{
  "error": "Username or email already exists"
}
```

### 2. 使用错误的密码登录

**请求：**
```
POST http://localhost:8080/users/login
Content-Type: application/json
```

**请求体：**
```json
{
  "username": "testuser",
  "password": "wrongpassword"
}
```

**预期响应：**
```
Status: 401 Unauthorized
```
```json
{
  "error": "Invalid username or password"
}
```

### 3. 访问不存在的文章

**请求：**
```
GET http://localhost:8080/posts/999
```

**预期响应：**
```
Status: 404 Not Found
```
```json
{
  "error": "Post not found"
}
```

### 4. 未认证尝试创建文章

**请求：**
```
POST http://localhost:8080/posts/
Content-Type: application/json
```

**请求体：**
```json
{
  "title": "未认证的文章",
  "content": "这篇文章不应该被创建。"
}
```

**预期响应：**
```
Status: 401 Unauthorized
```
```json
{
  "error": "Authorization header required"
}
```

### 5. 使用无效的 JWT token

**请求：**
```
POST http://localhost:8080/posts/
Content-Type: application/json
Authorization: Bearer invalidtoken
```

**请求体：**
```json
{
  "title": "使用无效token的文章",
  "content": "这篇文章不应该被创建。"
}
```

**预期响应：**
```
Status: 401 Unauthorized
```
```json
{
  "error": "Invalid token"
}
```

## 使用 curl 进行测试

你也可以使用 curl 命令行工具进行测试：

### 用户注册
```bash
curl -X POST http://localhost:8080/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"curluser","email":"curl@example.com","password":"password123"}'
```

### 用户登录
```bash
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"curluser","password":"password123"}'
```

### 创建文章
```bash
curl -X POST http://localhost:8080/posts/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d '{"title":"使用curl创建的文章","content":"这是使用curl命令创建的文章。"}'
```

### 获取所有文章
```bash
curl -X GET http://localhost:8080/posts/
```

## 测试结果示例

在正常情况下，你应该看到类似以下的测试结果：

1. 所有 API 端点都能正确响应
2. 用户注册和登录功能正常工作
3. JWT 认证机制正确验证请求
4. 用户只能操作自己的文章（更新和删除）
5. 文章和评论的创建、读取功能正常
6. 错误情况能返回适当的 HTTP 状态码和错误信息

如果所有测试都通过，说明博客系统的 API 功能正常工作。