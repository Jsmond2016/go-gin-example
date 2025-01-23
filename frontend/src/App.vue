<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'

const router = useRouter()
const authStore = useAuthStore()

// 检查认证状态
if (!authStore.isAuthenticated && router.currentRoute.value.path !== '/login') {
  router.push('/login')
}
</script>

<template>
  <el-config-provider :locale="zhCn">
    <router-view v-slot="{ Component }">
      <component :is="Component" />
    </router-view>
  </el-config-provider>
</template>

<style>
@import './assets/css/main.css';
</style>
