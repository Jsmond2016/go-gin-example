import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import DefaultLayout from '../layouts/DefaultLayout.vue'
import Login from '../views/Login.vue'

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth: boolean
  }
}

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { requiresAuth: false }
  },
  {
    path: '/',
    component: DefaultLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        redirect: '/dashboard'
      },
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue')
      },
      {
        path: 'articles',
        name: 'Articles',
        component: () => import('@/views/articles/List.vue')
      },
      {
        path: 'articles/create',
        name: 'CreateArticle',
        component: () => import('@/views/articles/Edit.vue')
      },
      {
        path: 'articles/:id/edit',
        name: 'EditArticle',
        component: () => import('@/views/articles/Edit.vue')
      },
      {
        path: 'tags',
        name: 'Tags',
        component: () => import('@/views/tags/List.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return '/login'
  }

  if (to.path === '/login' && authStore.isAuthenticated) {
    return '/'
  }
})

export default router 