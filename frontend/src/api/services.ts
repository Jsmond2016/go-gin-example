import request from './request'
import type { ListParams, ApiResponse, ListResponse } from '../types/api'

// 接口返回类型
interface ApiResponse<T = any> {
  code: number
  data: T
  msg: string
}

// Auth
export const auth = {
  login:  (username: string, password: string): Promise<ApiResponse> => {
   return request.post('/auth', { username, password })
  }
}

export interface Tag {
  id: number
  name: string
  state: number
  created_by: string
  created_on: number
}

export interface Article {
  id: number
  title: string
  desc: string
  content: string
  cover_image_url: string
  state: number
  tag_id: number
  tag: Tag
  created_by: string
  created_on: number
}

// Articles
export const articles = {
  getList: (params: ListParams) => 
    request.get<ApiResponse<ListResponse<Article>>>('/articles', { params }),
  
  getOne: (id: number) => 
    request.get<ApiResponse<Article>>(`/articles/${id}`),
  
  create: (data: Partial<Article>) => 
    request.post<ApiResponse>('/articles', data),
  
  update: (id: number, data: Partial<Article>) => 
    request.put<ApiResponse>(`/articles/${id}`, data),
  
  delete: (id: number) => 
    request.delete<ApiResponse>(`/articles/${id}`)
}

// Tags
export const tags = {
  getList: (params: ListParams) => 
    request.get<ApiResponse<ListResponse<Tag>>>('/tags', { params }),
  
  create: (data: Partial<Tag>) => 
    request.post<ApiResponse>('/tags', data),
  
  update: (id: number, data: Partial<Tag>) => 
    request.put<ApiResponse>(`/tags/${id}`, data),
  
  delete: (id: number) => 
    request.delete<ApiResponse>(`/tags/${id}`),

  export: (data?: { name?: string; state?: number }): Promise<ApiResponse> =>
    request.post('/tags/export', data),

  import: (file: File): Promise<ApiResponse> => {
    const formData = new FormData()
    formData.append('image', file)
    return request.post('/tags/import', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  }
} 