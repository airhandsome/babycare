<template>
  <section class="py-16">
    <div class="max-w-7xl mx-auto px-4">
      <h2 class="text-3xl font-bold mb-8 text-center">最新文章</h2>
      <div class="grid md:grid-cols-3 gap-8">
        <div v-for="article in articles" 
             :key="article.id" 
             class="bg-white rounded-xl shadow-md overflow-hidden transform transition duration-300 hover:scale-105">
          <img :src="article.image" :alt="article.title" class="w-full h-48 object-cover">
          <div class="p-6">
            <div class="text-sm text-blue-600 mb-2">{{ article.category }}</div>
            <h3 class="text-xl font-bold mb-2 text-gray-900">{{ article.title }}</h3>
            <p class="text-gray-600 mb-4">{{ article.excerpt }}</p>
            <router-link :to="article.link" class="text-blue-600 hover:text-blue-800 font-medium">
              阅读更多 →
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from 'vue'

const articles = ref([
  {
    id: 1,
    title: '婴儿睡眠指南：建立健康作息',
    excerpt: '了解如何帮助宝宝建立规律的睡眠习惯，促进健康成长。',
    category: '新生儿护理',
    image: 'https://images.unsplash.com/photo-1555252333-9f8e92e65df9',
    link: '/article/1'
  },
  {
    id: 2,
    title: '科学添加辅食的关键时期',
    excerpt: '详细解读6-12个月婴儿的辅食添加指南，让宝宝健康成长。',
    category: '婴儿营养',
    image: 'https://images.unsplash.com/photo-1544829099-b9a0c07fad1a',
    link: '/article/2'
  },
  {
    id: 3,
    title: '早教游戏：激发宝宝潜能',
    excerpt: '精选适合0-3岁婴幼儿的互动游戏，促进感知和认知发展。',
    category: '早期教育',
    image: 'https://images.unsplash.com/photo-1596178065887-1198b6148b2b',
    link: '/article/3'
  }
])
</script>