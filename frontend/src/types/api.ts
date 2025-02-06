export interface ListParams {
  page?: number
  page_size?: number
  tag_id?: number
  state?: number
  name?: string
}

export interface ApiResponse<T = any> {
  code: number
  msg?: string
  data: T
}

export interface ListResponse<T> {
  list: T[]
  total: number
} 