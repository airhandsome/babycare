<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">社区论坛</h1>
    
    <!-- 发帖区域 -->
    <PostCreator @post-created="handlePostCreated" />

    <!-- 分类筛选 -->
    <CategoryFilter 
      :categories="categories" 
      v-model="selectedCategory" 
      class="mb-8"
    />

    <!-- 帖子列表 -->
    <div v-if="posts.length > 0" class="space-y-6">
      <ForumPost 
        v-for="post in posts" 
        :key="post.id"
        :post="post"
        @like="handleLike"
        @comment="handleComment"
      />
      
      <!-- 分页控件 -->
      <div class="flex justify-center mt-8">
        <el-pagination
          v-model:current-page="currentPage"
          :page-size="pageSize"
          :total="total"
          :background="true"
          layout="prev, pager, next"
          @current-change="handlePageChange"
        />
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-else-if="loading" class="text-center py-12">
      <el-spinner class="inline-block" />
      <p class="mt-4 text-gray-600">加载中...</p>
    </div>

    <!-- 空状态 -->
    <div v-else class="text-center py-12">
      <el-empty description="暂无帖子" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import PostCreator from '../components/forum/PostCreator.vue'
import CategoryFilter from '../components/forum/CategoryFilter.vue'
import ForumPost from '../components/forum/ForumPost.vue'
import request from '../utils/request'
import { ElMessage } from 'element-plus'

const posts = ref([])
const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(20)
const total = ref(0)
const selectedCategory = ref('all')

const categories = [
  { id: 'all', name: '全部' },
  { id: '育儿经验', name: '育儿经验' },
  { id: '问题咨询', name: '问题咨询' },
  { id: '专业指导', name: '专业指导' },
  { id: '亲子活动', name: '亲子活动' }
]

// 加载帖子列表
const loadPosts = async () => {
  loading.value = true
  try {
    const response = await request.get('forum/posts', {
      params: {
        page: currentPage.value,
        page_size: pageSize.value,
        category: selectedCategory.value === 'all' ? '' : selectedCategory.value
      }
    })
    if (response && response.data) {
      posts.value = response.data
      total.value = response.total
      // 调试输出
      console.log('Posts loaded:', posts.value)
    } else {
      posts.value = []
      total.value = 0
      console.warn('No data in response:', response)
    }
  } catch (error) {
    console.error('Failed to load posts:', error)
    ElMessage.error('加载帖子失败')
  } finally {
    loading.value = false
  }
}

// 处理分页变化
const handlePageChange = (page) => {
  currentPage.value = page
  loadPosts()
}

// 监听分类变化
watch(selectedCategory, () => {
  console.log('Category changed to:', selectedCategory.value)
  currentPage.value = 1
  loadPosts()
})

// 处理点赞
const handleLike = async (postId) => {
  try {
    await request.post(`forum/posts/${postId}/like`)
    // 更新帖子点赞数
    const post = posts.value.find(p => p.id === postId)
    if (post) {
      post.likes++
    }
    ElMessage.success('点赞成功')
  } catch (error) {
    console.error('Failed to like post:', error)
    ElMessage.error('点赞失败')
  }
}

// 处理评论
const handleComment = async (postId, commentData) => {
  try {
    const response = await request.post(`forum/posts/${postId}/comments`, commentData)
    // 更新评论列表
    loadComments(postId)
    ElMessage.success('评论成功')
  } catch (error) {
    console.error('Failed to add comment:', error)
    ElMessage.error('评论失败')
  }
}

// 加载评论列表
const loadComments = async (postId) => {
  try {
    const response = await request.get(`forum/posts/${postId}/comments`)
    const post = posts.value.find(p => p.id === postId)
    if (post) {
      post.commentList = response.data
      post.comments = response.data.length
    }
  } catch (error) {
    console.error('Failed to load comments:', error)
    ElMessage.error('加载评论失败')
  }
}

// 处理新帖子创建
const handlePostCreated = (newPost) => {
  posts.value.unshift(newPost)
  total.value++
  ElMessage.success('发帖成功')
}

onMounted(() => {
  console.log('Component mounted, loading posts...')
  loadPosts()
})
</script>

<style scoped>
.el-pagination {
  --el-pagination-button-bg-color: white;
  --el-pagination-hover-color: #4f46e5;
}
</style>