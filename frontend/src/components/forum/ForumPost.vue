<template>
  <div class="bg-white rounded-lg shadow p-6">
    <!-- 作者信息 -->
    <div class="flex items-center mb-4">
      <img 
        :src="post.user.avatar || defaultAvatar" 
        :alt="post.user.username"
        class="w-10 h-10 rounded-full mr-3"
      >
      <div>
        <h3 class="font-medium">{{ post.user.username }}</h3>
        <span class="text-sm text-gray-500">{{ formatDate(post.created_at) }}</span>
      </div>
    </div>

    <!-- 帖子内容 -->
    <div class="mb-4">
      <p class="text-gray-800">{{ post.content }}</p>
      <!-- 图片展示 -->
      <div v-if="post.images" class="mt-4 grid grid-cols-2 gap-2">
        <img 
          v-for="(image, index) in imageList" 
          :key="index"
          :src="image"
          class="rounded-lg w-full h-48 object-cover"
        >
      </div>
    </div>

    <!-- 分类标签 -->
    <div class="mb-4">
      <span class="px-3 py-1 bg-indigo-100 text-indigo-800 rounded-full text-sm">
        {{ post.category }}
      </span>
    </div>

    <!-- 互动区域 -->
    <div class="flex items-center space-x-6">
      <button 
        class="flex items-center space-x-2 hover:text-indigo-600 transition-colors"
        :class="{ 'text-indigo-600': post.isLiked, 'text-gray-500': !post.isLiked }"
        @click="handleLike"
      >
        <el-icon>
          <component :is="post.isLiked ? 'StarFilled' : 'Star'" />
        </el-icon>
        <span>{{ post.likes || 0 }}</span>
      </button>

      <button 
        class="flex items-center space-x-2 hover:text-indigo-600 transition-colors"
        :class="{ 'text-indigo-600': post.isCommented, 'text-gray-500': !post.isCommented }"
        @click="handleCommentClick"
      >
        <el-icon>
          <ChatDotRound />
        </el-icon>
        <span>{{ post.comments || 0 }}</span>
      </button>

      <button 
        class="flex items-center space-x-2 text-gray-500 hover:text-indigo-600 transition-colors"
        @click="handleShare"
      >
        <el-icon>
          <Share />
        </el-icon>
        <span>分享</span>
      </button>
    </div>

    <!-- 评论列表 -->
    <div v-if="post.commentList?.length" class="mt-4 space-y-4">
      <div class="flex items-center justify-between border-b pb-2">
        <span class="text-sm text-gray-500">全部评论 ({{ post.comments }})</span>
        <el-button type="text" @click="handleCommentClick">
          {{ showCommentInput ? '收起评论' : '写评论' }}
        </el-button>
      </div>
      <div v-for="comment in post.commentList" 
           :key="comment.id" 
           class="bg-gray-50 rounded p-4 hover:bg-gray-100 transition-colors"
      >
        <div class="flex items-start space-x-3">
          <img 
            :src="comment.user.avatar || defaultAvatar" 
            :alt="comment.user.username"
            class="w-8 h-8 rounded-full"
          >
          <div class="flex-1">
            <div class="flex items-center justify-between">
              <span class="font-medium">{{ comment.user.username }}</span>
              <span class="text-sm text-gray-500">{{ formatDate(comment.created_at) }}</span>
            </div>
            <p class="mt-1 text-gray-800">{{ comment.content }}</p>
            <div class="mt-2 flex items-center space-x-4 text-sm text-gray-500">
              <button class="hover:text-indigo-600" @click="handleReply(comment)">
                回复
              </button>
              <button 
                class="hover:text-indigo-600 flex items-center space-x-1" 
                @click="handleLikeComment(comment)"
                :class="{ 'text-indigo-600': comment.isLiked }"
              >
                <el-icon class="mr-1">
                  <component :is="comment.isLiked ? 'StarFilled' : 'Star'" />
                </el-icon>
                <span>{{ comment.likes || 0 }}</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 评论输入框 -->
    <div v-if="showCommentInput" class="mt-4">
      <div class="space-y-3">
        <el-input
          v-model="commentText"
          type="textarea"
          :rows="2"
          placeholder="写下你的评论..."
          resize="none"
        />
        <div class="flex justify-end space-x-2">
          <el-button @click="handleCommentClick">取消</el-button>
          <el-button 
            type="primary" 
            :loading="submitting"
            @click="submitComment"
          >
            发布评论
          </el-button>
        </div>
      </div>
    </div>

    <!-- 分享对话框 -->
    <el-dialog
      v-model="showShareDialog"
      title="分享帖子"
      width="400px"
    >
      <div class="space-y-4">
        <div class="flex items-center justify-between p-4 bg-gray-50 rounded">
          <div class="flex items-center space-x-4">
            <el-icon class="text-2xl text-blue-500"><Link /></el-icon>
            <span>复制链接</span>
          </div>
          <el-button type="primary" @click="copyLink">复制</el-button>
        </div>
        
        <div class="grid grid-cols-4 gap-4 text-center">
          <div v-for="item in shareOptions" 
               :key="item.type" 
               class="cursor-pointer hover:text-indigo-600"
               @click="handleShareTo(item.type)"
          >
            <el-icon class="text-2xl mb-1">
              <component :is="item.icon" />
            </el-icon>
            <div class="text-sm">{{ item.name }}</div>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { 
  Pointer, 
  ChatDotRound, 
  Share, 
  Link, 
  ChatSquare,
  Message, 
  Location,
  Star,
  StarFilled
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import request from '../../utils/request'
const props = defineProps({
  post: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['like', 'comment'])

const defaultAvatar = '/images/default-avatar.png'
const showCommentInput = ref(false)
const commentText = ref('')
const submitting = ref(false)
const showShareDialog = ref(false)

const shareOptions = [
  { type: 'weixin', name: '微信', icon: 'ChatSquare' },
  { type: 'moments', name: '朋友圈', icon: 'Location' },
  { type: 'message', name: '私信', icon: 'Message' },
  { type: 'forward', name: '转发', icon: 'Share' },
]

// 格式化日期
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// 处理图片列表
const imageList = computed(() => {
  if (!props.post.images) return []
  return props.post.images.split(',').filter(Boolean)
})

// 处理评论点击
const handleCommentClick = () => {
  showCommentInput.value = !showCommentInput.value
  // 加载评论列表
  if (showCommentInput.value && !props.post.commentList) {
    loadComments()
  }
  if (!showCommentInput.value) {
    commentText.value = ''
  }
}

// 加载评论列表
const loadComments = async () => {
  try {
    const response = await request.get(`forum/posts/${props.post.id}/comments`)
    if (response && response.data) {
      props.post.commentList = response.data
      props.post.comments = response.total
    }
  } catch (error) {
    console.error('Failed to load comments:', error)
    ElMessage.error('加载评论失败')
  }
}

// 提交评论
const submitComment = async () => {
  if (!commentText.value.trim()) {
    ElMessage.warning('请输入评论内容')
    return
  }

  submitting.value = true
  try {
    const response = await request.post(`forum/posts/${props.post.id}/comments`, {
      content: commentText.value.trim()
    })
    
    commentText.value = ''
    props.post.comments++
    props.post.isCommented = true
    
    // 将新评论添加到列表开头
    if (!props.post.commentList) {
      props.post.commentList = []
    }
    props.post.commentList.unshift(response)
    
    ElMessage.success('评论成功')
  } catch (error) {
    console.error('Failed to submit comment:', error)
    ElMessage.error('评论失败')
  } finally {
    submitting.value = false
  }
}

// 处理分享
const handleShare = () => {
  showShareDialog.value = true
}

const copyLink = () => {
  const url = `${window.location.origin}/forum/post/${props.post.id}`
  navigator.clipboard.writeText(url)
    .then(() => {
      ElMessage.success('链接已复制到剪贴板')
      showShareDialog.value = false
    })
    .catch(() => {
      ElMessage.error('复制失败，请手动复制')
    })
}

const handleShareTo = (type) => {
  // 根据不同的分享类型处理
  switch (type) {
    case 'weixin':
      ElMessage.info('请使用微信扫描二维码分享')
      break
    case 'moments':
      ElMessage.info('请在微信中打开分享到朋友圈')
      break
    case 'message':
      ElMessage.info('私信功能开发中...')
      break
    case 'forward':
      ElMessage.info('转发功能开发中...')
      break
  }
  showShareDialog.value = false
}

// 添加点赞处理函数
const handleLike = async () => {
  try {
    await emit('like', props.post.id)
    props.post.isLiked = !props.post.isLiked
  } catch (error) {
    ElMessage.error('点赞失败')
  }
}

// 处理评论点赞
const handleLikeComment = async (comment) => {
  try {
    const response = await request.post(`forum/comments/${comment.id}/like`)
    // 更新评论数据
    Object.assign(comment, response)
    ElMessage.success('点赞成功')
  } catch (error) {
    console.error('Failed to like comment:', error)
    ElMessage.error('点赞失败')
  }
}

// 处理评论回复
const handleReply = (comment) => {
  commentText.value = `@${comment.user.username} `
  showCommentInput.value = true
}
</script>

<style scoped>
.el-icon {
  font-size: 1.25rem;
}

:deep(.el-input-group__append) {
  padding: 0;
}

:deep(.el-button) {
  border: none;
  height: 100%;
}
</style>