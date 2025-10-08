<template>
  <div class="dashboard">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card>
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #409EFF;">
              <el-icon :size="32"><Shop /></el-icon>
            </div>
            <div class="stat-info">
              <p class="stat-label">门店总数</p>
              <p class="stat-value">10</p>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card>
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #67C23A;">
              <el-icon :size="32"><Goods /></el-icon>
            </div>
            <div class="stat-info">
              <p class="stat-label">商品总数</p>
              <p class="stat-value">1,234</p>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card>
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #E6A23C;">
              <el-icon :size="32"><Document /></el-icon>
            </div>
            <div class="stat-info">
              <p class="stat-label">今日订单</p>
              <p class="stat-value">89</p>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card>
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #F56C6C;">
              <el-icon :size="32"><Coin /></el-icon>
            </div>
            <div class="stat-info">
              <p class="stat-label">今日营业额</p>
              <p class="stat-value">¥12,456</p>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    
    <el-row :gutter="20" style="margin-top: 20px;">
      <el-col :span="16">
        <el-card header="销售趋势">
          <div style="height: 300px; display: flex; align-items: center; justify-content: center; color: #999;">
            销售趋势图表（使用ECharts实现）
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="8">
        <el-card header="门店状态">
          <div class="store-status">
            <div class="status-item">
              <span class="status-dot" style="background-color: #67C23A;"></span>
              <span>正常运营：8家</span>
            </div>
            <div class="status-item">
              <span class="status-dot" style="background-color: #E6A23C;"></span>
              <span>维护中：1家</span>
            </div>
            <div class="status-item">
              <span class="status-dot" style="background-color: #F56C6C;"></span>
              <span>关闭：1家</span>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    
    <el-row :gutter="20" style="margin-top: 20px;">
      <el-col :span="12">
        <el-card header="库存预警">
          <el-table :data="[]" height="200">
            <el-table-column prop="product" label="商品" />
            <el-table-column prop="store" label="门店" />
            <el-table-column prop="quantity" label="当前库存" />
            <el-table-column prop="alert" label="预警数量" />
          </el-table>
        </el-card>
      </el-col>
      
      <el-col :span="12">
        <el-card header="最近订单">
          <el-table :data="[]" height="200">
            <el-table-column prop="order_no" label="订单号" />
            <el-table-column prop="amount" label="金额" />
            <el-table-column prop="status" label="状态" />
            <el-table-column prop="time" label="时间" />
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.dashboard {
  padding: 0;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 15px;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}

.stat-info {
  flex: 1;
}

.stat-label {
  margin: 0;
  color: #909399;
  font-size: 14px;
}

.stat-value {
  margin: 5px 0 0 0;
  font-size: 24px;
  font-weight: bold;
  color: #303133;
}

.store-status {
  padding: 10px 0;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 0;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}
</style>
