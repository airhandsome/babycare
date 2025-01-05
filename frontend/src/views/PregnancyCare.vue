<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">孕期护理</h1>
    
    <!-- 营养指导 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">营养指导</h2>
      <div class="grid md:grid-cols-2 gap-8">
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="/images/pregnancy/nutrition-guide.png" 
               alt="孕期营养" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">孕期营养食谱</h3>
          <p class="text-gray-600">为准妈妈提供全面的营养搭配建议，确保母婴健康。</p>
          <el-button type="primary" class="mt-4" @click="navigateToArticle('nutrition')">查看详情</el-button>
        </div>
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="/images/pregnancy/supplements.png" 
               alt="营养补充" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">营养补充指南</h3>
          <p class="text-gray-600">了解孕期必需的营养补充剂，科学补充营养。</p>
          <el-button type="primary" class="mt-4" @click="navigateToArticle('supplements')">查看详情</el-button>
        </div>
      </div>
    </section>

    <!-- 运动建议 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">运动建议</h2>
      <div class="grid md:grid-cols-3 gap-6">
        <CareCard v-for="item in exerciseItems" 
                 :key="item.id"
                 :title="item.title"
                 :description="item.description"
                 :icon="item.icon"
                 :image="item.image"
                 @click="navigateToArticle(item.id)" />
      </div>
    </section>

    <!-- 健康监测 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">健康监测</h2>
      <div class="bg-white rounded-xl shadow-md p-6">
        <div class="grid md:grid-cols-2 gap-8">
          <div v-for="item in healthItems" 
               :key="item.id"
               class="flex items-start space-x-4">
            <el-icon class="text-3xl text-blue-600 mt-1">
              <component :is="item.icon" />
            </el-icon>
            <div>
              <h3 class="text-lg font-bold mb-2">{{ item.title }}</h3>
              <p class="text-gray-600">{{ item.description }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 文章列表 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">最新文章</h2>
      <div class="grid md:grid-cols-3 gap-6">
        <div v-for="post in posts" 
             :key="post.id" 
             class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300 cursor-pointer"
             @click="goToDetail(post.id)">
          <div class="relative">
            <img :src="getArticleImage(post)" alt="文章配图" class="w-full h-48 object-cover">
            <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-4">
              <h2 class="text-white text-xl font-semibold">{{ post.title }}</h2>
            </div>
          </div>
          <div class="p-4">
            <p class="text-gray-600 line-clamp-3">{{ formatContent(post.content) }}</p>
            <div class="mt-4 flex items-center justify-between text-sm text-gray-500">
              <span class="flex items-center">
                <i class="fas fa-user mr-2"></i>
                {{ post.user.username }}
              </span>
              <span class="flex items-center">
                <i class="fas fa-calendar mr-2"></i>
                {{ formatDate(post.created_at) }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- 分页控件 -->
      <div class="flex justify-center mt-8 space-x-2">
        <button 
          @click="loadPosts(currentPage - 1)"
          :disabled="currentPage === 1"
          class="px-4 py-2 bg-indigo-600 text-white rounded-md disabled:bg-gray-400 hover:bg-indigo-700 transition-colors"
        >
          上一页
        </button>
        <span class="px-4 py-2 bg-white rounded-md shadow">第 {{ currentPage }} 页</span>
        <button 
          @click="loadPosts(currentPage + 1)"
          :disabled="!hasMorePages"
          class="px-4 py-2 bg-indigo-600 text-white rounded-md disabled:bg-gray-400 hover:bg-indigo-700 transition-colors"
        >
          下一页
        </button>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Bell, Monitor, Star, View, Calendar, Loading } from '@element-plus/icons-vue'
import CareCard from '../components/cards/CareCard.vue'
import request from '../utils/request'

const router = useRouter()
const posts = ref([])
const articles = ref([])
const currentPage = ref(1)
const pageSize = ref(6)
const hasMorePages = ref(true)

const exerciseItems = ref([
  {
    id: 'yoga',
    title: '孕期瑜伽',
    description: '安全的孕期瑜伽动作和注意事项',
    icon: Bell,
    image: '/images/pregnancy/yoga.png'
  },
  {
    id: 'walking',
    title: '散步指南',
    description: '合适的散步时间和强度建议',
    icon: Monitor,
    image: '/images/pregnancy/walking.png'
  },
  {
    id: 'swimming',
    title: '游泳建议',
    description: '孕期游泳的好处和安全事项',
    icon: Star,
    image: '/images/pregnancy/swimming.png'
  }
])

const healthItems = ref([
  {
    id: 1,
    title: '体重管理',
    description: '了解孕期合理的体重增长范围，保持健康的体重管理。',
    icon: View
  },
  {
    id: 2,
    title: '产检跟踪',
    description: '定期产检的重要性和检查项目说明。',
    icon: Bell
  },
  {
    id: 3,
    title: '营养监测',
    description: '监测营养摄入，确保母婴所需的各类营养素。',
    icon: Loading
  },
  {
    id: 4,
    title: '心理调节',
    description: '保持良好的心理状态，应对孕期情绪变化。',
    icon: Calendar
  }
])

const defaultImages = [
  '/images/pregnancy/care.png',
  '/images/pregnancy/exercise.png',
  '/images/pregnancy/care.png',
  '/images/pregnancy/health.png'
]

const loadPregnancyArticles = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'pregnancy' }
    })
    articles.value = response
  } catch (error) {
    console.error('Failed to load pregnancy articles:', error)
  }
}

onMounted(() => {
  loadPosts(1)
  loadPregnancyArticles()
})

const loadPosts = async (page) => {
  try {
    const response = await request.get('posts', {
      params: {
        page: page,
        page_size: pageSize.value,
        category: 'pregnancy'
      }
    })
    posts.value = response
    currentPage.value = page
    hasMorePages.value = response.length === pageSize.value
  } catch (error) {
    console.error('Failed to load posts:', error)
  }
}

const formatContent = (content) => {
  return content.length > 100 ? content.slice(0, 100) + '...' : content
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getArticleImage = (post) => {
  return defaultImages[post.id % defaultImages.length]
}

const goToDetail = (id) => {
  router.push(`/article/${id}`)
}

const navigateToArticle = (id) => {
  const article = articles.value.find(a => {
    switch(id) {
      case 'nutrition': return a.title.includes('营养食谱')
      case 'supplements': return a.title.includes('营养补充剂')
      case 'yoga': return a.title.includes('瑜伽')
      case 'walking': return a.title.includes('散步')
      case 'swimming': return a.title.includes('游泳')
      default: return false
    }
  })
  if (article) {
    router.push(`/article/${article.id}`)
  } else {
    ElMessage.error('未找到相关文章')
    console.error('Article not found for id:', id)
  }
}
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>