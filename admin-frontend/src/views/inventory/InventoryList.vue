<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center;">
          <span style="font-size: 18px; font-weight: bold;">库存管理</span>
          <div>
            <el-button type="success" @click="handleStockIn">
              <el-icon><Plus /></el-icon>
              批量入库
            </el-button>
            <el-button type="warning" @click="handleStockOut">
              <el-icon><Minus /></el-icon>
              批量出库
            </el-button>
            <el-button type="primary" @click="handleInventoryCount">
              <el-icon><List /></el-icon>
              库存盘点
            </el-button>
          </div>
        </div>
      </template>

      <!-- 搜索和筛选区域 -->
      <div style="margin-bottom: 20px;">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-select v-model="searchForm.storeId" placeholder="选择门店" clearable>
              <el-option
                v-for="store in storeOptions"
                :key="store.value"
                :label="store.label"
                :value="store.value"
              />
            </el-select>
          </el-col>
          <el-col :span="6">
            <el-input
              v-model="searchForm.keyword"
              placeholder="搜索商品名称或SKU"
              clearable
              @keyup.enter="handleSearch"
            >
              <template #prefix>
                <el-icon><Search /></el-icon>
              </template>
            </el-input>
          </el-col>
          <el-col :span="4">
            <el-select v-model="searchForm.lowStock" placeholder="库存状态" clearable>
              <el-option label="全部" value="" />
              <el-option label="正常" value="normal" />
              <el-option label="低库存" value="low" />
              <el-option label="缺货" value="out" />
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
            <el-button type="info" @click="handleViewLogs">
              <el-icon><Document /></el-icon>
              查看日志
            </el-button>
          </el-col>
        </el-row>
      </div>

      <!-- 库存统计卡片 -->
      <div style="margin-bottom: 20px;">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-number">{{ stats.totalProducts }}</div>
                <div class="stat-label">商品总数</div>
              </div>
              <el-icon class="stat-icon"><Box /></el-icon>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-number">{{ stats.lowStockCount }}</div>
                <div class="stat-label">低库存商品</div>
              </div>
              <el-icon class="stat-icon warning"><Warning /></el-icon>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-number">{{ stats.outOfStockCount }}</div>
                <div class="stat-label">缺货商品</div>
              </div>
              <el-icon class="stat-icon danger"><Close /></el-icon>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-number">{{ stats.totalValue.toFixed(2) }}</div>
                <div class="stat-label">库存总价值(元)</div>
              </div>
              <el-icon class="stat-icon success"><Money /></el-icon>
            </el-card>
          </el-col>
        </el-row>
      </div>

      <!-- 库存表格 -->
      <el-table 
        :data="inventoryList" 
        border 
        v-loading="loading"
        @selection-change="handleSelectionChange"
        row-key="id"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="sku" label="商品信息" min-width="250">
          <template #default="{ row }">
            <div class="product-info">
              <el-image
                v-if="row.sku && row.sku.product && row.sku.product.images && row.sku.product.images.length > 0"
                :src="row.sku.product.images[0]"
                style="width: 50px; height: 50px; border-radius: 4px; margin-right: 12px;"
                fit="cover"
              />
              <div class="product-details">
                <div class="product-name">{{ row.sku?.product?.name || '-' }}</div>
                <div class="sku-info">
                  <el-tag size="small" type="info">SKU: {{ row.sku?.sku_code || '-' }}</el-tag>
                  <span v-if="row.sku?.attributes" class="sku-attributes">
                    {{ formatSKUAttributes(row.sku.attributes) }}
                  </span>
                </div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="store" label="门店" width="120">
          <template #default="{ row }">
            <el-tag v-if="row.store">{{ row.store.name }}</el-tag>
            <span v-else style="color: #999;">-</span>
          </template>
        </el-table-column>
        <el-table-column prop="quantity" label="当前库存" width="120" align="center">
          <template #default="{ row }">
            <span 
              :style="{ 
                color: getStockColor(row.quantity, row.alert_quantity),
                fontWeight: 'bold',
                fontSize: '16px'
              }"
            >
              {{ row.quantity || 0 }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="alert_quantity" label="预警数量" width="100" align="center">
          <template #default="{ row }">
            {{ row.alert_quantity || 0 }}
          </template>
        </el-table-column>
        <el-table-column prop="unit_cost" label="成本价" width="100" align="center">
          <template #default="{ row }">
            <span style="color: #67c23a;">¥{{ row.unit_cost || 0 }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="total_value" label="库存价值" width="120" align="center">
          <template #default="{ row }">
            <span style="color: #e6a23c; font-weight: bold;">
              ¥{{ ((row.quantity || 0) * (row.unit_cost || 0)).toFixed(2) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="库存状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="getStockStatusType(row.quantity, row.alert_quantity)" size="small">
              {{ getStockStatusText(row.quantity, row.alert_quantity) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="updated_at" label="更新时间" width="160">
          <template #default="{ row }">
            {{ formatDate(row.updated_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" type="success" @click="handleQuickStockIn(row)">入库</el-button>
            <el-button size="small" type="warning" @click="handleQuickStockOut(row)">出库</el-button>
            <el-button size="small" @click="handleAdjust(row)">调整</el-button>
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

    <!-- 快速入库对话框 -->
    <el-dialog
      v-model="stockInDialogVisible"
      title="快速入库"
      width="500px"
    >
      <el-form :model="stockInForm" label-width="100px">
        <el-form-item label="商品">
          <el-input :value="selectedInventory?.sku?.product?.name" disabled />
        </el-form-item>
        <el-form-item label="SKU">
          <el-input :value="selectedInventory?.sku?.sku_code" disabled />
        </el-form-item>
        <el-form-item label="当前库存">
          <el-input :value="selectedInventory?.quantity" disabled />
        </el-form-item>
        <el-form-item label="入库数量" required>
          <el-input-number
            v-model="stockInForm.quantity"
            :min="1"
            :max="9999"
            style="width: 200px;"
          />
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="stockInForm.reason"
            type="textarea"
            placeholder="请输入入库原因"
            :rows="3"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="stockInDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmitStockIn" :loading="submitting">
          确认入库
        </el-button>
      </template>
    </el-dialog>

    <!-- 快速出库对话框 -->
    <el-dialog
      v-model="stockOutDialogVisible"
      title="快速出库"
      width="500px"
    >
      <el-form :model="stockOutForm" label-width="100px">
        <el-form-item label="商品">
          <el-input :value="selectedInventory?.sku?.product?.name" disabled />
        </el-form-item>
        <el-form-item label="SKU">
          <el-input :value="selectedInventory?.sku?.sku_code" disabled />
        </el-form-item>
        <el-form-item label="当前库存">
          <el-input :value="selectedInventory?.quantity" disabled />
        </el-form-item>
        <el-form-item label="出库数量" required>
          <el-input-number
            v-model="stockOutForm.quantity"
            :min="1"
            :max="selectedInventory?.quantity || 0"
            style="width: 200px;"
          />
        </el-form-item>
        <el-form-item label="出库原因" required>
          <el-select v-model="stockOutForm.reason" placeholder="请选择出库原因">
            <el-option label="销售出库" value="销售出库" />
            <el-option label="损耗出库" value="损耗出库" />
            <el-option label="调拨出库" value="调拨出库" />
            <el-option label="其他" value="其他" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="stockOutForm.remark"
            type="textarea"
            placeholder="请输入备注信息"
            :rows="3"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="stockOutDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmitStockOut" :loading="submitting">
          确认出库
        </el-button>
      </template>
    </el-dialog>

    <!-- 库存调整对话框 -->
    <el-dialog
      v-model="adjustDialogVisible"
      title="库存调整"
      width="500px"
    >
      <el-form :model="adjustForm" label-width="100px">
        <el-form-item label="商品">
          <el-input :value="selectedInventory?.sku?.product?.name" disabled />
        </el-form-item>
        <el-form-item label="SKU">
          <el-input :value="selectedInventory?.sku?.sku_code" disabled />
        </el-form-item>
        <el-form-item label="当前库存">
          <el-input :value="selectedInventory?.quantity" disabled />
        </el-form-item>
        <el-form-item label="调整数量" required>
          <el-input-number
            v-model="adjustForm.quantity"
            :min="-9999"
            :max="9999"
            style="width: 200px;"
          />
          <span style="margin-left: 10px; color: #666;">
            (正数为增加，负数为减少)
          </span>
        </el-form-item>
        <el-form-item label="调整原因" required>
          <el-input
            v-model="adjustForm.reason"
            type="textarea"
            placeholder="请输入调整原因"
            :rows="3"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="adjustDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmitAdjust" :loading="submitting">
          确认调整
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  Plus, 
  Minus, 
  List, 
  Search, 
  Refresh, 
  Document,
  Box,
  Warning,
  Close,
  Money
} from '@element-plus/icons-vue'
import { 
  getInventoryByStore,
  stockIn,
  stockOut,
  adjustInventory,
  getLowStockItems,
  getInventoryLogs
} from '@/api/product'

// 响应式数据
const loading = ref(false)
const submitting = ref(false)
const inventoryList = ref([])
const selectedInventories = ref([])
const storeOptions = ref([])
const selectedInventory = ref(null)

// 搜索表单
const searchForm = reactive({
  storeId: '',
  keyword: '',
  lowStock: ''
})

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

// 统计数据
const stats = reactive({
  totalProducts: 0,
  lowStockCount: 0,
  outOfStockCount: 0,
  totalValue: 0
})

// 对话框状态
const stockInDialogVisible = ref(false)
const stockOutDialogVisible = ref(false)
const adjustDialogVisible = ref(false)

// 表单数据
const stockInForm = reactive({
  quantity: 1,
  reason: ''
})

const stockOutForm = reactive({
  quantity: 1,
  reason: '',
  remark: ''
})

const adjustForm = reactive({
  quantity: 0,
  reason: ''
})

// 方法
const loadInventory = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.pageSize,
      ...searchForm
    }
    
    const response = await getInventoryByStore(searchForm.storeId || 'default', params)
    inventoryList.value = response.data.data || []
    pagination.total = response.data.total || 0
    
    // 计算统计数据
    calculateStats()
  } catch (error) {
    ElMessage.error('加载库存数据失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const loadLowStockItems = async () => {
  try {
    const response = await getLowStockItems(searchForm.storeId, 10)
    stats.lowStockCount = response.data.data?.length || 0
  } catch (error) {
    console.error('加载低库存数据失败:', error)
  }
}

const calculateStats = () => {
  stats.totalProducts = inventoryList.value.length
  stats.lowStockCount = inventoryList.value.filter(item => 
    item.quantity <= item.alert_quantity && item.quantity > 0
  ).length
  stats.outOfStockCount = inventoryList.value.filter(item => 
    item.quantity <= 0
  ).length
  stats.totalValue = inventoryList.value.reduce((sum, item) => 
    sum + ((item.quantity || 0) * (item.unit_cost || 0)), 0
  )
}

const handleSearch = () => {
  pagination.page = 1
  loadInventory()
}

const handleReset = () => {
  searchForm.storeId = ''
  searchForm.keyword = ''
  searchForm.lowStock = ''
  pagination.page = 1
  loadInventory()
}

const handleSelectionChange = (selection: any[]) => {
  selectedInventories.value = selection
}

const handleQuickStockIn = (row: any) => {
  selectedInventory.value = row
  stockInForm.quantity = 1
  stockInForm.reason = ''
  stockInDialogVisible.value = true
}

const handleQuickStockOut = (row: any) => {
  selectedInventory.value = row
  stockOutForm.quantity = 1
  stockOutForm.reason = ''
  stockOutForm.remark = ''
  stockOutDialogVisible.value = true
}

const handleAdjust = (row: any) => {
  selectedInventory.value = row
  adjustForm.quantity = 0
  adjustForm.reason = ''
  adjustDialogVisible.value = true
}

const handleSubmitStockIn = async () => {
  submitting.value = true
  try {
    await stockIn({
      store_id: selectedInventory.value.store_id,
      sku_id: selectedInventory.value.sku_id,
      quantity: stockInForm.quantity,
      reason: stockInForm.reason || '手动入库'
    })
    ElMessage.success('入库成功')
    stockInDialogVisible.value = false
    loadInventory()
  } catch (error) {
    ElMessage.error('入库失败')
    console.error(error)
  } finally {
    submitting.value = false
  }
}

const handleSubmitStockOut = async () => {
  submitting.value = true
  try {
    await stockOut({
      store_id: selectedInventory.value.store_id,
      sku_id: selectedInventory.value.sku_id,
      quantity: stockOutForm.quantity,
      reason: stockOutForm.reason,
      remark: stockOutForm.remark
    })
    ElMessage.success('出库成功')
    stockOutDialogVisible.value = false
    loadInventory()
  } catch (error) {
    ElMessage.error('出库失败')
    console.error(error)
  } finally {
    submitting.value = false
  }
}

const handleSubmitAdjust = async () => {
  submitting.value = true
  try {
    await adjustInventory({
      store_id: selectedInventory.value.store_id,
      sku_id: selectedInventory.value.sku_id,
      quantity: adjustForm.quantity,
      reason: adjustForm.reason
    })
    ElMessage.success('调整成功')
    adjustDialogVisible.value = false
    loadInventory()
  } catch (error) {
    ElMessage.error('调整失败')
    console.error(error)
  } finally {
    submitting.value = false
  }
}

const handleStockIn = () => {
  ElMessage.info('批量入库功能开发中')
}

const handleStockOut = () => {
  ElMessage.info('批量出库功能开发中')
}

const handleInventoryCount = () => {
  ElMessage.info('库存盘点功能开发中')
}

const handleViewLogs = () => {
  ElMessage.info('库存日志功能开发中')
}

const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  loadInventory()
}

const handleCurrentChange = (page: number) => {
  pagination.page = page
  loadInventory()
}

const getStockColor = (quantity: number, alertQuantity: number) => {
  if (quantity <= 0) return '#f56c6c' // 缺货 - 红色
  if (quantity <= alertQuantity) return '#e6a23c' // 低库存 - 橙色
  return '#67c23a' // 正常 - 绿色
}

const getStockStatusType = (quantity: number, alertQuantity: number) => {
  if (quantity <= 0) return 'danger'
  if (quantity <= alertQuantity) return 'warning'
  return 'success'
}

const getStockStatusText = (quantity: number, alertQuantity: number) => {
  if (quantity <= 0) return '缺货'
  if (quantity <= alertQuantity) return '低库存'
  return '正常'
}

const formatSKUAttributes = (attributes: any) => {
  if (!attributes) return ''
  return Object.entries(attributes).map(([key, value]) => `${key}:${value}`).join(', ')
}

const formatDate = (date: string) => {
  if (!date) return '-'
  return new Date(date).toLocaleString('zh-CN')
}

// 生命周期
onMounted(() => {
  loadInventory()
  loadLowStockItems()
})
</script>

<style scoped>
.stat-card {
  position: relative;
  overflow: hidden;
}

.stat-card .stat-content {
  position: relative;
  z-index: 2;
}

.stat-card .stat-icon {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 40px;
  opacity: 0.3;
  z-index: 1;
}

.stat-card .stat-icon.success {
  color: #67c23a;
}

.stat-card .stat-icon.warning {
  color: #e6a23c;
}

.stat-card .stat-icon.danger {
  color: #f56c6c;
}

.stat-number {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.product-info {
  display: flex;
  align-items: center;
}

.product-details {
  flex: 1;
}

.product-name {
  font-weight: bold;
  margin-bottom: 4px;
}

.sku-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.sku-attributes {
  font-size: 12px;
  color: #666;
}
</style>
