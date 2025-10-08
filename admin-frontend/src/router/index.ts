import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/',
    component: () => import('@/layouts/MainLayout.vue'),
    redirect: '/dashboard',
    meta: { requiresAuth: true },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue'),
        meta: { title: '仪表盘' }
      },
      {
        path: 'stores',
        name: 'Stores',
        component: () => import('@/views/stores/StoreList.vue'),
        meta: { title: '门店管理' }
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('@/views/products/ProductList.vue'),
        meta: { title: '商品管理' }
      },
      {
        path: 'categories',
        name: 'Categories',
        component: () => import('@/views/products/CategoryList.vue'),
        meta: { title: '分类管理' }
      },
      {
        path: 'inventory',
        name: 'Inventory',
        component: () => import('@/views/inventory/InventoryList.vue'),
        meta: { title: '库存管理' }
      },
      {
        path: 'orders',
        name: 'Orders',
        component: () => import('@/views/orders/OrderList.vue'),
        meta: { title: '订单管理' }
      },
      {
        path: 'members',
        name: 'Members',
        component: () => import('@/views/members/MemberList.vue'),
        meta: { title: '会员管理' }
      },
      {
        path: 'coupons',
        name: 'Coupons',
        component: () => import('@/views/marketing/CouponList.vue'),
        meta: { title: '优惠券管理' }
      },
      {
        path: 'access',
        name: 'Access',
        component: () => import('@/views/access/AccessLog.vue'),
        meta: { title: '门禁记录' }
      },
      {
        path: 'monitoring',
        name: 'Monitoring',
        component: () => import('@/views/monitoring/MonitoringView.vue'),
        meta: { title: '监控中心' }
      },
      {
        path: 'finance',
        name: 'Finance',
        component: () => import('@/views/finance/FinanceReport.vue'),
        meta: { title: '财务报表' }
      },
      {
        path: 'analytics',
        name: 'Analytics',
        component: () => import('@/views/analytics/AnalyticsView.vue'),
        meta: { title: '数据分析' }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  if (to.meta.requiresAuth !== false && !userStore.token) {
    next('/login')
  } else if (to.path === '/login' && userStore.token) {
    next('/')
  } else {
    next()
  }
})

export default router
