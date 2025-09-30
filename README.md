# Melody Cure Backend

音美疗愈 - 儿童康复治疗平台后端服务

## 项目简介

Melody Cure 是一个专为儿童康复治疗设计的综合性平台，提供AI陪伴、虚拟疗愈导师、儿童档案管理、疗愈日志记录等功能，帮助儿童在康复过程中获得更好的治疗体验和效果。

## 技术栈

- **语言**: Go 1.24.4
- **框架**: Gin Web Framework
- **数据库**: MySQL (使用 GORM ORM)
- **缓存**: Redis
- **认证**: JWT
- **依赖注入**: Google Wire
- **配置管理**: Viper
- **API文档**: Swagger
- **图床服务**: 七牛云

## 项目结构

```
├── DAO/                    # 数据访问层
├── api/                    # API相关
│   ├── request/           # 请求结构体
│   └── response/          # 响应结构体
├── config/                # 配置文件
├── controller/            # 控制器层
├── docs/                  # Swagger文档
├── middleware/            # 中间件
├── model/                 # 数据模型
├── routes/                # 路由配置
├── service/               # 业务逻辑层
└── tool/                  # 工具类
```

## 功能特性

### 🔐 用户管理系统
- 用户注册/登录
- JWT认证机制
- 个人信息管理
- 密码修改
- 专业认证申请（机构认证/康复师认证）

### 🤖 AI陪伴功能
- 创建个性化AI陪伴角色
- 多种陪伴类型选择
- 个性化设定和语音类型

### 👨‍⚕️ 虚拟疗愈导师
- 创建专业虚拟疗愈导师
- 专业领域分类
- 经验等级设定

### 👶 儿童档案管理
- 完整的儿童档案创建
- 病情诊断记录
- 治疗方案管理
- 康复进度跟踪

### 📝 疗愈日志系统
- 记录儿童成长进步
- 疗愈前后对比（文字、照片等）
- 时间线浏览功能
- 媒体文件管理（图片、视频）

### ⭐ 收藏管理
- 课程收藏
- 游戏收藏
- 文章收藏

### 📚 内容管理
- 课程列表和详情
- 游戏列表和详情

### 🖼️ 图床服务
- 七牛云图片上传
- 安全的上传Token获取

## 快速开始

### 前置要求
- Docker (20.10+)
- Docker Compose (2.0+)
- Git

### 一键部署

1. **克隆项目**
```bash
git clone https://github.com/your-username/Melody-Cure-Backend.git
cd Melody-Cure-Backend
```

2. **配置环境**
```bash
# 复制配置文件
cp config/example.yaml config/config.yaml

# 编辑配置文件（必须修改数据库密码、JWT密钥等）
nano config/config.yaml
```

3. **启动服务**
```bash
docker-compose up -d
```

4. **验证部署**
```bash
# 检查服务状态
docker-compose ps

# 健康检查
curl http://localhost/health

# 访问 API 文档
# 浏览器打开: http://localhost/swagger/index.html
```

详细部署指南请参考 [DEPLOY.md](DEPLOY.md)

## 开发环境

### 环境要求

- Go 1.24.4+
- MySQL 8.0+
- Redis 6.0+

### 安装依赖

```bash
go mod download
```

### 配置文件

复制 `config/example.yaml` 为 `config/config.yaml` 并修改相应配置：

```yaml
# 数据库配置
database:
  host: localhost
  port: 3306
  username: your_username
  password: your_password
  database: melody_cure

# Redis配置
redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

# JWT配置
jwt:
  secret: your_jwt_secret
  expire: 24h
```

### 运行项目

```bash
go run main.go
```

服务将在 `http://localhost:8080` 启动

## API 接口文档

### 认证说明

除了注册和登录接口外，其他接口都需要在请求头中携带 JWT Token：

```
Authorization: Bearer <your_jwt_token>
```

### 公开接口（无需认证）

#### 用户注册

- **POST** `/api/user/register`
- **描述**: 用户注册
- **请求体**:

```json
{
  "name": "用户名",
  "password": "密码",
  "email": "邮箱",
  "phone": "手机号",
  "identity": "身份类型"
}
```

#### 用户登录

- **POST** `/api/user/login`
- **描述**: 用户登录
- **请求体**:

```json
{
  "name": "用户名",
  "password": "密码"
}
```

### 用户管理接口（需要认证）

#### 获取个人信息

- **GET** `/api/user/profile`
- **描述**: 获取当前用户的个人信息

#### 更新个人信息

