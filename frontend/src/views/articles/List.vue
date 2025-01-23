<template>
  <div class="bg-white shadow overflow-hidden sm:rounded-lg">
    <div class="px-4 py-5 sm:px-6 flex justify-between items-center">
      <h3 class="text-lg leading-6 font-medium text-gray-900">Articles</h3>
      <router-link
        to="/articles/create"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"
      >
        Create Article
      </router-link>
    </div>
    <div class="border-t border-gray-200">
      <div v-if="loading" class="text-center py-4">Loading...</div>
      <div v-else-if="error" class="text-center py-4 text-red-600">{{ error }}</div>
      <div v-else>
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Description</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">State</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Created By</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="article in articles" :key="article.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ article.title }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ article.desc }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span
                  :class="[
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                    article.state === 1 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                  ]"
                >
                  {{ article.state === 1 ? 'Active' : 'Inactive' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ article.created_by }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <router-link
                  :to="`/articles/${article.id}`"
                  class="text-indigo-600 hover:text-indigo-900 mr-4"
                >
                  Edit
                </router-link>
                <button
                  @click="deleteArticle(article.id)"
                  class="text-red-600 hover:text-red-900"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { articles as articleApi } from '../../api/services'

interface Article {
  id: number
  title: string
  desc: string
  state: number
  created_by: string
}

const loading = ref(false)
const error = ref('')
const articles = ref<Article[]>([])

const fetchArticles = async () => {
  loading.value = true
  error.value = ''
  try {
    const response = await articleApi.getList()
    if (response.code === 200) {
      articles.value = response.data
    } else {
      error.value = response.msg || 'Failed to fetch articles'
    }
  } catch (err) {
    error.value = 'An error occurred while fetching articles'
    console.error('Error fetching articles:', err)
  } finally {
    loading.value = false
  }
}

const deleteArticle = async (id: number) => {
  if (!confirm('Are you sure you want to delete this article?')) {
    return
  }

  try {
    const response = await articleApi.delete(id)
    if (response.code === 200) {
      await fetchArticles()
    } else {
      alert(response.msg || 'Failed to delete article')
    }
  } catch (err) {
    console.error('Error deleting article:', err)
    alert('An error occurred while deleting the article')
  }
}

onMounted(fetchArticles)
</script> 