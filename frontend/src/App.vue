<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'

const router = useRouter()
const authStore = useAuthStore()

// 检查认证状态
if (!authStore.isAuthenticated && router.currentRoute.value.path !== '/login') {
  router.push('/login')
}
</script>

<template>
  <router-view></router-view>
</template>

<style>
@import './assets/css/main.css';
</style>
