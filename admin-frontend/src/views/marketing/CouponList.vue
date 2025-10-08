<template>
  <div>
    <el-card header="优惠券管理">
      <el-button type="primary" style="margin-bottom: 15px;">新增优惠券</el-button>
      <el-table :data="[]" border>
        <el-table-column prop="name" label="优惠券名称" />
        <el-table-column prop="type" label="类型" />
        <el-table-column prop="value" label="面值" />
        <el-table-column prop="total" label="总量" />
        <el-table-column prop="used" label="已使用" />
        <el-table-column prop="status" label="状态" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>
