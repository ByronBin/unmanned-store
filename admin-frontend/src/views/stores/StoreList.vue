<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between;">
          <span>门店列表</span>
          <el-button type="primary" @click="handleAdd">新增门店</el-button>
        </div>
      </template>
      
      <el-table :data="stores" border>
        <el-table-column prop="name" label="门店名称" />
        <el-table-column prop="code" label="门店编码" />
        <el-table-column prop="address" label="地址" />
        <el-table-column prop="phone" label="电话" />
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag :type="row.status === 'active' ? 'success' : 'info'">
              {{ row.status === 'active' ? '营业中' : '已关闭' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getStoreList } from '@/api/store'

const stores = ref([])

const fetchStores = async () => {
  try {
    const res: any = await getStoreList()
    stores.value = res.data || []
  } catch (error) {
    console.error(error)
  }
}

const handleAdd = () => {
  // TODO: 打开新增对话框
}

const handleEdit = (row: any) => {
  // TODO: 打开编辑对话框
}

const handleDelete = (row: any) => {
  // TODO: 删除门店
}

onMounted(() => {
  fetchStores()
})
</script>
