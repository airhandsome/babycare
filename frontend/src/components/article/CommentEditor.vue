<template>
  <div class="bg-white rounded-lg shadow-md p-6 mb-8">
    <el-form @submit.prevent="handleSubmit">
      <el-input
        v-model="content"
        type="textarea"
        :rows="4"
        placeholder="写下您的评论..."
        class="mb-4"
      />
      <div class="flex justify-between items-center">
        <div class="flex items-center space-x-2">
          <el-upload
            action="#"
            :show-file-list="false"
            :on-change="handleImageUpload"
            accept="image/*"
          >
            <el-button text>
              <el-icon><Picture /></el-icon>
              添加图片
            </el-button>
          </el-upload>
        </div>
        <el-button type="primary" @click="handleSubmit">发布评论</el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Picture } from '@element-plus/icons-vue'

const emit = defineEmits(['submit'])

const content = ref('')
const images = ref([])

const handleImageUpload = (file) => {
  images.value.push(file.url)
}

const handleSubmit = () => {
  if (!content.value.trim()) return
  
  emit('submit', {
    content: content.value,
    images: images.value,
    date: new Date().toISOString()
  })

  // Reset form
  content.value = ''
  images.value = []
}
</script>