// pages/products/products.js
const app = getApp()

Page({
  data: {
    categories: [],
    currentCategory: '',
    products: []
  },

  onLoad() {
    this.loadCategories()
    this.loadProducts()
  },

  loadCategories() {
    app.request({
      url: '/categories'
    }).then(res => {
      this.setData({
        categories: res || []
      })
    })
  },

  loadProducts() {
    const params = {
      page: 1,
      page_size: 20
    }
    if (this.data.currentCategory) {
      params.category_id = this.data.currentCategory
    }

    app.request({
      url: '/products',
      data: params
    }).then(res => {
      this.setData({
        products: res.data || []
      })
    })
  },

  selectCategory(e) {
    const id = e.currentTarget.dataset.id
    this.setData({
      currentCategory: id
    }, () => {
      this.loadProducts()
    })
  },

  onSearch(e) {
    // TODO: 实现搜索功能
    console.log('搜索', e.detail.value)
  },

  goToDetail(e) {
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
