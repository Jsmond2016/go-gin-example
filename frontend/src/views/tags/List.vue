<template>
  <div class="tag-list">
    <el-card class="box-card">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-xl font-bold">标签管理</span>
          <el-button type="primary" @click="handleAdd">新建标签</el-button>
        </div>
      </template>

      <!-- 搜索区域 -->
      <el-form :inline="true" :model="searchForm" class="mb-4">
        <el-form-item label="名称">
          <el-input
            v-model="searchForm.name"
            placeholder="标签名称"
            clearable
          />
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
      <el-table v-loading="loading" :data="tags" border style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="名称" min-width="150" />
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
              <el-button type="primary" :icon="Edit" @click="handleEdit(row)">
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

    <!-- 标签编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="editingTag.id ? '编辑标签' : '新建标签'"
      width="500px"
    >
      <el-form
        ref="formRef"
        :model="editingTag"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item label="名称" prop="name">
          <el-input v-model="editingTag.name" placeholder="请输入标签名称" />
        </el-form-item>
        <el-form-item label="状态" prop="state">
          <el-select v-model="editingTag.state" placeholder="请选择状态">
            <el-option :value="1" label="启用" />
            <el-option :value="0" label="禁用" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button
            type="primary"
            :loading="submitLoading"
            @click="handleSubmit"
          >
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { Delete, Edit } from "@element-plus/icons-vue"
import { tags as tagApi } from "../../api/services"
import { useAuthStore } from "../../stores/auth"
import dayjs from 'dayjs'

const authStore = useAuthStore()

interface Tag {
  id: number
  name: string
  state: number
  created_by: string
  created_on: number
}

const loading = ref(false)
const submitLoading = ref(false)
const tags = ref<Tag[]>([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const dialogVisible = ref(false)
const formRef = ref()

const searchForm = ref({
  name: "",
  state: undefined,
})

const editingTag = ref({
  id: 0,
  name: "",
  state: 1,
  created_by: authStore.username || "",
  modified_by: authStore.username || "",
})

const rules = {
  name: [{ required: true, message: "请输入标签名称", trigger: "blur" }],
  state: [{ required: true, message: "请选择状态", trigger: "change" }],
}

// 获取标签列表
const fetchTags = async () => {
  loading.value = true
  try {
    const response = await tagApi.getList({
      page: currentPage.value,
      page_size: pageSize.value,
      ...searchForm.value,
    })
    if (response.code === 200) {
      tags.value = response.data.list
      total.value = response.data.total
    }
  } catch (err) {
    console.error("Error fetching tags:", err)
    ElMessage.error("获取标签列表失败")
  } finally {
    loading.value = false
  }
}

// 新建标签
const handleAdd = () => {
  editingTag.value = {
    id: 0,
    name: "",
    state: 1,
    created_by: authStore.username || "",
    modified_by: authStore.username || "",
  }
  dialogVisible.value = true
}

// 编辑标签
const handleEdit = (row: Tag) => {
  editingTag.value = {
    ...row,
    modified_by: authStore.username || "",
  }
  dialogVisible.value = true
}

// 删除标签
const handleDelete = (row: Tag) => {
  ElMessageBox.confirm(`确定要删除标签 "${row.name}" 吗？`, "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      try {
        const response = await tagApi.delete(row.id)
        if (response.code === 200) {
          ElMessage.success("删除成功")
          fetchTags()
        } else {
          ElMessage.error(response.msg || "删除失败")
        }
      } catch (err) {
        console.error("Error deleting tag:", err)
        ElMessage.error("删除失败")
      }
    })
    .catch(() => {})
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid: boolean) => {
    if (!valid) return

    submitLoading.value = true
    try {
      const api = editingTag.value.id ? tagApi.update : tagApi.create
      const response = await api({
        name: editingTag.value.name,
        state: editingTag.value.state,
        created_by: authStore.username || "",
        // modified_by: authStore.username || "",
      })
      
      if (response.code === 200) {
        ElMessage.success(editingTag.value.id ? "更新成功" : "创建成功")
        dialogVisible.value = false
        fetchTags()
      } else {
        ElMessage.error(
          response.msg || (editingTag.value.id ? "更新失败" : "创建失败")
        )
      }
    } catch (err) {
      console.error("Error saving tag:", err)
      ElMessage.error("保存失败")
    } finally {
      submitLoading.value = false
    }
  })
}

// 搜索
const handleSearch = () => {
  currentPage.value = 1
  fetchTags()
}

// 重置搜索
const resetSearch = () => {
  searchForm.value = {
    name: "",
    state: undefined,
  }
  currentPage.value = 1
  fetchTags()
}

// 分页
const handleSizeChange = (val: number) => {
  pageSize.value = val
  fetchTags()
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
  fetchTags()
}

onMounted(() => {
  fetchTags()
})
</script>

<style scoped>
.tag-list {
  padding: 20px;
}
</style>
