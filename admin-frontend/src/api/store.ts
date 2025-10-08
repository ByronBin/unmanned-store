import request from './request'

export const getStoreList = (params?: any) => {
  return request({
    url: '/stores',
    method: 'get',
    params
  })
}

export const getStore = (id: string) => {
  return request({
    url: `/stores/${id}`,
    method: 'get'
  })
}

export const createStore = (data: any) => {
  return request({
    url: '/stores',
    method: 'post',
    data
  })
}

export const updateStore = (id: string, data: any) => {
  return request({
    url: `/stores/${id}`,
    method: 'put',
    data
  })
}

export const deleteStore = (id: string) => {
  return request({
    url: `/stores/${id}`,
    method: 'delete'
  })
}
