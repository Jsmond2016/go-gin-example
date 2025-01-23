import { defineStore } from 'pinia'
import request from '../api/request'

interface LoginForm {
  username: string
  password: string
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    username: localStorage.getItem('username') || ''
  }),
  
  getters: {
    isAuthenticated: state => !!state.token
  },
  
  actions: {
    async login(form: LoginForm) {
      try {
        const response = await request.post('/auth', form)
        if (response.code === 200) {
          this.token = response.data.token
          this.username = form.username
          localStorage.setItem('token', response.data.token)
          localStorage.setItem('username', form.username)
          return true
        }
        return false
      } catch {
        return false
      }
    },
    
    logout() {
      this.token = ''
      this.username = ''
      localStorage.removeItem('token')
      localStorage.removeItem('username')
    }
  }
}) 