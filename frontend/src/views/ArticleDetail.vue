<template>
  <div class="article-detail">
    <div class="max-w-4xl mx-auto px-4 py-8">
      <!-- 返回按钮 -->
      <button 
        @click="$router.back()" 
        class="mb-6 flex items-center text-indigo-600 hover:text-indigo-800"
      >
        <i class="fas fa-arrow-left mr-2"></i>
        返回列表
      </button>

      <!-- 文章内容 -->
      <div v-if="article" class="bg-white rounded-lg shadow-lg overflow-hidden">
        <div class="relative">
          <img :src="getArticleImage()" alt="文章配图" class="w-full h-64 object-cover">
          <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6">
            <h1 class="text-white text-3xl font-bold">{{ article.title }}</h1>
          </div>
        </div>

        <div class="p-6">
          <!-- 文章元信息卡片 -->
          <div class="bg-gray-50 rounded-lg p-4 mb-6">
            <div class="flex flex-wrap items-center gap-6 text-sm text-gray-600">
              <!-- 作者信息 -->
              <div class="flex items-center">
                <img :src="getAuthorAvatar()" alt="作者头像" class="w-10 h-10 rounded-full mr-3">
                <div>
                  <div class="font-medium text-gray-900">{{ article.user.username }}</div>
                  <div class="text-xs">孕期护理专家</div>
                </div>
              </div>

              <!-- 文章统计信息 -->
              <div class="flex items-center space-x-6">
                <span class="flex items-center">
                  <i class="fas fa-calendar mr-2"></i>
                  {{ formatDate(article.created_at) }}
                </span>
                <span class="flex items-center">
                  <i class="fas fa-eye mr-2"></i>
                  {{ article.views }} 次阅读
                </span>
                <span class="flex items-center">
                  <i class="fas fa-file-alt mr-2"></i>
                  {{ article.word_count }} 字
                </span>
                <span class="flex items-center">
                  <i class="fas fa-clock mr-2"></i>
                  预计阅读 {{ article.read_time }} 分钟
                </span>
              </div>
            </div>

            <!-- 文章标签 -->
            <div class="mt-4 flex flex-wrap gap-2">
              <span 
                v-for="tag in articleTags" 
                :key="tag"
                class="px-3 py-1 bg-indigo-100 text-indigo-800 rounded-full text-sm"
              >
                {{ tag }}
              </span>
            </div>
          </div>

          <!-- 文章正文 -->
          <div class="prose max-w-none">
            <div class="mb-6 text-gray-600 text-sm">
              <i class="fas fa-info-circle mr-2"></i>
              温馨提示：本文已经过专业医生审核，仅供参考，不作为医疗诊断依据。
            </div>
            
            <div class="article-content">
              {{ article.content }}
            </div>
          </div>

          <!-- 文章底部 -->
          <div class="mt-8 pt-6 border-t border-gray-200">
            <div class="flex justify-between items-center">
              <div class="text-gray-600 text-sm">
                最后更新于 {{ formatDate(article.updated_at) }}
              </div>
              <div class="flex space-x-4">
                <button class="flex items-center text-gray-600 hover:text-indigo-600">
                  <i class="fas fa-share-alt mr-2"></i>
                  分享
                </button>
                <button class="flex items-center text-gray-600 hover:text-indigo-600">
                  <i class="fas fa-bookmark mr-2"></i>
                  收藏
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 加载状态 -->
      <div v-else class="text-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600 mx-auto"></div>
      </div>
    </div>
  </div>
</template>

<script>
import request from '../utils/request'

export default {
  name: 'ArticleDetail',
  data() {
    return {
      article: null,
      defaultImages: [
        '/images/pregnancy/nutrition.jpg',
        '/images/pregnancy/exercise.jpg',
        '/images/pregnancy/care.jpg',
        '/images/pregnancy/health.jpg'
      ],
      defaultAvatars: [
        '/images/avatars/expert1.jpg',
        '/images/avatars/expert2.jpg'
      ]
    }
  },
  computed: {
    articleTags() {
      return this.article?.tags ? this.article.tags.split(',') : []
    }
  },
  created() {
    this.loadArticle()
  },
  methods: {
    async loadArticle() {
      try {
        const id = this.$route.params.id
        const response = await request.get(`posts/${id}`)
        this.article = response
        // 增加阅读量
        this.incrementViews(id)
      } catch (error) {
        console.error('Failed to load article:', error)
      }
    },
    async incrementViews(id) {
      try {
        await request.post(`posts/${id}/views`)
      } catch (error) {
        console.error('Failed to increment views:', error)
      }
    },
    formatDate(dateString) {
      return new Date(dateString).toLocaleDateString('zh-CN', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    },
    getArticleImage() {
      return this.article ? 
        this.defaultImages[this.article.id % this.defaultImages.length] : 
        this.defaultImages[0]
    },
    getAuthorAvatar() {
      return this.article ? 
        this.defaultAvatars[this.article.user.id % this.defaultAvatars.length] :
        this.defaultAvatars[0]
    }
  }
}
</script>

<style scoped>
.article-detail {
  background-color: #f3f4f6;
  min-height: 100vh;
}

.prose {
  white-space: pre-line;
  line-height: 1.8;
  color: #374151;
}

.article-content {
  font-size: 1.1rem;
}

.article-content p {
  margin-bottom: 1.5em;
}

/* 添加响应式样式 */
@media (max-width: 640px) {
  .flex.flex-wrap.items-center.gap-6 {
    gap: 1rem;
  }
  
  .flex.items-center.space-x-6 {
    flex-wrap: wrap;
    gap: 0.5rem;
  }
}
</style> 