import { defineStore } from 'pinia'
import { auth } from '../api/services'

interface AuthState {
  token: string | null
  username: string | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    token: localStorage.getItem('token'),
    username: localStorage.getItem('username')
  }),

  actions: {
    async login(username: string, password: string) {
      try {
        const response = await auth.login(username, password)
        if (response.code === 200 && response.data.token) {
          this.token = response.data.token
          this.username = username
          localStorage.setItem('token', response.data.token)
          localStorage.setItem('username', username)
          return true
        }
        return false
      } catch (error) {
        console.error('Login failed:', error)
        return false
      }
    },

    logout() {
      this.token = null
      this.username = null
      localStorage.removeItem('token')
      localStorage.removeItem('username')
    }
  },

  getters: {
    isAuthenticated: (state) => !!state.token
  }
}) 