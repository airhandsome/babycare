<template>
  <nav class="bg-white shadow-lg">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex justify-between h-16">
        <div class="flex">
          <router-link to="/" class="flex items-center">
            <span class="text-2xl font-bold text-blue-600">BabyCare</span>
          </router-link>
        </div>
        
        <div class="hidden md:flex items-center space-x-4">
          <router-link v-for="item in menuItems" 
                      :key="item.path" 
                      :to="item.path"
                      class="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
            {{ item.name }}
          </router-link>
        </div>

        <div class="flex items-center space-x-4">
          <el-input
            v-model="searchQuery"
            placeholder="搜索..."
            class="w-48"
            :prefix-icon="Search"
          />
          <el-button type="primary" @click="showLoginDialog">登录</el-button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'

const searchQuery = ref('')
const menuItems = [
  { name: '首页', path: '/' },
  { name: '孕期护理', path: '/pregnancy-care' },
  { name: '新生儿护理', path: '/newborn-care' },
  { name: '婴幼儿成长', path: '/infant-growth' },
  { name: '幼儿教育', path: '/preschool' },
  { name: '专家专栏', path: '/expert-columns' },
  { name: '社区论坛', path: '/community' },
]

const showLoginDialog = () => {
  // TODO: Implement login dialog
}
</script>