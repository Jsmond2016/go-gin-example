<template>
  <div class="bg-white shadow overflow-hidden sm:rounded-lg">
    <div class="px-4 py-5 sm:px-6">
      <h3 class="text-lg leading-6 font-medium text-gray-900">
        Welcome to Blog Admin
      </h3>
      <p class="mt-1 max-w-2xl text-sm text-gray-500">
        Manage your articles and tags here.
      </p>
    </div>
    <div class="border-t border-gray-200">
      <dl>
        <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
          <dt class="text-sm font-medium text-gray-500">Articles</dt>
          <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
            <router-link
              to="/articles"
              class="text-indigo-600 hover:text-indigo-900"
            >
              Manage Articles
            </router-link>
          </dd>
        </div>
        <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
          <dt class="text-sm font-medium text-gray-500">Tags</dt>
          <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
            <router-link
              to="/tags"
              class="text-indigo-600 hover:text-indigo-900"
            >
              Manage Tags
            </router-link>
          </dd>
        </div>
      </dl>
    </div>
  </div>
</template> 