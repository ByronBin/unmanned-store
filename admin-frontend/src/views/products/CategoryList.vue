<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center;">
          <span style="font-size: 18px; font-weight: bold;">分类管理</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            新增分类
          </el-button>
        </div>
      </template>

      <!-- 操作工具栏 -->
      <div style="margin-bottom: 20px;">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-input
              v-model="searchKeyword"
              placeholder="搜索分类名称"
              clearable
              @input="handleSearch"
            >
              <template #prefix>
                <el-icon><Search /></el-icon>
              </template>
            </el-input>
          </el-col>
          <el-col :span="8">
            <el-select v-model="statusFilter" placeholder="状态筛选" clearable @change="handleSearch">
              <el-option label="全部" value="" />
              <el-option label="启用" value="active" />
              <el-option label="禁用" value="inactive" />
            </el-select>
          </el-col>
          <el-col :span="8">
            <el-button @click="handleRefresh">
              <el-icon><Refresh /></el-icon>
              刷新
            </el-button>
            <el-button @click="handleExpandAll">
              <el-icon><ArrowDown /></el-icon>
              展开全部
            </el-button>
            <el-button @click="handleCollapseAll">
              <el-icon><ArrowUp /></el-icon>
              收起全部
            </el-button>
          </el-col>
        </el-row>
      </div>

      <!-- 分类树 -->
      <el-tree
        ref="treeRef"
        :data="filteredCategoryTree"
        :props="{ label: 'name', children: 'children' }"
        :expand-on-click-node="false"
        :default-expand-all="false"
        node-key="id"
        v-loading="loading"
        @node-click="handleNodeClick"
      >
        <template #default="{ node, data }">
          <div class="tree-node">
            <div class="node-content">
              <el-icon v-if="data.icon" class="category-icon">
                <component :is="data.icon" />
              </el-icon>
              <span class="node-label">{{ data.name }}</span>
              <el-tag 
                :type="data.status === 'active' ? 'success' : 'danger'" 
                size="small"
                style="margin-left: 8px;"
              >
                {{ data.status === 'active' ? '启用' : '禁用' }}
              </el-tag>
              <span v-if="data.sort !== undefined" class="sort-number">
                排序: {{ data.sort }}
              </span>
            </div>
            <div class="node-actions">
              <el-button size="small" @click.stop="handleAddChild(data)">添加子分类</el-button>
              <el-button size="small" @click.stop="handleEdit(data)">编辑</el-button>
              <el-button 
                size="small" 
                :type="data.status === 'active' ? 'warning' : 'success'"
                @click.stop="handleToggleStatus(data)"
              >
                {{ data.status === 'active' ? '禁用' : '启用' }}
              </el-button>
              <el-button size="small" type="danger" @click.stop="handleDelete(data)">删除</el-button>
            </div>
          </div>
        </template>
      </el-tree>

      <!-- 分类编辑对话框 -->
      <el-dialog
        v-model="dialogVisible"
        :title="dialogTitle"
        width="600px"
        :before-close="handleDialogClose"
      >
        <el-form
          ref="formRef"
          :model="formData"
          :rules="formRules"
          label-width="100px"
        >
          <el-form-item label="分类名称" prop="name">
            <el-input v-model="formData.name" placeholder="请输入分类名称" />
          </el-form-item>
          
          <el-form-item label="父级分类" prop="parentId">
            <el-tree-select
              v-model="formData.parentId"
              :data="categoryTree"
              :props="{ label: 'name', value: 'id', children: 'children' }"
              placeholder="请选择父级分类（不选择则为顶级分类）"
              clearable
              check-strictly
            />
          </el-form-item>

          <el-form-item label="排序" prop="sort">
            <el-input-number
              v-model="formData.sort"
              :min="0"
              :max="999"
              placeholder="数字越小排序越靠前"
              style="width: 200px;"
            />
          </el-form-item>

          <el-form-item label="图标" prop="icon">
            <el-input v-model="formData.icon" placeholder="请输入图标名称（如：House、ShoppingBag）" />
          </el-form-item>

          <el-form-item label="状态" prop="status">
            <el-radio-group v-model="formData.status">
              <el-radio value="active">启用</el-radio>
              <el-radio value="inactive">禁用</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-form>

        <template #footer>
          <div class="dialog-footer">
            <el-button @click="handleDialogClose">取消</el-button>
            <el-button type="primary" @click="handleSubmit" :loading="submitting">
              {{ isEdit ? '更新' : '创建' }}
            </el-button>
          </div>
        </template>
      </el-dialog>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, nextTick } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Search, Refresh, ArrowDown, ArrowUp } from '@element-plus/icons-vue'
import { 
  getCategoryTree, 
  createCategory, 
  updateCategory, 
  deleteCategory 
} from '@/api/product'

// 响应式数据
const loading = ref(false)
const submitting = ref(false)
const categoryTree = ref([])
const treeRef = ref()
const dialogVisible = ref(false)
const isEdit = ref(false)
const searchKeyword = ref('')
const statusFilter = ref('')
const selectedNode = ref(null)

