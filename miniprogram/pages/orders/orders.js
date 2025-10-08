// pages/orders/orders.js
const app = getApp()

Page({
  data: {
    currentTab: 'all',
    orders: []
  },

  onShow() {
    this.loadOrders()
  },

  switchTab(e) {
    const tab = e.currentTarget.dataset.tab
    this.setData({
      currentTab: tab
    }, () => {
      this.loadOrders()
    })
  },

  loadOrders() {
    const params = {
      page: 1,
      page_size: 20
    }
    if (this.data.currentTab !== 'all') {
      params.status = this.data.currentTab
    }

    app.request({
      url: '/orders',
      data: params
    }).then(res => {
      const orders = (res.data || []).map(order => ({
        ...order,
        statusText: this.getStatusText(order.status)
      }))
      this.setData({
        orders
      })
    })
  },

  getStatusText(status) {
    const statusMap = {
      'pending': '待支付',
      'paid': '已支付',
      'completed': '已完成',
      'cancelled': '已取消',
      'refunded': '已退款'
    }
    return statusMap[status] || status
  },

  goToDetail(e) {
    const id = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/order-detail/order-detail?id=${id}`
    })
  },

  pay(e) {
    const id = e.currentTarget.dataset.id
    // TODO: 发起支付
    wx.showLoading({ title: '支付中...' })
  }
})
