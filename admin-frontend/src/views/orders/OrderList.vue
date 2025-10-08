<template>
  <div>
    <el-card header="订单列表">
      <el-table :data="[]" border>
        <el-table-column prop="order_no" label="订单号" />
        <el-table-column prop="store" label="门店" />
        <el-table-column prop="amount" label="金额" />
        <el-table-column prop="status" label="状态" />
        <el-table-column prop="created_at" label="下单时间" />
        <el-table-column label="操作">
          <template #default>
            <el-button size="small">详情</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>
