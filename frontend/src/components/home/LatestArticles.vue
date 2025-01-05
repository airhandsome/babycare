<template>
  <div class="py-16 bg-gray-50">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex justify-between items-center mb-8">
        <h2 class="text-3xl font-bold">最新文章</h2>
        <el-button type="primary" @click="router.push('/articles')">查看更多</el-button>
      </div>
      
      <div class="grid md:grid-cols-3 gap-8">
        <div v-for="article in articles" 
             :key="article.id" 
             class="bg-white rounded-lg shadow-md overflow-hidden cursor-pointer hover:shadow-lg transition-shadow"
             @click="goToDetail(article.id)">
          <img :src="getArticleImage(article)" 
               :alt="article.title" 
               class="w-full h-48 object-cover">
          <div class="p-6">
            <h3 class="text-xl font-bold mb-2">{{ article.title }}</h3>
            <p class="text-gray-600 mb-4 line-clamp-2">{{ article.content }}</p>
            <div class="flex justify-between items-center text-sm text-gray-500">
              <span>{{ formatDate(article.created_at) }}</span>
              <span>阅读量: {{ article.views }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import request from '../../utils/request'

const router = useRouter()
const articles = ref([])

const defaultImages = [
  '/images/articles/article1.png',
  '/images/articles/article2.png',
  '/images/articles/article3.png',
  '/images/articles/article4.png'
]

const loadLatestArticles = async () => {
  try {
    const response = await request.get('posts', {
      params: {
        page: 1,
        page_size: 6,
        sort: 'created_at desc'
      }
    })
    articles.value = response
  } catch (error) {
    console.error('Failed to load articles:', error)
  }
}

onMounted(() => {
  loadLatestArticles()
})

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getArticleImage = (article) => {
  return defaultImages[article.id % defaultImages.length]
}

const goToDetail = (id) => {
  router.push(`/article/${id}`)
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>