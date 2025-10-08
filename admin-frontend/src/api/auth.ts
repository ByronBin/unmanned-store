import request from './request'

export const login = (username: string, password: string) => {
  return request({
    url: '/auth/login',
    method: 'post',
    data: { username, password }
  })
}

export const register = (data: any) => {
  return request({
    url: '/auth/register',
    method: 'post',
    data
  })
}

export const getUserInfo = () => {
  return request({
    url: '/members/profile',
    method: 'get'
  })
}
