<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">孕期护理</h1>
    
    <!-- 营养与健康 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">营养与健康</h2>
      <div class="grid md:grid-cols-2 gap-8">
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="https://images.unsplash.com/photo-1493770348161-369560ae357d" 
               alt="孕期营养" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">孕期营养指南</h3>
          <p class="text-gray-600">了解孕期各阶段的营养需求，为宝宝的健康发展打下基础。</p>
        </div>
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="https://images.unsplash.com/photo-1518611012118-696072aa579a" 
               alt="孕期运动" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">安全运动建议</h3>
          <p class="text-gray-600">合适的运动可以帮助缓解不适，保持身体健康。</p>
        </div>
      </div>
    </section>

    <!-- 产前准备 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">产前准备</h2>
      <div class="bg-white rounded-xl shadow-md p-6">
        <div class="grid md:grid-cols-3 gap-6">
          <div v-for="item in preparationItems" :key="item.id" class="text-center">
            <el-icon class="text-4xl text-blue-600 mb-4">
              <component :is="item.icon" />
            </el-icon>
            <h3 class="font-bold mb-2">{{ item.title }}</h3>
            <p class="text-gray-600">{{ item.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- 互动社区 -->
    <section>
      <h2 class="text-2xl font-bold mb-6">互动社区</h2>
      <div class="bg-white rounded-xl shadow-md p-6">
        <div class="mb-6">
          <el-input
            v-model="newPost"
            type="textarea"
            placeholder="分享您的经验和问题..."
            :rows="3"
          />
          <div class="mt-4 text-right">
            <el-button type="primary">发布</el-button>
          </div>
        </div>
        <div class="space-y-6">
          <div v-for="post in communityPosts" :key="post.id" class="border-b pb-6">
            <div class="flex items-start space-x-4">
              <el-avatar :src="post.avatar" />
              <div class="flex-1">
                <div class="flex justify-between items-center mb-2">
                  <span class="font-bold">{{ post.username }}</span>
                  <span class="text-gray-500 text-sm">{{ post.time }}</span>
                </div>
                <p class="text-gray-700">{{ post.content }}</p>
                <div class="mt-2 flex items-center space-x-4">
                  <el-button text>
                    <el-icon><ChatDotRound /></el-icon>
                    回复 ({{ post.replies }})
                  </el-button>
                  <el-button text>
                    <el-icon><Star /></el-icon>
                    收藏
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ChatDotRound, Star, Location, Timer, Box } from '@element-plus/icons-vue'

const newPost = ref('')

const preparationItems = ref([
  {
    id: 1,
    title: '医院选择',
    description: '了解各医院的特点，选择适合的分娩医院',
    icon: Location
  },
  {
    id: 2,
    title: '待产准备',
    description: '待产包清单和注意事项',
    icon: Box
  },
  {
    id: 3,
    title: '分娩知识',
    description: '了解分娩过程和注意事项',
    icon: Timer
  }
])

const communityPosts = ref([
  {
    id: 1,
    username: '准妈妈小王',
    avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80',
    content: '最近总是失眠，有什么好的解决办法吗？',
    time: '2小时前',
    replies: 5
  },
  {
    id: 2,
    username: '快乐妈咪',
    avatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330',
    content: '分享一下我的孕期营养餐谱，希望对大家有帮助！',
    time: '4小时前',
    replies: 8
  }
])
</script>