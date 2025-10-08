# 无人超市后端服务

基于 Go + Gin + PostgreSQL + Redis 的无人超市管理系统后端服务。

## 技术栈

- **语言**: Go 1.21+
- **框架**: Gin (HTTP), GORM (ORM)
- **数据库**: PostgreSQL 15+
- **缓存**: Redis 7+
- **认证**: JWT
- **日志**: Zap

## 项目结构

```
backend/
├── cmd/                    # 入口程序
│   ├── api/               # HTTP API服务
│   └── worker/            # 异步任务处理
├── internal/              # 私有应用代码
│   ├── domain/            # 领域模型
│   ├── service/           # 业务逻辑
│   ├── repository/        # 数据访问
│   ├── handler/           # HTTP处理器
│   └── middleware/        # 中间件
├── pkg/                   # 公共库
│   ├── auth/
│   ├── logger/
│   └── utils/
├── migrations/            # 数据库迁移
├── config/                # 配置文件
└── go.mod
```

## 快速开始

### 1. 安装依赖

```bash
cd backend
go mod download
```

### 2. 配置数据库

创建 PostgreSQL 数据库：

```sql
CREATE DATABASE unmanned_store;
```

运行迁移脚本：

```bash
psql -U postgres -d unmanned_store -f migrations/001_initial_schema.sql
```

### 3. 配置环境

修改 `config/config.yaml`：

```yaml
database:
  host: localhost
  port: 5432
  user: postgres
  password: your_password
  dbname: unmanned_store

redis:
  host: localhost
  port: 6379
```

### 4. 运行服务

```bash
go run cmd/api/main.go
```

服务将在 `http://localhost:8080` 启动

### 5. 测试API

默认管理员账号：
- 用户名: `admin`
- 密码: `admin123`

登录获取token：

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

## API文档

启动服务后访问：
- Swagger UI: http://localhost:8080/swagger/index.html

## 核心模块

- **认证授权**: JWT + RBAC
- **门店管理**: 多门店数据隔离
- **商品管理**: 商品、SKU、分类
- **库存管理**: 实时库存、预警
- **订单系统**: 订单创建、支付、退款
- **会员营销**: 会员、积分、优惠券
- **门禁系统**: 扫码开门、黑名单
- **监控系统**: 视频监控、告警
- **财务分析**: 报表、统计

## 开发指南

### 添加新模块

1. 在 `internal/domain` 添加模型
2. 在 `internal/repository` 添加数据访问层
3. 在 `internal/service` 添加业务逻辑
4. 在 `internal/handler` 添加HTTP处理器
5. 在 `cmd/api/main.go` 注册路由

### 数据库迁移

添加新的迁移文件到 `migrations/` 目录，按顺序命名：
- `001_initial_schema.sql`
- `002_add_xxx_table.sql`
- ...

## License

MIT
