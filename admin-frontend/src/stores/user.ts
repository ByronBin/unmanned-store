import { defineStore } from 'pinia'
import { ref } from 'vue'
import { login as apiLogin, getUserInfo } from '@/api/auth'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref<any>(null)

  const setToken = (newToken: string) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  const setUserInfo = (info: any) => {
    userInfo.value = info
  }

  const login = async (username: string, password: string) => {
    const res = await apiLogin(username, password)
    setToken(res.access_token)
    return res
  }

  const logout = () => {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
  }

  const fetchUserInfo = async () => {
    const info = await getUserInfo()
    setUserInfo(info)
    return info
  }

  return {
    token,
    userInfo,
    setToken,
    setUserInfo,
    login,
    logout,
    fetchUserInfo
  }
})
