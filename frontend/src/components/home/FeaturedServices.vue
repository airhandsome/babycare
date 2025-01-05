<template>
  <section class="py-16 bg-gray-50">
    <div class="max-w-7xl mx-auto px-4">
      <h2 class="text-3xl font-bold mb-8 text-center">特色服务</h2>
      <div class="grid md:grid-cols-3 gap-8">
        <div v-for="service in services" 
             :key="service.id" 
             class="bg-white p-8 rounded-xl shadow-md text-center transform transition duration-300 hover:shadow-lg">
          <el-icon class="text-4xl text-blue-600 mb-6">
            <component :is="service.icon" />
          </el-icon>
          <h3 class="text-xl font-bold mb-4">{{ service.title }}</h3>
          <p class="text-gray-600 mb-4">{{ service.description }}</p>
          <el-button type="primary" plain>了解更多</el-button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from 'vue'
import { ChatDotRound, User, Calendar } from '@element-plus/icons-vue'

const services = ref([
  {
    id: 1,
    title: '在线咨询',
    description: '随时与育儿专家在线交流，解答您的疑惑，提供专业建议。',
    icon: ChatDotRound
  },
  {
    id: 2,
    title: '专家问答',
    description: '资深儿科医生、心理专家为您答疑解惑，提供权威指导。',
    icon: User
  },
  {
    id: 3,
    title: '成长记录',
    description: '记录宝宝的每个重要时刻，科学追踪成长轨迹。',
    icon: Calendar
  }
])
</script>