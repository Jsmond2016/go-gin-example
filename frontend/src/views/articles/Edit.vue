<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span class="text-xl font-bold">{{
          isEdit ? "编辑文章" : "创建文章"
        }}</span>
      </div>
    </template>

    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="120px"
      class="mt-4"
      @submit.prevent="handleSubmit"
    >
      <el-form-item label="标题" prop="title">
        <el-input v-model="form.title" placeholder="请输入文章标题" />
      </el-form-item>

      <el-form-item label="描述" prop="desc">
        <el-input v-model="form.desc" placeholder="请输入文章描述" />
      </el-form-item>

      <el-form-item label="内容" prop="content">
        <el-input
          v-model="form.content"
          type="textarea"
          :rows="6"
          placeholder="请输入文章内容"
        />
      </el-form-item>

      <el-form-item label="标签" prop="tag_id">
        <el-select
          style="display: block; width: 100%"
          v-model="form.tag_id"
          placeholder="请选择标签"
        >
          <el-option
            v-for="tag in tags"
            :key="tag.id"
            :label="tag.name"
            :value="tag.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="状态" prop="state">
        <el-select style="display: block; width: 100%;" v-model="form.state" placeholder="请选择状态">
          <el-option :value="1" label="启用" />
          <el-option :value="0" label="禁用" />
        </el-select>
      </el-form-item>

      <el-form-item label="封面图" prop="cover_image_url">
        <el-input
          v-model="form.cover_image_url"
          placeholder="请输入封面图URL"
        />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" :loading="loading" @click="handleSubmit">
          {{ loading ? "保存中..." : "保存" }}
        </el-button>
        <el-button @click="$router.push('/articles')">取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script setup lang="ts">
import { ref, computed } from "vue"
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from "element-plus"
import { articles as articleApi, tags as tagApi } from "../../api/services"
import { useAuthStore } from "../../stores/auth"

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const formRef = ref()

interface Tag {
  id: number
  name: string
}

const loading = ref(false)
const tags = ref<Tag[]>([])

const isEdit = computed(() => !!route.params.id)

const form = ref({
  title: "",
  desc: "",
  content: "",
  tag_id: "",
  state: 1,
  created_by: authStore.username || "",
  modified_by: authStore.username || "",
  cover_image_url: "",
})

const rules = {
  title: [{ required: true, message: "请输入标题", trigger: "blur" }],
  desc: [{ required: true, message: "请输入描述", trigger: "blur" }],
  content: [{ required: true, message: "请输入内容", trigger: "blur" }],
  tag_id: [{ required: true, message: "请选择标签", trigger: "change" }],
  state: [{ required: true, message: "请选择状态", trigger: "change" }],
}

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

const fetchArticle = async () => {
  if (!isEdit.value) return

  try {
    const response = await articleApi.getOne(Number(route.params.id))
    if (response.code === 200) {
      const article = response.data
      form.value = {
        ...form.value,
        ...article,
        tag_id: article.tag_id,
      }
    }
  } catch (err) {
    console.error("Error fetching article:", err)
    ElMessage.error("获取文章详情失败")
    router.push("/articles")
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid: boolean) => {
    if (!valid) return

    loading.value = true
    try {
      if (isEdit.value) {
        const updateData = {
          tag_id: form.value.tag_id,
          title: form.value.title,
          desc: form.value.desc,
          content: form.value.content,
          state: form.value.state,
          modified_by: form.value.modified_by,
          cover_image_url: form.value.cover_image_url,
        }
        const response = await articleApi.update(
          Number(route.params.id),
          updateData
        )
        if (response.code === 200) {
          ElMessage.success("更新成功")
          router.push("/articles")
        } else {
          ElMessage.error(response.msg || "更新失败")
        }
      } else {
        const createData = {
          tag_id: Number(form.value.tag_id),
          title: form.value.title,
          desc: form.value.desc,
          content: form.value.content,
          state: form.value.state,
          created_by: form.value.created_by,
          cover_image_url: form.value.cover_image_url,
        }
        const response = await articleApi.create(createData)
        if (response.code === 200) {
          ElMessage.success("创建成功")
          router.push("/articles")
        } else {
          ElMessage.error(response.msg || "创建失败")
        }
      }
    } catch (err) {
      console.error("Error saving article:", err)
      ElMessage.error("保存失败")
    } finally {
      loading.value = false
    }
  })
}

// 组件挂载时获取数据
fetchTags()
if (route.params.id) {
  fetchArticle()
}
</script>