- **PUT** `/api/user/profile`
- **描述**: 更新用户个人信息
- **请求体**:

```json
{
  "name": "姓名",
  "phone": "手机号",
  "address": "地址",
  "image": "头像URL"
}
```

#### 用户登出

- **POST** `/api/user/logout`
- **描述**: 用户登出，使当前token失效

#### 修改密码

- **PUT** `/api/user/password`
- **描述**: 修改用户密码
- **请求体**:

```json
{
  "old_password": "旧密码",
  "new_password": "新密码"
}
```

### 认证管理接口

#### 申请认证

- **POST** `/api/user/certification/apply`
- **描述**: 申请专业认证（机构认证/康复师认证）
- **请求体**:

```json
{
  "certificate_type": "认证类型",
  "certificate_name": "证书名称",
  "certificate_no": "证书编号",
  "issuing_authority": "颁发机构",
  "issue_date": "颁发日期",
  "expiry_date": "过期日期"
}
```

#### 获取认证状态

- **GET** `/api/user/certification/status`
- **描述**: 获取用户的认证状态

### AI陪伴功能

#### 创建AI陪伴

- **POST** `/api/user/ai-companion`
- **描述**: 创建AI陪伴角色
- **请求体**:

```json
{
  "companion_type": "陪伴类型",
  "name": "陪伴名称",
  "avatar": "头像URL",
  "personality": "性格描述",
  "voice_type": "语音类型"
}
```

#### 获取AI陪伴列表

- **GET** `/api/user/ai-companions`
- **描述**: 获取用户的所有AI陪伴列表

### 虚拟疗愈导师

#### 创建虚拟疗愈导师

- **POST** `/api/user/virtual-therapist`
- **描述**: 创建虚拟疗愈导师
- **请求体**:

```json
{
  "therapist_type": "导师类型",
  "name": "导师名称",
  "avatar": "头像URL",
  "specialization": "专业领域",
  "experience": 5
}
```

#### 获取虚拟疗愈导师列表

- **GET** `/api/user/virtual-therapists`
- **描述**: 获取用户的虚拟疗愈导师列表

### 儿童档案管理

#### 创建儿童档案

- **POST** `/api/user/child-archive`
- **描述**: 创建儿童档案
- **请求体**:

```json
{
  "child_name": "儿童姓名",
  "gender": "性别",
  "birth_date": "出生日期",
  "avatar": "头像URL",
  "condition": "病情描述",
  "diagnosis": "诊断结果",
  "treatment": "治疗方案",
  "progress": "康复进度",
  "notes": "备注"
}
```

#### 获取儿童档案列表

- **GET** `/api/user/child-archives`
- **描述**: 获取用户的儿童档案列表

#### 更新儿童档案

- **PUT** `/api/user/child-archive/:id`
- **描述**: 更新指定的儿童档案
- **参数**: `id` - 档案ID
- **请求体**: 同创建儿童档案

#### 删除儿童档案

- **DELETE** `/api/user/child-archive/:id`
- **描述**: 删除指定的儿童档案
- **参数**: `id` - 档案ID

### 收藏管理

#### 添加收藏

- **POST** `/api/user/favorite`
- **描述**: 添加内容到收藏夹
- **请求体**:

```json
{
  "resource_type": "资源类型（course/game/article）",
  "resource_id": "资源ID"
}
```

#### 获取收藏列表

- **GET** `/api/user/favorites`
- **描述**: 获取用户的收藏列表

#### 移除收藏

- **DELETE** `/api/user/favorite`
- **描述**: 从收藏夹移除内容
- **请求体**:

```json
{
  "resource_type": "资源类型",
  "resource_id": "资源ID"
}
```

### 内容管理

#### 获取课程列表

- **GET** `/api/user/courses`
- **描述**: 获取所有课程列表

#### 获取课程详情

- **GET** `/api/user/course/:id`
- **描述**: 获取指定课程的详细信息
- **参数**: `id` - 课程ID

### 游戏管理

#### 获取游戏列表

- **GET** `/api/user/games`
- **描述**: 获取所有游戏列表

#### 获取游戏详情

- **GET** `/api/user/game/:id`
- **描述**: 获取指定游戏的详细信息
- **参数**: `id` - 游戏ID

### 图床服务

#### 获取七牛云上传Token

- **GET** `/api/image/qiniu/token`
- **描述**: 获取七牛云上传Token，用于客户端直传文件
- **需要认证**: 是
- **响应**:

