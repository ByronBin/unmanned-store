<template>
  <div>
    <el-card header="监控中心">
      <el-row :gutter="20">
        <el-col :span="6" v-for="i in 4" :key="i">
          <div class="video-panel">
            <div class="video-placeholder">
              <el-icon :size="48"><VideoCamera /></el-icon>
              <p>摄像头 {{ i }}</p>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.video-panel {
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  overflow: hidden;
}

.video-placeholder {
  background-color: #f5f7fa;
  height: 180px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #909399;
}
</style>
