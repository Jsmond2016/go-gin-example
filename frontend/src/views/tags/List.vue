<template>
  <div class="bg-white shadow overflow-hidden sm:rounded-lg">
    <div class="px-4 py-5 sm:px-6 flex justify-between items-center">
      <h3 class="text-lg leading-6 font-medium text-gray-900">Tags</h3>
      <div class="flex space-x-3">
        <button
          @click="showCreateModal = true"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"
        >
          Create Tag
        </button>
        <button
          @click="handleExport"
          class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
        >
          Export
        </button>
        <label
          class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 cursor-pointer"
        >
          Import
          <input
            type="file"
            @change="handleImport"
            class="hidden"
            accept=".xlsx,.xls,.csv"
          />
        </label>
      </div>
    </div>

    <div class="border-t border-gray-200">
      <div v-if="loading" class="text-center py-4">Loading...</div>
      <div v-else-if="error" class="text-center py-4 text-red-600">{{ error }}</div>
      <div v-else>
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">State</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="tag in tags" :key="tag.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ tag.name }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span
                  :class="[
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                    tag.state === 1 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                  ]"
                >
                  {{ tag.state === 1 ? 'Active' : 'Inactive' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button
                  @click="editTag(tag)"
                  class="text-indigo-600 hover:text-indigo-900 mr-4"
                >
                  Edit
                </button>
                <button
                  @click="deleteTag(tag.id)"
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

    <!-- Create/Edit Modal -->
    <div v-if="showCreateModal || editingTag" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
      <div class="bg-white rounded-lg p-6 max-w-md w-full">
        <h3 class="text-lg font-medium text-gray-900 mb-4">
          {{ editingTag ? 'Edit Tag' : 'Create Tag' }}
        </h3>
        <form @submit.prevent="handleSubmit" class="space-y-4">
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700">Name</label>
            <input
              type="text"
              id="name"
              v-model="form.name"
              required
              class="mt-1 block w-full shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm border-gray-300 rounded-md"
            />
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

          <div class="flex justify-end space-x-3 mt-6">
            <button
              type="button"
              @click="closeModal"
              class="inline-flex justify-center py-2 px-4 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"
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
import { ref } from 'vue'
import { useAuthStore } from '../../stores/auth'
import { tags as tagApi } from '../../api/services'

const authStore = useAuthStore()

interface Tag {
  id: number
  name: string
  state: number
}

const loading = ref(false)
const error = ref('')
const tags = ref<Tag[]>([])
const showCreateModal = ref(false)
const editingTag = ref<Tag | null>(null)

const form = ref({
  name: '',
  state: 1,
  created_by: authStore.username || '',
  modified_by: authStore.username || ''
})

const fetchTags = async () => {
  loading.value = true
  error.value = ''
  try {
    const response = await tagApi.getList()
    if (response.code === 200) {
      tags.value = response.data
    } else {
      error.value = response.msg || 'Failed to fetch tags'
    }
  } catch (err) {
    error.value = 'An error occurred while fetching tags'
    console.error('Error fetching tags:', err)
  } finally {
    loading.value = false
  }
}

const closeModal = () => {
  showCreateModal.value = false
  editingTag.value = null
  form.value = {
    name: '',
    state: 1,
    created_by: authStore.username || '',
    modified_by: authStore.username || ''
  }
}

const editTag = (tag: Tag) => {
  editingTag.value = tag
  form.value = {
    name: tag.name,
    state: tag.state,
    created_by: authStore.username || '',
    modified_by: authStore.username || ''
  }
}

const handleSubmit = async () => {
  loading.value = true
  try {
    const response = editingTag.value
      ? await tagApi.update(editingTag.value.id, {
          name: form.value.name,
          state: form.value.state,
          modified_by: form.value.modified_by
        })
      : await tagApi.create({
          name: form.value.name,
          state: form.value.state,
          created_by: Number(form.value.created_by)
        })

    if (response.code === 200) {
      closeModal()
      await fetchTags()
    } else {
      alert(response.msg || 'Failed to save tag')
    }
  } catch (err) {
    console.error('Error saving tag:', err)
    alert('An error occurred while saving the tag')
  } finally {
    loading.value = false
  }
}

const deleteTag = async (id: number) => {
  if (!confirm('Are you sure you want to delete this tag?')) {
    return
  }

  try {
    const response = await tagApi.delete(id)
    if (response.code === 200) {
      await fetchTags()
    } else {
      alert(response.msg || 'Failed to delete tag')
    }
  } catch (err) {
    console.error('Error deleting tag:', err)
    alert('An error occurred while deleting the tag')
  }
}

const handleExport = async () => {
  try {
    const response = await tagApi.export()
    if (response.code === 200) {
      // 处理导出文件
      const blob = new Blob([response.data], { type: 'application/vnd.ms-excel' })
      const url = window.URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = 'tags.xlsx'
      document.body.appendChild(a)
      a.click()
      window.URL.revokeObjectURL(url)
      document.body.removeChild(a)
    } else {
      alert(response.msg || 'Failed to export tags')
    }
  } catch (err) {
    console.error('Error exporting tags:', err)
    alert('An error occurred while exporting tags')
  }
}

const handleImport = async (event: Event) => {
  const input = event.target as HTMLInputElement
  if (!input.files?.length) return

  const file = input.files[0]
  try {
    const response = await tagApi.import(file)
    if (response.code === 200) {
      await fetchTags()
      alert('Tags imported successfully')
    } else {
      alert(response.msg || 'Failed to import tags')
    }
  } catch (err) {
    console.error('Error importing tags:', err)
    alert('An error occurred while importing tags')
  } finally {
    input.value = '' // 清除文件选择
  }
}

// 初始加载标签
fetchTags()
</script> 