```json
{
  "code": 200,
  "data": {
    "token": "上传Token",
    "domain": "CDN域名",
    "bucket": "存储空间名",
    "expires_at": "过期时间",
    "use_https": true
  }
}
```

### 疗愈日志管理

#### 创建疗愈日志

- **POST** `/api/healing-log`
- **描述**: 创建一条新的疗愈日志，记录儿童成长进步和疗愈前后对比
- **需要认证**: 是
- **请求体**:

```json
{
  "child_archive_id": 1,
  "content": "今天孩子的表现很好，能够主动与其他小朋友交流...",
  "media": [
    {
      "media_type": "image",
      "url": "https://example.com/image1.jpg"
    },
    {
      "media_type": "video", 
      "url": "https://example.com/video1.mp4"
    }
  ]
}
```

#### 获取儿童疗愈日志列表

- **GET** `/api/healing-log/child/:child_id`
- **描述**: 获取指定儿童的所有疗愈日志，按时间线排序显示成长进步
- **需要认证**: 是
- **参数**: `child_id` - 儿童档案ID

#### 获取疗愈日志详情

- **GET** `/api/healing-log/:log_id`
- **描述**: 获取单个疗愈日志的详细信息，包括文字内容和媒体文件
- **需要认证**: 是
- **参数**: `log_id` - 日志ID

#### 删除疗愈日志

- **DELETE** `/api/healing-log/:log_id`
- **描述**: 删除指定的疗愈日志及其关联的媒体文件
- **需要认证**: 是
- **参数**: `log_id` - 日志ID

## 响应格式

### 成功响应

```json
{
  "code": 200,
  "message": "操作成功",
  "data": {}
}
```

### 错误响应

```json
{
  "code": 400,
  "message": "错误信息"
}
```

## 状态码说明

- `200` - 请求成功
- `400` - 请求参数错误
- `401` - 未授权或token无效
- `500` - 服务器内部错误

## Swagger API 文档

项目集成了 Swagger 自动生成 API 文档功能。

### 访问文档

启动服务后，可以通过以下地址访问 API 文档：

- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **JSON格式**: `http://localhost:8080/swagger/doc.json`

### 重新生成文档

当添加或修改 API 接口后，需要重新生成 Swagger 文档：

```bash
swag init
```

### 注释规范

在控制器方法上添加 Swagger 注释：

```go
// CreateHealingLog 创建疗愈日志
// @Summary 创建疗愈日志
// @Description 创建一条新的疗愈日志，记录儿童成长进步和疗愈前后对比
// @Tags 疗愈日志
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param healing_log body model.HealingLog true "疗愈日志信息"
// @Success 200 {object} response.SuccessResponse "创建成功"
// @Failure 400 {object} response.ErrorResponse "参数错误"
// @Failure 401 {object} response.ErrorResponse "未认证"
// @Failure 500 {object} response.ErrorResponse "创建失败"
// @Router /api/healing-log [post]
func (c *HealingLogController) CreateHealingLog(ctx *gin.Context) {
    // 实现代码...
}
```

## 开发指南

### 添加新接口

1. 在 `api/request/` 中定义请求结构体
2. 在 `api/response/` 中定义响应结构体
3. 在 `model/` 中定义数据模型
4. 在 `DAO/` 中实现数据访问层
5. 在 `service/` 中实现业务逻辑
6. 在 `controller/` 中实现控制器方法
7. 在 `routes/` 中添加路由配置
8. 在 `wire.go` 中添加依赖注入
9. 添加 Swagger 注释
10. 运行 `wire` 和 `swag init` 重新生成代码和文档

### 数据库迁移

使用 GORM 的自动迁移功能：

```go
db.AutoMigrate(
    &User{}, 
    &Certification{}, 
    &AICompanion{}, 
    &VirtualTherapist{}, 
    &ChildArchive{}, 
    &UserFavorite{}, 
    &Course{}, 
    &Game{}, 
    &HealingLog{}, 
    &LogMedia{},
    &model.ImageToken{},
)
```

## 部署

### 环境要求
- Go 1.21+
- MySQL 8.0+
- Redis 6.0+
- Docker & Docker Compose (推荐)

### Docker 部署 (推荐)

#### 使用服务器现有的 MySQL 和 Redis

如果你的服务器已经安装了 MySQL 和 Redis，可以使用以下方式部署：

**方式一：使用 host 网络模式 (推荐)**
```bash
# 使用简化版配置，直接访问宿主机服务
docker-compose -f docker-compose.simple.yml up -d

# 查看服务状态
docker-compose -f docker-compose.simple.yml ps

# 查看日志
docker-compose -f docker-compose.simple.yml logs -f backend
```

