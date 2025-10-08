// pages/store/store.js
const app = getApp()

Page({
  data: {
    stores: []
  },

  onLoad() {
    this.loadStores()
  },

  loadStores() {
    app.request({
      url: '/stores'
    }).then(res => {
      this.setData({
        stores: res.data || []
      })
    })
  },

  onSearch(e) {
    // TODO: 搜索门店
  },

  selectStore(e) {
    const id = e.currentTarget.dataset.id
    wx.navigateBack()
  }
})
