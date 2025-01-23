import api from './config'
// 接口返回类型
interface ApiResponse<T = any> {
  code: number
  data: T
  msg: string
}

// Auth
export const auth = {
  login:  (username: string, password: string): Promise<ApiResponse> => {
   return api.post('/api/v1/auth', { username, password })
  }
}

// Articles
export const articles = {
  getList: (params?: { tag_id?: number; state?: number; created_by?: number }): Promise<ApiResponse> =>
    api.get('/api/v1/articles', { params }),

  getOne: (id: number): Promise<ApiResponse> =>
    api.get(`/api/v1/articles/${id}`),

  create: (data: {
    tag_id: number
    title: string
    desc: string
    content: string
    created_by: string
    state: number
  }): Promise<ApiResponse> =>
    api.post('/api/v1/articles', data),

  update: (id: number, data: {
    tag_id?: string
    title?: string
    desc?: string
    content?: string
    modified_by: string
    state?: number
  }): Promise<ApiResponse> =>
    api.put(`/api/v1/articles/${id}`, data),

  delete: (id: number): Promise<ApiResponse> =>
    api.delete(`/api/v1/articles/${id}`)
}

// Tags
export const tags = {
  getList: (params?: { name?: string; state?: number }): Promise<ApiResponse> =>
    api.get('/api/v1/tags', { params }),

  create: (data: {
    name: string
    state?: number
    created_by?: number
  }): Promise<ApiResponse> =>
    api.post('/api/v1/tags', data),

  update: (id: number, data: {
    name: string
    state?: number
    modified_by: string
  }): Promise<ApiResponse> =>
    api.put(`/api/v1/tags/${id}`, data),

  delete: (id: number): Promise<ApiResponse> =>
    api.delete(`/api/v1/tags/${id}`),

  export: (data?: { name?: string; state?: number }): Promise<ApiResponse> =>
    api.post('/api/v1/tags/export', data),

  import: (file: File): Promise<ApiResponse> => {
    const formData = new FormData()
    formData.append('image', file)
    return api.post('/api/v1/tags/import', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  }
} 