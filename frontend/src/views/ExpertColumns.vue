<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">专家专栏</h1>
    
    <!-- 专家列表 -->
    <section class="mb-12">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">育儿专家</h2>
        <el-button type="primary" plain @click="loadExperts">
          <el-icon><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
      <div class="grid md:grid-cols-3 gap-6">
        <ExpertCard v-for="expert in experts" 
                   :key="expert.id"
                   :name="expert.name"
                   :title="expert.title"
                   :avatar="expert.avatar"
                   :description="expert.summary"
                   :tags="expert.tags?.split(',')"
                   class="cursor-pointer hover:shadow-lg transition-shadow"
                   @click="handleExpertClick(expert)" />
      </div>
    </section>

    <!-- 最新文章 -->
    <section class="mb-12">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">专家文章</h2>
        <el-button type="primary" plain @click="loadArticles">
          <el-icon><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
      <div class="grid md:grid-cols-2 gap-8">
        <ArticleCard v-for="article in articles"
                    :key="article.id"
                    :title="article.title"
                    :summary="article.summary"
                    :author="article.user?.username"
                    :date="formatDate(article.created_at)"
                    :tags="article.tags?.split(',')"
                    :views="article.views"
                    :read-time="article.read_time"
                    class="cursor-pointer hover:shadow-lg transition-shadow"
                    @click="handleArticleClick(article)" />
      </div>
    </section>

    <!-- 热门话题 -->
    <section class="mb-12">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">热门话题</h2>
        <el-button type="primary" plain @click="loadTopics">
          <el-icon><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
      <div class="bg-white rounded-xl shadow-md p-6">
        <TopicList :topics="topics" @topic-click="handleTopicClick" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { formatDate } from '../utils/date'
import request from '../utils/request'
import ExpertCard from '../components/cards/ExpertCard.vue'
import ArticleCard from '../components/cards/ArticleCard.vue'
import TopicList from '../components/lists/TopicList.vue'

const router = useRouter()
const experts = ref([])
const articles = ref([])
const topics = ref([])

// 加载专家列表
const loadExperts = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'expert_profile' }
    })
    experts.value = response
  } catch (error) {
    console.error('Failed to load experts:', error)
    ElMessage.error('加载专家列表失败')
  }
}

// 加载专家文章
const loadArticles = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'expert_article' }
    })
    articles.value = response
  } catch (error) {
    console.error('Failed to load articles:', error)
    ElMessage.error('加载文章失败')
  }
}

// 加载热门话题
const loadTopics = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'expert_topic' }
    })
    topics.value = response
  } catch (error) {
    console.error('Failed to load topics:', error)
    ElMessage.error('加载话题失败')
  }
}

// 处理专家点击
const handleExpertClick = (expert) => {
  router.push(`/posts/${expert.id}`)
}

// 处理文章点击
const handleArticleClick = (article) => {
  router.push(`/posts/${article.id}`)
}

// 处理话题点击
const handleTopicClick = (topic) => {
  router.push(`/posts/${topic.id}`)
}

onMounted(() => {
  loadExperts()
  loadArticles()
  loadTopics()
})
</script>