<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between;">
          <span>商品列表</span>
          <el-button type="primary">新增商品</el-button>
        </div>
      </template>
      
      <el-table :data="[]" border>
        <el-table-column prop="name" label="商品名称" />
        <el-table-column prop="code" label="商品编码" />
        <el-table-column prop="category" label="分类" />
        <el-table-column prop="price" label="价格" />
        <el-table-column prop="status" label="状态" />
        <el-table-column label="操作" width="200">
          <template #default>
            <el-button size="small">编辑</el-button>
            <el-button size="small" type="danger">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>