**方式二：使用桥接网络模式**
```bash
# 修改 docker-compose.yml 中的数据库和Redis连接配置
# 将 DB_HOST 和 REDIS_HOST 改为服务器的实际IP地址
docker-compose up -d
```

#### 开发环境部署 (包含所有服务)
如果需要完整的开发环境，包括 MySQL 和 Redis 容器：

1. 克隆项目
```bash
git clone <repository-url>
cd Melody-Cure-Backend
```

2. 使用完整版配置启动
```bash
# 恢复 MySQL 和 Redis 容器配置后使用
# docker-compose up -d
```

3. 访问服务
- API 服务: http://localhost:8080
- Swagger 文档: http://localhost:8080/swagger/index.html

#### 生产环境部署
1. 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，配置生产环境的密码和密钥
```

2. 启动生产环境
```bash
# 使用生产环境配置启动
docker-compose -f docker-compose.prod.yml up -d

# 查看服务状态
docker-compose -f docker-compose.prod.yml ps
```

3. SSL 配置 (可选)
```bash
# 将 SSL 证书放入 ssl 目录
mkdir ssl
cp your-cert.pem ssl/cert.pem
cp your-key.pem ssl/key.pem

# 编辑 nginx.conf 启用 HTTPS 配置
# 重启 nginx 服务
docker-compose -f docker-compose.prod.yml restart nginx
```

### 传统部署

#### 部署步骤
1. 克隆项目
```bash
git clone <repository-url>
cd Melody-Cure-Backend
```

2. 配置环境变量
```bash
cp config/example.yaml config/config.yaml
# 编辑 config/config.yaml 文件，配置数据库、Redis、JWT等信息
```

3. 安装依赖
```bash
go mod download
```

4. 运行数据库迁移
```bash
go run main.go
```

5. 启动服务
```bash
go run main.go
```

### Docker 命令参考

#### 使用服务器现有 MySQL/Redis 的命令
```bash
# 使用简化版配置 (推荐)
docker-compose -f docker-compose.simple.yml up -d
docker-compose -f docker-compose.simple.yml down
docker-compose -f docker-compose.simple.yml logs -f backend
docker-compose -f docker-compose.simple.yml restart backend

# 使用标准配置 (需要修改连接地址)
docker-compose up -d
docker-compose down
docker-compose logs -f backend
docker-compose restart backend
```

#### 生产环境命令
```bash
# 生产环境部署
docker-compose -f docker-compose.prod.yml up -d
docker-compose -f docker-compose.prod.yml down
docker-compose -f docker-compose.prod.yml logs -f backend
docker-compose -f docker-compose.prod.yml restart nginx
```

#### 通用命令
```bash
# 构建镜像
docker-compose build

# 进入容器
docker-compose exec backend sh

# 查看容器状态
docker ps

# 查看镜像
docker images
```

### 数据库准备

由于使用服务器现有的 MySQL 和 Redis，需要手动准备数据库：

#### MySQL 数据库设置
```sql
-- 连接到 MySQL
mysql -u root -p

-- 创建数据库
CREATE DATABASE IF NOT EXISTS melody_cure CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建用户并授权
CREATE USER IF NOT EXISTS 'melody_cure'@'%' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON melody_cure.* TO 'melody_cure'@'%';
FLUSH PRIVILEGES;

-- 设置时区
SET time_zone = '+08:00';
```

#### Redis 设置
确保 Redis 服务正在运行，并且如果设置了密码，请在环境变量中正确配置。

```bash
# 检查 Redis 状态
redis-cli ping

# 如果有密码
redis-cli -a your_password ping
```

### Docker 部署

```dockerfile
FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config
CMD ["./main"]
```

### 环境变量

- `DB_HOST` - 数据库主机
- `DB_PORT` - 数据库端口
- `DB_USER` - 数据库用户名
- `DB_PASSWORD` - 数据库密码
- `DB_NAME` - 数据库名称
- `REDIS_HOST` - Redis主机
- `REDIS_PORT` - Redis端口
- `JWT_SECRET` - JWT密钥

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 联系方式

如有问题或建议，请通过以下方式联系：

- 项目地址: [https://github.com/your-username/Melody-Cure-Backend](https://github.com/your-username/Melody-Cure-Backend)
- 问题反馈: [Issues](https://github.com/your-username/Melody-Cure-Backend/issues)
