import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: () => import('../views/Layout.vue'),
      children: [
        {
          path: '',
          name: 'home',
          component: () => import('../views/Home.vue')
        },
        {
          path: 'articles',
          name: 'articles',
          component: () => import('../views/articles/List.vue')
        },
        {
          path: 'articles/create',
          name: 'articles-create',
          component: () => import('../views/articles/Edit.vue')
        },
        {
          path: 'articles/:id',
          name: 'articles-edit',
          component: () => import('../views/articles/Edit.vue')
        },
        {
          path: 'tags',
          name: 'tags',
          component: () => import('../views/tags/List.vue')
        }
      ],
      meta: { requiresAuth: true }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue')
    }
  ]
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!authStore.isAuthenticated) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router 