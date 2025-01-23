<template>
  <div class="bg-white shadow sm:rounded-lg">
    <div class="px-4 py-5 sm:p-6">
      <h3 class="text-lg leading-6 font-medium text-gray-900">
        {{ isEdit ? 'Edit Article' : 'Create Article' }}
      </h3>
      <div class="mt-5">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div>
            <label for="title" class="block text-sm font-medium text-gray-700">Title</label>
            <input
              type="text"
              id="title"
              v-model="form.title"
              required
              class="mt-1 block w-full shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm border-gray-300 rounded-md"
            />
          </div>

          <div>
            <label for="desc" class="block text-sm font-medium text-gray-700">Description</label>
            <input
              type="text"
              id="desc"
              v-model="form.desc"
              required
              class="mt-1 block w-full shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm border-gray-300 rounded-md"
            />
          </div>

          <div>
            <label for="content" class="block text-sm font-medium text-gray-700">Content</label>
            <textarea
              id="content"
              v-model="form.content"
              rows="4"
              required
              class="mt-1 block w-full shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm border-gray-300 rounded-md"
            ></textarea>
          </div>

          <div>
            <label for="tag_id" class="block text-sm font-medium text-gray-700">Tag</label>
            <select
              id="tag_id"
              v-model="form.tag_id"
              required
              class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
            >
              <option v-for="tag in tags" :key="tag.id" :value="tag.id">
                {{ tag.name }}
              </option>
            </select>
          </div>

          <div>
            <label for="state" class="block text-sm font-medium text-gray-700">State</label>
            <select
              id="state"
              v-model="form.state"
              required
              class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
            >
              <option :value="1">Active</option>
              <option :value="0">Inactive</option>
            </select>
          </div>

          <div class="flex justify-end space-x-3">
            <button
              type="button"
              @click="$router.push('/articles')"
              class="inline-flex justify-center py-2 px-4 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              {{ loading ? 'Saving...' : 'Save' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { articles as articleApi, tags as tagApi } from '../../api/services'
import { useAuthStore } from '../../stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

interface Tag {
  id: number
  name: string
}

const loading = ref(false)
const tags = ref<Tag[]>([])

const isEdit = computed(() => !!route.params.id)

const form = ref({
  title: '',
  desc: '',
  content: '',
  tag_id: '',
  state: 1,
  created_by: authStore.username || '',
  modified_by: authStore.username || ''
})

const fetchTags = async () => {
  try {
    const response = await tagApi.getList()
    if (response.code === 200) {
      tags.value = response.data
    }
  } catch (err) {
    console.error('Error fetching tags:', err)
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
        tag_id: article.tag_id.toString()
      }
    }
  } catch (err) {
    console.error('Error fetching article:', err)
    alert('Failed to fetch article')
    router.push('/articles')
  }
}

const handleSubmit = async () => {
  loading.value = true
  try {
    if (isEdit.value) {
      const updateData = {
        tag_id: form.value.tag_id,
        title: form.value.title,
        desc: form.value.desc,
        content: form.value.content,
        state: form.value.state,
        modified_by: form.value.modified_by
      }
      const response = await articleApi.update(Number(route.params.id), updateData)
      if (response.code === 200) {
        router.push('/articles')
      } else {
        alert(response.msg || 'Failed to update article')
      }
    } else {
      const createData = {
        tag_id: Number(form.value.tag_id),
        title: form.value.title,
        desc: form.value.desc,
        content: form.value.content,
        state: form.value.state,
        created_by: form.value.created_by
      }
      const response = await articleApi.create(createData)
      if (response.code === 200) {
        router.push('/articles')
      } else {
        alert(response.msg || 'Failed to create article')
      }
    }
  } catch (err) {
    console.error('Error saving article:', err)
    alert('An error occurred while saving the article')
  } finally {
    loading.value = false
  }
}

// 组件挂载时获取数据
fetchTags()
if (route.params.id) {
  fetchArticle()
}
</script> 