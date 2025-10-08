<template>
  <div>
    <el-card header="会员列表">
      <el-table :data="[]" border>
        <el-table-column prop="nickname" label="昵称" />
        <el-table-column prop="phone" label="手机号" />
        <el-table-column prop="member_level" label="等级" />
        <el-table-column prop="points" label="积分" />
        <el-table-column prop="status" label="状态" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>
