// pages/index/index.js
const app = getApp()

Page({
  data: {
    banners: [
      'https://via.placeholder.com/750x300',
      'https://via.placeholder.com/750x300',
      'https://via.placeholder.com/750x300'
    ],
    hotProducts: []
  },

  onLoad() {
    this.loadHotProducts()
  },

  loadHotProducts() {
    app.request({
      url: '/products',
      data: {
        page: 1,
        page_size: 6
      }
    }).then(res => {
      this.setData({
        hotProducts: res.data || []
      })
    }).catch(err => {
      console.error('加载商品失败', err)
    })
  },

  goToStore() {
    wx.navigateTo({
      url: '/pages/store/store'
    })
  },

  goToScan() {
    wx.scanCode({
      success: (res) => {
        console.log('扫码结果', res)
        // TODO: 处理扫码结果
      }
    })
  },

  goToOrders() {
    wx.switchTab({
      url: '/pages/orders/orders'
    })
  },

  goToCoupons() {
    wx.navigateTo({
      url: '/pages/coupons/coupons'
    })
  },

  goToProducts() {
    wx.switchTab({
      url: '/pages/products/products'
    })
  },

  goToProductDetail(e) {
    const id = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/product-detail/product-detail?id=${id}`
    })
  },

  addToCart(e) {
    const id = e.currentTarget.dataset.id
    // TODO: 添加到购物车
    wx.showToast({
      title: '已添加到购物车',
      icon: 'success'
    })
  }
})
