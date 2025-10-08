# 无人超市管理后台

基于 Vue 3 + TypeScript + Element Plus 的无人超市管理系统前端。

## 技术栈

- **框架**: Vue 3
- **语言**: TypeScript
- **UI组件**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router
- **HTTP客户端**: Axios
- **构建工具**: Vite

## 快速开始

### 1. 安装依赖

```bash
cd admin-frontend
npm install
```

### 2. 启动开发服务器

```bash
npm run dev
```

访问: http://localhost:3000

### 3. 构建生产版本

```bash
npm run build
```

## 项目结构

```
admin-frontend/
├── src/
│   ├── api/              # API接口
│   ├── assets/           # 静态资源
│   ├── components/       # 公共组件
│   ├── layouts/          # 布局组件
│   ├── router/           # 路由配置
│   ├── stores/           # Pinia状态
│   ├── views/            # 页面视图
│   ├── App.vue           # 根组件
│   └── main.ts           # 入口文件
├── public/               # 公共资源
├── index.html            # HTML模板
├── vite.config.ts        # Vite配置
└── package.json
```

## 功能模块

- 仪表盘：数据概览
- 门店管理：门店CRUD操作
- 商品管理：商品、分类管理
- 库存管理：入库、出库、调拨
- 订单管理：订单列表、详情
- 会员营销：会员、优惠券
- 门禁系统：门禁记录、黑名单
- 监控中心：视频监控、告警
- 财务报表：日报、月报
- 数据分析：销售统计、热销商品

## 默认登录

- 用户名: `admin`
- 密码: `admin123`

## License

MIT
