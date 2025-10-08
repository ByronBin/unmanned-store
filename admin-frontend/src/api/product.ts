import request from './request'

// 商品管理API
export const getProductList = (params?: any) => {
  return request({
    url: '/products',
    method: 'get',
    params
  })
}

export const searchProducts = (keyword: string, params?: any) => {
  return request({
    url: '/products/search',
    method: 'get',
    params: { keyword, ...params }
  })
}

export const getHotProducts = (limit = 10) => {
  return request({
    url: '/products/hot',
    method: 'get',
    params: { limit }
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

export const updateProductStatus = (id: string, status: string) => {
  return request({
    url: `/products/${id}/status`,
    method: 'put',
    data: { status }
  })
}

export const deleteProduct = (id: string) => {
  return request({
    url: `/products/${id}`,
    method: 'delete'
  })
}

// SKU管理API
export const createSKU = (productId: string, data: any) => {
  return request({
    url: `/products/${productId}/skus`,
    method: 'post',
    data
  })
}

export const updateSKU = (id: string, data: any) => {
  return request({
    url: `/products/skus/${id}`,
    method: 'put',
    data
  })
}

export const deleteSKU = (id: string) => {
  return request({
    url: `/products/skus/${id}`,
    method: 'delete'
  })
}

export const getSKU = (id: string) => {
  return request({
    url: `/products/skus/${id}`,
    method: 'get'
  })
}

// 分类管理API
export const getCategoryList = () => {
  return request({
    url: '/categories',
    method: 'get'
  })
}

export const getCategoryTree = () => {
  return request({
    url: '/categories/tree',
    method: 'get'
  })
}

export const getCategory = (id: string) => {
  return request({
    url: `/categories/${id}`,
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

export const updateCategory = (id: string, data: any) => {
  return request({
    url: `/categories/${id}`,
    method: 'put',
    data
  })
}

export const deleteCategory = (id: string) => {
  return request({
    url: `/categories/${id}`,
    method: 'delete'
  })
}

// 库存管理API
export const getInventoryByStore = (storeId: string, params?: any) => {
  return request({
    url: '/inventory',
    method: 'get',
    params: { store_id: storeId, ...params }
  })
}

export const getInventoryBySKU = (skuId: string, storeId?: string) => {
  return request({
    url: `/inventory/sku/${skuId}`,
    method: 'get',
    params: storeId ? { store_id: storeId } : {}
  })
}

export const getInventoryByProduct = (productId: string, storeId?: string) => {
  return request({
    url: `/inventory/product/${productId}`,
    method: 'get',
    params: storeId ? { store_id: storeId } : {}
  })
}

export const adjustInventory = (data: any) => {
  return request({
    url: '/inventory/adjust',
    method: 'post',
    data
  })
}

export const stockIn = (data: any) => {
  return request({
    url: '/inventory/stock-in',
    method: 'post',
    data
  })
}

export const stockOut = (data: any) => {
  return request({
    url: '/inventory/stock-out',
    method: 'post',
    data
  })
}

export const getLowStockItems = (storeId?: string, threshold = 10) => {
  return request({
    url: '/inventory/low-stock',
    method: 'get',
    params: storeId ? { store_id: storeId, threshold } : { threshold }
  })
}

export const getInventoryLogs = (params?: any) => {
  return request({
    url: '/inventory/logs',
    method: 'get',
    params
  })
}

// 库存盘点API
export const createInventoryCount = (data: any) => {
  return request({
    url: '/inventory/counts',
    method: 'post',
    data
  })
}

export const getInventoryCounts = (storeId: string, status?: string) => {
  return request({
    url: '/inventory/counts',
    method: 'get',
    params: { store_id: storeId, status }
  })
}

export const submitInventoryCount = (id: string, data: any) => {
  return request({
    url: `/inventory/counts/${id}/submit`,
    method: 'post',
    data
  })
}
