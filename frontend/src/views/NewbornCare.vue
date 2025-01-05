<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">新生儿护理</h1>
    
    <!-- 喂养指导 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">喂养指导</h2>
      <div class="grid md:grid-cols-2 gap-8">
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="/images/newborn/breastfeeding.png" 
               alt="母乳喂养" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">母乳喂养指南</h3>
          <p class="text-gray-600">了解正确的母乳喂养姿势和技巧，确保宝宝获得充足营养。</p>
          <el-button type="primary" class="mt-4" @click="handleFeedingClick('breastfeeding')">查看详情</el-button>
        </div>
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="/images/newborn/formula.png" 
               alt="人工喂养" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">人工喂养建议</h3>
          <p class="text-gray-600">选择合适的奶粉和喂养工具，掌握科学的喂养方法。</p>
          <el-button type="primary" class="mt-4" @click="handleFeedingClick('formula')">查看详情</el-button>
        </div>
      </div>
    </section>

    <!-- 日常护理 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">日常护理</h2>
      <div class="grid md:grid-cols-3 gap-6">
        <CareCard v-for="item in careItems" 
                 :key="item.id"
                 :title="item.title"
                 :description="item.description"
                 :icon="item.icon"
                 :image="item.image"
                 @click="handleCardClick(item.id)" />
      </div>
    </section>

    <!-- 早期发展 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">早期发展</h2>
      <div class="bg-white rounded-xl shadow-md p-6">
        <div class="grid md:grid-cols-2 gap-8">
          <div v-for="item in developmentItems" 
               :key="item.id"
               class="flex items-start space-x-4 cursor-pointer hover:bg-gray-50 p-4 rounded-lg transition-colors"
               @click="handleDevelopmentClick(item.id)">
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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Bell, Monitor, Star, View, Calendar, Loading } from '@element-plus/icons-vue'
import CareCard from '../components/cards/CareCard.vue'
import request from '../utils/request'
import { ElMessage } from 'element-plus'

const router = useRouter()
const articles = ref([])

// 获取新生儿护理文章
const loadNewbornArticles = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'newborn' }
    })
    articles.value = response
  } catch (error) {
    ElMessage.error('加载文章失败')
    console.error('Failed to load newborn articles:', error)
  }
}

onMounted(() => {
  loadNewbornArticles()
})

const careItems = ref([
  {
    id: 1,
    title: '脐带护理',
    description: '正确的脐带护理方法和注意事项',
    icon: Bell,
    image: '/images/newborn/umbilical.png'
  },
  {
    id: 2,
    title: '洗澡技巧',
    description: '新生儿洗澡的步骤和安全事项',
    icon: Monitor,
    image: '/images/newborn/bath.png'
  },
  {
    id: 3,
    title: '换尿布指南',
    description: '如何正确更换尿布并预防红臀',
    icon: Star,
    image: '/images/newborn/diaper.png'
  }
])

const developmentItems = ref([
  {
    id: 1,
    title: '视觉发展',
    description: '了解宝宝视觉发展的关键阶段，选择适合的视觉刺激玩具。',
    icon: View
  },
  {
    id: 2,
    title: '听觉训练',
    description: '通过音乐和语言互动促进宝宝听觉系统发展。',
    icon: Bell
  },
  {
    id: 3,
    title: '运动技能',
    description: '支持宝宝抬头、翻身等早期运动技能的发展。',
    icon: Loading
  },
  {
    id: 4,
    title: '作息规律',
    description: '建立健康的睡眠和喂养规律，促进身心发展。',
    icon: Calendar
  }
])

const handleCardClick = (id) => {
  const article = articles.value.find(a => {
    switch(id) {
      case 1: return a.title.includes('脐带护理')
      case 2: return a.title.includes('洗澡')
      case 3: return a.title.includes('换尿布')
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

const handleFeedingClick = (type) => {
  const article = articles.value.find(a => 
    type === 'breastfeeding' ? 
      a.title.includes('母乳喂养') : 
      a.title.includes('人工喂养')
  )
  if (article) {
    router.push(`/article/${article.id}`)
  } else {
    ElMessage.error('未找到相关文章')
    console.error('Article not found for type:', type)
  }
}

const handleDevelopmentClick = (id) => {
  const article = articles.value.find(a => {
    switch(id) {
      case 1: return a.title.includes('视觉发展')
      case 2: return a.title.includes('听觉训练')
      case 3: return a.title.includes('运动技能')
      case 4: return a.title.includes('作息')
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