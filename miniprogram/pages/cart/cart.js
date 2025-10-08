// pages/cart/cart.js
Page({
  data: {
    cartItems: [],
    allChecked: false,
    totalPrice: 0
  },

  onShow() {
    this.loadCart()
  },

  loadCart() {
    // TODO: 从缓存或后端加载购物车数据
    this.setData({
      cartItems: []
    })
    this.calculateTotal()
  },

  toggleAll() {
    const allChecked = !this.data.allChecked
    const cartItems = this.data.cartItems.map(item => ({
      ...item,
      checked: allChecked
    }))
    this.setData({
      allChecked,
      cartItems
    })
    this.calculateTotal()
  },

  increaseQty(e) {
    const id = e.currentTarget.dataset.id
    // TODO: 增加数量
  },

  decreaseQty(e) {
    const id = e.currentTarget.dataset.id
    // TODO: 减少数量
  },

  calculateTotal() {
    const total = this.data.cartItems
      .filter(item => item.checked)
      .reduce((sum, item) => sum + item.price * item.quantity, 0)
    this.setData({
      totalPrice: total.toFixed(2)
    })
  },

  checkout() {
    const selectedItems = this.data.cartItems.filter(item => item.checked)
    if (selectedItems.length === 0) {
      wx.showToast({
        title: '请选择商品',
        icon: 'none'
      })
      return
    }
    // TODO: 跳转到结算页面
  }
})
