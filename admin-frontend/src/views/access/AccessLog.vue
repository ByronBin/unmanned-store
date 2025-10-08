<template>
  <div>
    <el-card header="门禁记录">
      <el-table :data="[]" border>
        <el-table-column prop="store" label="门店" />
        <el-table-column prop="user" label="用户" />
        <el-table-column prop="action" label="动作" />
        <el-table-column prop="method" label="方式" />
        <el-table-column prop="status" label="状态" />
        <el-table-column prop="created_at" label="时间" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>
