import request from './request'

export const getProductList = (params?: any) => {
  return request({
    url: '/products',
    method: 'get',
    params
  })
}

export const getProduct = (id: string) => {
  return request({
    url: `/products/${id}`,
    method: 'get'
  })
}

export const createProduct = (data: any) => {
  return request({
    url: '/products',
    method: 'post',
    data
  })
}

export const updateProduct = (id: string, data: any) => {
  return request({
    url: `/products/${id}`,
    method: 'put',
    data
  })
}

export const deleteProduct = (id: string) => {
  return request({
    url: `/products/${id}`,
    method: 'delete'
  })
}

export const getCategoryList = () => {
  return request({
    url: '/categories',
    method: 'get'
  })
}

export const createCategory = (data: any) => {
  return request({
    url: '/categories',
    method: 'post',
    data
  })
}
