<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span class="text-xl font-bold">仪表盘</span>
      </div>
    </template>
    
    <div class="dashboard-content">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card shadow="hover" class="stat-card">
            <template #header>
              <div class="stat-header">
                <el-icon><Document /></el-icon>
                <span>文章总数</span>
              </div>
            </template>
            <div class="stat-value">{{ stats.articleCount }}</div>
          </el-card>
        </el-col>
        
        <el-col :span="6">
          <el-card shadow="hover" class="stat-card">
            <template #header>
              <div class="stat-header">
                <el-icon><Collection /></el-icon>
                <span>标签总数</span>
              </div>
            </template>
            <div class="stat-value">{{ stats.tagCount }}</div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Document, Collection } from '@element-plus/icons-vue'
import { articles as articleApi, tags as tagApi } from '../api/services'

interface Stats {
  articleCount: number
  tagCount: number
}

const stats = ref<Stats>({
  articleCount: 0,
  tagCount: 0
})

const fetchStats = async () => {
  try {
    const [articlesRes, tagsRes] = await Promise.all([
      articleApi.getList({}),
      tagApi.getList({})
    ])
    
    if (articlesRes.code === 200 && tagsRes.code === 200) {
      stats.value = {
        articleCount: articlesRes.data.total || 0,
        tagCount: tagsRes.data.total || 0
      }
    }
  } catch (err) {
    console.error('Error fetching stats:', err)
  }
}

onMounted(() => {
  fetchStats()
})
</script>

<style scoped>
.dashboard-content {
  padding: 20px 0;
}

.stat-card {
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-5px);
}

.stat-header {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #606266;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  text-align: center;
  padding: 10px 0;
}
</style> 