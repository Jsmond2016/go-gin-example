<template>
  <el-container class="layout-container">
    <!-- 侧边栏 -->
    <el-aside width="220px" class="aside">
      <div class="logo">
        <h1 class="text-xl font-bold text-white">博客管理系统</h1>
      </div>
      
      <el-menu
        :default-active="route.path"
        class="el-menu-vertical"
        background-color="#304156"
        text-color="#fff"
        active-text-color="#409EFF"
        :router="true"
      >
        <el-menu-item index="/dashboard">
          <el-icon><DataLine /></el-icon>
          <span>仪表盘</span>
        </el-menu-item>
        
        <el-menu-item index="/articles">
          <el-icon><Document /></el-icon>
          <span>文章管理</span>
        </el-menu-item>
        
        <el-menu-item index="/tags">
          <el-icon><Collection /></el-icon>
          <span>标签管理</span>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <!-- 主要内容区 -->
    <el-container>
      <!-- 顶部导航 -->
      <el-header class="header">
        <div class="flex justify-between items-center h-full">
          <el-breadcrumb>
            <el-breadcrumb-item v-for="item in breadcrumbs" :key="item.path" :to="item.path">
              {{ item.name }}
            </el-breadcrumb-item>
          </el-breadcrumb>
          
          <el-dropdown @command="handleCommand">
            <span class="el-dropdown-link">
              {{ authStore.username }}
              <el-icon class="el-icon--right"><arrow-down /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 内容区 -->
      <el-main class="main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import { DataLine, Document, Collection, ArrowDown } from '@element-plus/icons-vue'
import { useAuthStore } from '../stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

// 面包屑导航
const breadcrumbs = computed(() => {
  const paths = route.path.split('/').filter(Boolean)
  const items = [{ path: '/', name: '首页' }]
  
  let currentPath = ''
  paths.forEach(path => {
    currentPath += `/${path}`
    items.push({
      path: currentPath,
      name: getBreadcrumbName(path)
    })
  })
  
  return items
})

// 获取面包屑名称
function getBreadcrumbName(path: string): string {
  const nameMap: Record<string, string> = {
    dashboard: '仪表盘',
    articles: '文章管理',
    tags: '标签管理',
    create: '新建',
    edit: '编辑'
  }
  return nameMap[path] || path
}

// 处理下拉菜单命令
const handleCommand = async (command: string) => {
  if (command === 'logout') {
    try {
      await ElMessageBox.confirm(
        '确定要退出登录吗？',
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
      await authStore.logout()
      router.push('/login')
    } catch {
      // 用户取消操作
    }
  } else if (command === 'profile') {
    // 处理个人信息
    router.push('/profile')
  }
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
}

.aside {
  background-color: #304156;
  color: #fff;
}

.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #2b3648;
}

.header {
  background-color: #fff;
  border-bottom: 1px solid #dcdfe6;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
  padding: 0 20px;
}

.main {
  background-color: #f0f2f5;
  padding: 20px;
}

.el-menu {
  border-right: none;
}

.el-dropdown-link {
  cursor: pointer;
  display: flex;
  align-items: center;
  color: #606266;
}

:deep(.el-menu-item) {
  display: flex;
  align-items: center;
}

:deep(.el-menu-item .el-icon) {
  margin-right: 10px;
}
</style> 