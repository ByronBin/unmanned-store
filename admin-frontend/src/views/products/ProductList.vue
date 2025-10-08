<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center;">
          <span style="font-size: 18px; font-weight: bold;">商品管理</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增商品
          </el-button>
        </div>
      </template>
      
      <!-- 搜索和筛选区域 -->
      <div style="margin-bottom: 20px;">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-input
              v-model="searchForm.keyword"
              placeholder="请输入商品名称或编码"
              clearable
              @keyup.enter="handleSearch"
            >
              <template #prefix>
                <el-icon><Search /></el-icon>
              </template>
            </el-input>
          </el-col>
          <el-col :span="4">
            <el-select v-model="searchForm.categoryId" placeholder="选择分类" clearable>
              <el-option
                v-for="category in categoryOptions"
                :key="category.value"
                :label="category.label"
                :value="category.value"
              />
            </el-select>
          </el-col>
          <el-col :span="4">
            <el-select v-model="searchForm.status" placeholder="商品状态" clearable>
              <el-option label="全部" value="" />
              <el-option label="上架" value="active" />
              <el-option label="下架" value="inactive" />
              <el-option label="草稿" value="draft" />
            </el-select>
          </el-col>
          <el-col :span="6">
            <el-button type="primary" @click="handleSearch">
              <el-icon><Search /></el-icon>
              搜索
            </el-button>
            <el-button @click="handleReset">
              <el-icon><Refresh /></el-icon>
              重置
            </el-button>
          </el-col>
        </el-row>
      </div>

      <!-- 批量操作 -->
      <div v-if="selectedProducts.length > 0" style="margin-bottom: 15px;">
        <el-alert
          :title="`已选择 ${selectedProducts.length} 个商品`"
          type="info"
          show-icon
          :closable="false"
        >
          <template #default>
            <el-button size="small" @click="handleBatchStatus('active')">批量上架</el-button>
            <el-button size="small" @click="handleBatchStatus('inactive')">批量下架</el-button>
            <el-button size="small" type="danger" @click="handleBatchDelete">批量删除</el-button>
          </template>
        </el-alert>
      </div>

      <!-- 商品表格 -->
      <el-table 
        :data="productList" 
        border 
        v-loading="loading"
        @selection-change="handleSelectionChange"
        row-key="id"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="images" label="商品图片" width="100">
          <template #default="{ row }">
            <el-image
              v-if="row.images && row.images.length > 0"
              :src="row.images[0]"
              :preview-src-list="row.images"
              style="width: 60px; height: 60px; border-radius: 4px;"
              fit="cover"
            />
            <span v-else style="color: #999;">暂无图片</span>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="商品名称" min-width="200">
          <template #default="{ row }">
            <div>
              <div style="font-weight: bold; margin-bottom: 4px;">{{ row.name }}</div>
              <div style="color: #666; font-size: 12px;">编码: {{ row.code }}</div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="category" label="分类" width="120">
          <template #default="{ row }">
            <el-tag v-if="row.category">{{ row.category.name }}</el-tag>
            <span v-else style="color: #999;">未分类</span>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="价格" width="100">
          <template #default="{ row }">
            <span style="color: #e6a23c; font-weight: bold;">¥{{ row.price }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="stock" label="库存" width="80">
          <template #default="{ row }">
            <span :style="{ color: row.stock <= 10 ? '#f56c6c' : '#67c23a' }">
              {{ row.stock || 0 }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag 
              :type="getStatusType(row.status)"
              size="small"
            >
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button 
              size="small" 
              :type="row.status === 'active' ? 'warning' : 'success'"
              @click="handleStatusChange(row)"
            >
              {{ row.status === 'active' ? '下架' : '上架' }}
            </el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div style="margin-top: 20px; display: flex; justify-content: center;">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Search, Refresh } from '@element-plus/icons-vue'
import { 
  getProductList, 
  searchProducts, 
  updateProductStatus, 
  deleteProduct,
  getCategoryTree 
} from '@/api/product'

// 响应式数据
const loading = ref(false)
const productList = ref([])
const selectedProducts = ref([])
const categoryOptions = ref([])

// 搜索表单
const searchForm = reactive({
  keyword: '',
  categoryId: '',
  status: ''
})

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

// 计算属性
const getStatusType = (status: string) => {
  const statusMap = {
    active: 'success',
    inactive: 'danger',
    draft: 'info'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status: string) => {
  const statusMap = {
    active: '上架',
    inactive: '下架',
    draft: '草稿'
  }
  return statusMap[status] || '未知'
}

// 方法
const loadProducts = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.pageSize,
      ...searchForm
    }
    
    let response
    if (searchForm.keyword) {
      response = await searchProducts(searchForm.keyword, params)
    } else {
      response = await getProductList(params)
    }
    
    productList.value = response.data.data || []
    pagination.total = response.data.total || 0
  } catch (error) {
    ElMessage.error('加载商品列表失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const loadCategories = async () => {
  try {
    const response = await getCategoryTree()
    categoryOptions.value = formatCategoryOptions(response.data.data || [])
  } catch (error) {
    console.error('加载分类失败:', error)
  }
}

const formatCategoryOptions = (categories: any[]): any[] => {
  const options: any[] = []
  const formatCategory = (category: any) => {
    options.push({
      label: category.name,
      value: category.id
    })
    if (category.children && category.children.length > 0) {
      category.children.forEach(formatCategory)
    }
  }
  categories.forEach(formatCategory)
  return options
}

const handleSearch = () => {
  pagination.page = 1
  loadProducts()
}

const handleReset = () => {
  searchForm.keyword = ''
  searchForm.categoryId = ''
  searchForm.status = ''
  pagination.page = 1
  loadProducts()
}

const handleAdd = () => {
  // TODO: 跳转到新增商品页面
  ElMessage.info('跳转到新增商品页面')
}

const handleEdit = (row: any) => {
  // TODO: 跳转到编辑商品页面
  ElMessage.info(`编辑商品: ${row.name}`)
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除商品"${row.name}"吗？此操作不可撤销。`,
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
    
    await deleteProduct(row.id)
    ElMessage.success('删除成功')
    loadProducts()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
      console.error(error)
    }
  }
}

const handleStatusChange = async (row: any) => {
  try {
    const newStatus = row.status === 'active' ? 'inactive' : 'active'
    await updateProductStatus(row.id, newStatus)
    ElMessage.success(`${newStatus === 'active' ? '上架' : '下架'}成功`)
    loadProducts()
  } catch (error) {
    ElMessage.error('状态修改失败')
    console.error(error)
  }
}

const handleSelectionChange = (selection: any[]) => {
  selectedProducts.value = selection
}

const handleBatchStatus = async (status: string) => {
  try {
    await ElMessageBox.confirm(
      `确定要批量${status === 'active' ? '上架' : '下架'} ${selectedProducts.value.length} 个商品吗？`,
      '批量操作确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
    
    const promises = selectedProducts.value.map((product: any) => 
      updateProductStatus(product.id, status)
    )
    await Promise.all(promises)
    
    ElMessage.success('批量操作成功')
    loadProducts()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('批量操作失败')
      console.error(error)
    }
  }
}

const handleBatchDelete = async () => {
  try {
    await ElMessageBox.confirm(
      `确定要删除 ${selectedProducts.value.length} 个商品吗？此操作不可撤销。`,
      '批量删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
    
    const promises = selectedProducts.value.map((product: any) => 
      deleteProduct(product.id)
    )
    await Promise.all(promises)
    
    ElMessage.success('批量删除成功')
    loadProducts()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('批量删除失败')
      console.error(error)
    }
  }
}

const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  loadProducts()
}

const handleCurrentChange = (page: number) => {
  pagination.page = page
  loadProducts()
}

const formatDate = (date: string) => {
  if (!date) return '-'
  return new Date(date).toLocaleString('zh-CN')
}

// 生命周期
onMounted(() => {
  loadProducts()
  loadCategories()
})
</script>
