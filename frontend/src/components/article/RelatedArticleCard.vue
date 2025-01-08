<template>
  <router-link 
    :to="`/article/${article.id}`"
    class="block bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition duration-300"
  >
    <div class="md:flex">
      <div class="md:w-1/3">
        <img 
          :src="article.image" 
          :alt="article.title"
          class="h-48 w-full object-cover md:h-full"
        >
      </div>
      <div class="p-4 md:w-2/3">
        <h3 class="text-lg font-bold mb-2 line-clamp-2">{{ article.title }}</h3>
        <p class="text-gray-600 text-sm mb-2 line-clamp-2">{{ article.summary }}</p>
        <div class="flex items-center text-sm text-gray-500">
          <el-icon class="mr-1"><View /></el-icon>
          <span>{{ article.views || 0 }} 阅读</span>
          <span class="mx-2">·</span>
          <el-icon class="mr-1"><ChatDotRound /></el-icon>
          <span>{{ article.comments || 0 }} 评论</span>
        </div>
      </div>
    </div>
  </router-link>
</template>

<script setup>
import { View, ChatDotRound } from '@element-plus/icons-vue'

defineProps({
  article: {
    type: Object,
    required: true,
    validator: (article) => {
      return article && 
             typeof article.id === 'string' &&
             typeof article.title === 'string' &&
             typeof article.summary === 'string' &&
             typeof article.image === 'string'
    }
  }
})
</script>