// 表单数据
const formData = reactive({
  id: '',
  name: '',
  parentId: '',
  sort: 0,
  icon: '',
  status: 'active'
})

// 表单验证规则
const formRules = {
  name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' },
    { min: 1, max: 50, message: '分类名称长度在1到50个字符', trigger: 'blur' }
  ],
  sort: [
    { required: true, message: '请输入排序值', trigger: 'blur' }
  ]
}

// 计算属性
const dialogTitle = computed(() => {
  return isEdit.value ? '编辑分类' : '新增分类'
})

const filteredCategoryTree = computed(() => {
  if (!searchKeyword.value && !statusFilter.value) {
    return categoryTree.value
  }
  return filterTree(categoryTree.value, searchKeyword.value, statusFilter.value)
})

// 方法
const loadCategories = async () => {
  loading.value = true
  try {
    const response = await getCategoryTree()
    categoryTree.value = response.data.data || []
  } catch (error) {
    ElMessage.error('加载分类失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const filterTree = (tree: any[], keyword: string, status: string): any[] => {
  return tree.filter(node => {
    const matchesKeyword = !keyword || node.name.toLowerCase().includes(keyword.toLowerCase())
    const matchesStatus = !status || node.status === status
    
    const filteredChildren = node.children ? filterTree(node.children, keyword, status) : []
    const hasMatchingChildren = filteredChildren.length > 0
    
    return (matchesKeyword && matchesStatus) || hasMatchingChildren
  }).map(node => ({
    ...node,
    children: node.children ? filterTree(node.children, keyword, status) : []
  }))
}

const handleSearch = () => {
  // 搜索逻辑已在computed中实现
}

const handleRefresh = () => {
  loadCategories()
}

const handleExpandAll = () => {
  nextTick(() => {
    const expandNodes = (nodes: any[]) => {
      nodes.forEach(node => {
        treeRef.value?.setExpanded(node.id, true)
        if (node.children) {
          expandNodes(node.children)
        }
      })
    }
    expandNodes(filteredCategoryTree.value)
  })
}

const handleCollapseAll = () => {
  nextTick(() => {
    const collapseNodes = (nodes: any[]) => {
      nodes.forEach(node => {
        treeRef.value?.setExpanded(node.id, false)
        if (node.children) {
          collapseNodes(node.children)
        }
      })
    }
    collapseNodes(filteredCategoryTree.value)
  })
}

const handleNodeClick = (data: any) => {
  selectedNode.value = data
}

const handleAdd = () => {
  isEdit.value = false
  resetForm()
  dialogVisible.value = true
}

const handleAddChild = (parentData: any) => {
  isEdit.value = false
  resetForm()
  formData.parentId = parentData.id
  dialogVisible.value = true
}

const handleEdit = (data: any) => {
  isEdit.value = true
  selectedNode.value = data
  Object.assign(formData, {
    id: data.id,
    name: data.name,
    parentId: data.parentId || '',
    sort: data.sort || 0,
    icon: data.icon || '',
    status: data.status || 'active'
  })
  dialogVisible.value = true
}

const handleDelete = async (data: any) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除分类"${data.name}"吗？删除后其子分类将变为顶级分类。`,
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
    
    await deleteCategory(data.id)
    ElMessage.success('删除成功')
    loadCategories()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
      console.error(error)
    }
  }
}

const handleToggleStatus = async (data: any) => {
  try {
    const newStatus = data.status === 'active' ? 'inactive' : 'active'
    await updateCategory(data.id, { ...data, status: newStatus })
    ElMessage.success(`${newStatus === 'active' ? '启用' : '禁用'}成功`)
    loadCategories()
  } catch (error) {
    ElMessage.error('状态修改失败')
    console.error(error)
  }
}

const handleSubmit = async () => {
  submitting.value = true
  try {
    const data = { ...formData }
    if (data.parentId === '') {
      delete data.parentId
    }
    
    if (isEdit.value) {
      await updateCategory(formData.id, data)
      ElMessage.success('更新成功')
    } else {
      await createCategory(data)
      ElMessage.success('创建成功')
    }
    
    handleDialogClose()
    loadCategories()
  } catch (error) {
    ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
    console.error(error)
  } finally {
    submitting.value = false
  }
}

const handleDialogClose = () => {
  dialogVisible.value = false
  resetForm()
}

const resetForm = () => {
  Object.assign(formData, {
    id: '',
    name: '',
    parentId: '',
    sort: 0,
    icon: '',
    status: 'active'
  })
}

// 生命周期
onMounted(() => {
  loadCategories()
})
</script>

<style scoped>
.tree-node {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding: 4px 0;
}

.node-content {
  display: flex;
  align-items: center;
  flex: 1;
}

.category-icon {
  margin-right: 8px;
  color: #409eff;
}

.node-label {
  font-weight: 500;
}

.sort-number {
  margin-left: 8px;
  color: #999;
  font-size: 12px;
}

.node-actions {
  display: flex;
  gap: 8px;
  opacity: 0;
  transition: opacity 0.3s;
}

.tree-node:hover .node-actions {
  opacity: 1;
}

.dialog-footer {
  text-align: right;
}
</style>
