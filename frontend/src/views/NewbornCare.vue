<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">新生儿护理</h1>
    
    <!-- 喂养指导 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">喂养指导</h2>
      <div class="grid md:grid-cols-2 gap-8">
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="https://images.unsplash.com/photo-1590612847213-5c54a3c3f103" 
               alt="母乳喂养" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">母乳喂养指南</h3>
          <p class="text-gray-600">了解正确的母乳喂养姿势和技巧，确保宝宝获得充足营养。</p>
          <el-button type="primary" class="mt-4">查看详情</el-button>
        </div>
        <div class="bg-white rounded-xl shadow-md p-6">
          <img src="https://images.unsplash.com/photo-1594612076467-d532b39c6559" 
               alt="人工喂养" 
               class="w-full h-48 object-cover rounded-lg mb-4">
          <h3 class="text-xl font-bold mb-2">人工喂养建议</h3>
          <p class="text-gray-600">选择合适的奶粉和喂养工具，掌握科学的喂养方法。</p>
          <el-button type="primary" class="mt-4">查看详情</el-button>
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
                 :image="item.image" />
      </div>
    </section>

    <!-- 早期发展 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">早期发展</h2>
      <div class="bg-white rounded-xl shadow-md p-6">
        <div class="grid md:grid-cols-2 gap-8">
          <div v-for="item in developmentItems" 
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
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Bell, Monitor, Star, View, Calendar, Loading } from '@element-plus/icons-vue'
import CareCard from '../components/cards/CareCard.vue'

const careItems = ref([
  {
    id: 1,
    title: '脐带护理',
    description: '正确的脐带护理方法和注意事项',
    icon: Bell,
    image: 'https://images.unsplash.com/photo-1555252333-9f8e92e65df9'
  },
  {
    id: 2,
    title: '洗澡技巧',
    description: '新生儿洗澡的步骤和安全事项',
    icon: Monitor,
    image: 'https://images.unsplash.com/photo-1544829099-b9a0c07fad1a'
  },
  {
    id: 3,
    title: '换尿布指南',
    description: '如何正确更换尿布并预防红臀',
    icon: Star,
    image: 'https://images.unsplash.com/photo-1596178065887-1198b6148b2b'
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
</script>