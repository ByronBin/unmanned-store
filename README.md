# 无人值守24小时超市系统

一个完整的无人超市管理系统，支持10店规模运营，包含后端API、管理后台、微信小程序、门禁监控等完整功能。

## 系统架构

```
无人超市系统
├── 后端服务 (Go + Gin + PostgreSQL + Redis)
├── 管理后台 (Vue 3 + TypeScript + Element Plus)
├── 微信小程序 (原生小程序开发)
└── 硬件对接 (门禁、监控、电子标签)
```

## 技术栈

### 后端
- **语言**: Go 1.21+
- **框架**: Gin (HTTP), GORM (ORM)
- **数据库**: PostgreSQL 15+
- **缓存**: Redis 7+
- **消息队列**: RabbitMQ
- **认证**: JWT + RBAC

### 管理后台
- **框架**: Vue 3 + TypeScript
- **UI组件**: Element Plus
- **状态管理**: Pinia
- **构建工具**: Vite

### 微信小程序
- **开发方式**: 原生小程序开发
- **UI风格**: 简洁现代

## 核心功能

### 后端系统
- ✅ 用户认证与授权（JWT + RBAC）
- ✅ 多租户门店管理（10店数据隔离）
- ✅ 商品管理（商品、SKU、分类）
- ✅ 库存管理（实时库存、预警、调拨）
- ✅ 订单系统（订单创建、支付、退款）
- ✅ 支付对接（微信支付、支付宝）
- ✅ 会员营销（会员、积分、优惠券）
- ✅ 门禁系统（扫码开门、黑名单）
- ✅ 监控系统（视频监控、设备告警）
- ✅ 财务报表（日报、月报、汇总）
- ✅ 数据分析（销售统计、热销商品）

### 管理后台
- 🎯 仪表盘（数据概览）
- 🏪 门店管理
- 📦 商品与库存管理
- 📝 订单管理
- 👥 会员管理
- 🎫 优惠券管理
- 🚪 门禁记录
- 📹 监控中心
- 💰 财务报表
- 📊 数据分析

### 微信小程序
- 🏠 首页展示
- 🛍️ 商品浏览与搜索
- 🛒 购物车管理
- 📋 订单管理
- 🚪 扫码开门
- 👤 个人中心
- 🏪 门店选择

## 项目结构

```
yyj/
├── backend/                    # Go后端
│   ├── cmd/                    # 入口程序
│   ├── internal/              # 私有应用代码
│   │   ├── domain/            # 领域模型
│   │   ├── service/           # 业务逻辑
│   │   ├── repository/        # 数据访问
│   │   ├── handler/           # HTTP处理器
│   │   └── middleware/        # 中间件
│   ├── pkg/                   # 公共库
│   ├── migrations/            # 数据库迁移
│   └── config/                # 配置文件
├── admin-frontend/            # Vue3管理后台
│   ├── src/
│   │   ├── views/            # 页面
│   │   ├── components/       # 组件
│   │   ├── api/              # API接口
│   │   ├── stores/           # Pinia状态
│   │   └── router/           # 路由
│   └── package.json
├── miniprogram/              # 微信小程序
│   ├── pages/                # 页面
│   ├── components/           # 组件
│   └── app.json
├── docker/                   # Docker配置
│   ├── docker-compose.yml
│   ├── Dockerfile.backend
│   └── Dockerfile.frontend
└── docs/                     # 文档
```

## 快速开始

### 使用Docker Compose（推荐）

```bash
cd yyj/docker
docker-compose up -d
```

服务访问地址：
- 后端API: http://localhost:8080
- 管理后台: http://localhost:3000
- PostgreSQL: localhost:5432
- Redis: localhost:6379
- RabbitMQ管理界面: http://localhost:15672

### 手动启动

#### 1. 后端服务

```bash
# 安装依赖
cd backend
go mod download

# 创建数据库
createdb unmanned_store
psql -d unmanned_store -f migrations/001_initial_schema.sql

# 修改配置
cp config/config.yaml config/config.local.yaml
# 编辑 config.local.yaml 修改数据库等配置

# 启动服务
go run cmd/api/main.go
```

#### 2. 管理后台

```bash
cd admin-frontend
npm install
npm run dev
```

访问: http://localhost:3000

#### 3. 微信小程序

1. 使用微信开发者工具打开 `miniprogram` 目录
2. 修改 `project.config.json` 中的 `appid`
3. 修改 `app.js` 中的 `apiUrl` 为后端地址
4. 编译预览

## 默认账号

- **用户名**: admin
- **密码**: admin123

## 硬件对接

### 支持的硬件设备

1. **门禁设备**
   - 海康威视、大华、萤石
   - 支持二维码开门、人脸识别

2. **监控摄像头**
   - 支持RTSP视频流
   - 移动侦测告警

3. **电子价签**
   - 汉朔、京东方
   - 远程价格更新

4. **自助收银机**
   - 扫码识别、支付、打印

## 数据库设计

系统采用PostgreSQL作为主数据库，包含以下核心表：

- stores - 门店表
- users - 用户表
- products - 商品表
- product_skus - SKU表
- inventory - 库存表
- orders - 订单表
- payments - 支付记录
- coupons - 优惠券表
- access_logs - 门禁日志
- monitoring_devices - 监控设备

详见: `backend/migrations/001_initial_schema.sql`

## API文档

启动后端服务后，访问 Swagger 文档：
http://localhost:8080/swagger/index.html

## 开发路线图

- [x] 阶段一：基础框架搭建
- [ ] 阶段二：核心购物流程
- [ ] 阶段三：门禁与监控
- [ ] 阶段四：会员与营销
- [ ] 阶段五：多门店管理
- [ ] 阶段六：财务与数据分析
- [ ] 阶段七：测试与优化

## 贡献指南

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 联系方式

如有问题，请提交 Issue。
