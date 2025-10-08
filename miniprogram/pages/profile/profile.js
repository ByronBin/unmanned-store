// pages/profile/profile.js
const app = getApp()

Page({
  data: {
    userInfo: {},
    isLoggedIn: false
  },

  onShow() {
    this.loadUserInfo()
  },

  loadUserInfo() {
    if (!app.checkLoginStatus()) {
      return
    }

    app.request({
      url: '/members/profile'
    }).then(res => {
      this.setData({
        userInfo: res,
        isLoggedIn: true
      })
    }).catch(() => {
      this.setData({
        isLoggedIn: false
      })
    })
  },

  openSettings() {
    wx.navigateTo({
      url: '/pages/settings/settings'
    })
  },

  logout() {
    wx.showModal({
      title: '提示',
      content: '确定要退出登录吗？',
      success: (res) => {
        if (res.confirm) {
          app.globalData.token = null
          wx.removeStorageSync('token')
          this.setData({
            isLoggedIn: false,
            userInfo: {}
          })
        }
      }
    })
  }
})
