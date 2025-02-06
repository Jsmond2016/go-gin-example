<template>
  <div class="article-list">
    <el-card class="box-card">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-xl font-bold">文章列表</span>
          <el-button type="primary" @click="$router.push('/articles/create')">
            新建文章
          </el-button>
        </div>
      </template>

      <!-- 搜索区域 -->
      <el-form :inline="true" :model="searchForm" class="mb-4">
        <el-form-item label="标签">
          <el-select
            v-model="searchForm.tag_id"
            placeholder="选择标签"
            clearable
          >
            <el-option
              v-for="tag in tags"
              :key="tag.id"
              :label="tag.name"
              :value="tag.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select
            v-model="searchForm.state"
            placeholder="选择状态"
            clearable
          >
            <el-option :value="1" label="启用" />
            <el-option :value="0" label="禁用" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <!-- 表格区域 -->
      <el-table v-loading="loading" :data="articles" border style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="title" label="标题" min-width="200">
          <template #default="{ row }">
            <el-link
              type="primary"
              @click="$router.push(`/articles/${row.id}`)"
            >
              {{ row.title }}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column
          prop="desc"
          label="描述"
          min-width="250"
          show-overflow-tooltip
        />
        <el-table-column prop="tag.name" label="标签" width="120" />
        <el-table-column prop="state" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.state === 1 ? 'success' : 'info'">
              {{ row.state === 1 ? "启用" : "禁用" }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_by" label="创建者" width="120" />
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="{ row }">
            {{ dayjs(row.created_at).format("YYYY-MM-DD HH:mm:ss") }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button-group>
              <el-button
                type="primary"
                :icon="Edit"
                @click="$router.push(`/articles/${row.id}/edit`)"
              >
                编辑
              </el-button>
              <el-button
                type="danger"
                :icon="Delete"
                @click="handleDelete(row)"
              >
                删除
              </el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="flex justify-end mt-4">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          :page-sizes="[10, 20, 30, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRouter } from "vue-router"
import { ElMessage, ElMessageBox } from "element-plus"
import { Delete, Edit } from "@element-plus/icons-vue"
import { articles as articleApi, tags as tagApi } from "../../api/services"
import dayjs from "dayjs"

interface Tag {
  id: number
  name: string
}

interface Article {
  id: number
  title: string
  desc: string
  content: string
  tag: Tag
  state: number
  created_by: string
  created_on: number
}

const router = useRouter()
const loading = ref(false)
const articles = ref<Article[]>([])
const tags = ref<Tag[]>([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)

const searchForm = ref({
  tag_id: undefined,
  state: undefined,
})

// 获取标签列表
const fetchTags = async () => {
  try {
    const response = await tagApi.getList()
    if (response.code === 200) {
      tags.value = response.data.list
    }
  } catch (err) {
    console.error("Error fetching tags:", err)
    ElMessage.error("获取标签列表失败")
  }
}

// 获取文章列表
const fetchArticles = async () => {
  loading.value = true
  try {
    const response = await articleApi.getList({
      page: currentPage.value,
      page_size: pageSize.value,
      ...searchForm.value,
    })
    if (response.code === 200) {
      articles.value = response.data.list
      total.value = response.data.total
    }
  } catch (err) {
    console.error("Error fetching articles:", err)
    ElMessage.error("获取文章列表失败")
  } finally {
    loading.value = false
  }
}

// 处理删除
const handleDelete = (row: Article) => {
  ElMessageBox.confirm(`确定要删除文章 "${row.title}" 吗？`, "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      try {
        const response = await articleApi.delete(row.id)
        if (response.code === 200) {
          ElMessage.success("删除成功")
          fetchArticles()
        } else {
          ElMessage.error(response.msg || "删除失败")
        }
      } catch (err) {
        console.error("Error deleting article:", err)
        ElMessage.error("删除失败")
      }
    })
    .catch(() => {})
}

// 搜索
const handleSearch = () => {
  currentPage.value = 1
  fetchArticles()
}

// 重置搜索
const resetSearch = () => {
  searchForm.value = {
    tag_id: undefined,
    state: undefined,
  }
  currentPage.value = 1
  fetchArticles()
}

// 分页
const handleSizeChange = (val: number) => {
  pageSize.value = val
  fetchArticles()
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
  fetchArticles()
}

onMounted(() => {
  fetchTags()
  fetchArticles()
})
</script>

<style scoped>
.article-list {
  padding: 20px;
}
</style>
