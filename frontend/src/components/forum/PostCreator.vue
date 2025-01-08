<template>
  <div class="bg-white rounded-lg shadow p-6 mb-8">
    <div class="mb-4">
      <el-input
        v-model="content"
        type="textarea"
        :rows="4"
        placeholder="分享您的育儿经验和问题..."
        resize="none"
      />
    </div>

    <div class="flex items-center justify-between">
      <div class="flex items-center space-x-4">
        <!-- 分类选择 -->
        <el-select v-model="category" placeholder="选择分类" class="w-32">
          <el-option
            v-for="item in categories"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>

        <!-- 图片上传 -->
        <el-upload
          action="#"
          list-type="picture-card"
          :auto-upload="false"
          :on-change="handleImageChange"
          :on-remove="handleImageRemove"
          :limit="4"
        >
          <template #default>
            <i class="fas fa-plus"></i>
          </template>
        </el-upload>
      </div>

      <!-- 发布按钮 -->
      <el-button 
        type="primary" 
        :loading="publishing"
        :disabled="!isValid"
        @click="handlePublish"
      >
        发布
      </el-button>
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="mt-2 text-red-500 text-sm">
      {{ error }}
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import request from '../../utils/request'
import { ElMessage } from 'element-plus'

const emit = defineEmits(['post-created'])

const content = ref('')
const category = ref('')
const images = ref([])
const publishing = ref(false)
const error = ref('')

const categories = [
  { id: '育儿经验', name: '育儿经验' },
  { id: '问题咨询', name: '问题咨询' },
  { id: '专业指导', name: '专业指导' },
  { id: '亲子活动', name: '亲子活动' }
]

// 表单验证
const isValid = computed(() => {
  return content.value.trim() && category.value
})

// 处理图片变化
const handleImageChange = (file) => {
  // 这里可以添加图片预览、压缩等处理
  if (file.status === 'ready') {
    images.value.push(file.raw)
  }
}

// 处理图片移除
const handleImageRemove = (file) => {
  const index = images.value.indexOf(file.raw)
  if (index > -1) {
    images.value.splice(index, 1)
  }
}

// 上传图片到服务器
const uploadImages = async () => {
  if (images.value.length === 0) return ''

  const formData = new FormData()
  images.value.forEach(image => {
    formData.append('images', image)
  })

  try {
    const response = await request.post('upload/images', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    return response.urls.join(',')
  } catch (error) {
    console.error('Failed to upload images:', error)
    throw new Error('图片上传失败')
  }
}

// 发布帖子
const handlePublish = async () => {
  if (!isValid.value) {
    error.value = '请填写内容并选择分类'
    return
  }

  publishing.value = true
  error.value = ''

  try {
    let imageUrls = ''
    if (images.value.length > 0) {
      imageUrls = await uploadImages()
    }

    const response = await request.post('forum/posts', {
      content: content.value.trim(),
      category: category.value,
      images: imageUrls
    })

    // 发布成功，清空表单
    content.value = ''
    category.value = ''
    images.value = []
    
    // 通知父组件
    emit('post-created', response)
    ElMessage.success('发布成功')
  } catch (err) {
    console.error('Failed to publish post:', err)
    error.value = '发布失败，请稍后重试'
    ElMessage.error('发布失败')
  } finally {
    publishing.value = false
  }
}
</script>

<style scoped>
:deep(.el-upload--picture-card) {
  width: 80px;
  height: 80px;
  line-height: 80px;
}

:deep(.el-upload-list--picture-card .el-upload-list__item) {
  width: 80px;
  height: 80px;
}
</style>