<template>
  <div>
    <el-card header="库存管理">
      <el-table :data="[]" border>
        <el-table-column prop="product" label="商品" />
        <el-table-column prop="store" label="门店" />
        <el-table-column prop="quantity" label="库存数量" />
        <el-table-column prop="alert_qty" label="预警数量" />
        <el-table-column label="操作">
          <template #default>
            <el-button size="small">入库</el-button>
            <el-button size="small">出库</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>
