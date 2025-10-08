// app.js
App({
  globalData: {
    userInfo: null,
    token: null,
    apiUrl: 'http://localhost:8080/api/v1'
  },

  onLaunch() {
    // 获取本地存储的token
    const token = wx.getStorageSync('token')
    if (token) {
      this.globalData.token = token
    }

    // 检查登录状态
    this.checkLoginStatus()
  },

  checkLoginStatus() {
    if (!this.globalData.token) {
      return false
    }
    // TODO: 验证token有效性
    return true
  },

  login() {
    return new Promise((resolve, reject) => {
      wx.login({
        success: (res) => {
          if (res.code) {
            // 发送 res.code 到后台换取 token
            this.request({
              url: '/auth/wechat/login',
              method: 'POST',
              data: {
                code: res.code
              }
            }).then(data => {
              this.globalData.token = data.token
              wx.setStorageSync('token', data.token)
              resolve(data)
            }).catch(reject)
          } else {
            reject(res.errMsg)
          }
        },
        fail: reject
      })
    })
  },

  request(options) {
    return new Promise((resolve, reject) => {
      wx.request({
        url: this.globalData.apiUrl + options.url,
        method: options.method || 'GET',
        data: options.data || {},
        header: {
          'Content-Type': 'application/json',
          'Authorization': this.globalData.token ? `Bearer ${this.globalData.token}` : ''
        },
        success: (res) => {
          if (res.statusCode === 200) {
            resolve(res.data)
          } else {
            wx.showToast({
              title: res.data.error || '请求失败',
              icon: 'none'
            })
            reject(res.data)
          }
        },
        fail: (err) => {
          wx.showToast({
            title: '网络错误',
            icon: 'none'
          })
          reject(err)
        }
      })
    })
  }
})
