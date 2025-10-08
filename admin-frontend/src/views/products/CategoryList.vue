<template>
  <div>
    <el-card header="分类管理">
      <el-button type="primary" style="margin-bottom: 15px;">新增分类</el-button>
      <el-tree :data="[]" :props="{ label: 'name', children: 'children' }" />
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>
