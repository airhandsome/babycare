<template>
  <div class="py-16">
    <div class="max-w-7xl mx-auto px-4">
      <h2 class="text-3xl font-bold text-center mb-12">特色服务</h2>
      
      <div class="grid md:grid-cols-3 gap-8">
        <div v-for="service in services" 
             :key="service.id" 
             class="bg-white rounded-lg shadow-md p-6 text-center hover:shadow-lg transition-shadow cursor-pointer"
             @click="navigateToService(service.route)">
          <el-icon class="text-4xl text-indigo-600 mb-4">
            <component :is="service.icon" />
          </el-icon>
          <h3 class="text-xl font-bold mb-2">{{ service.title }}</h3>
          <p class="text-gray-600">{{ service.description }}</p>
          <el-button type="primary" plain class="mt-4">了解更多</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Sunrise, FirstAidKit, School } from '@element-plus/icons-vue'

const router = useRouter()

const services = ref([
  {
    id: 1,
    title: '孕期护理',
    description: '为准妈妈提供全面的孕期指导和营养建议',
    icon: Sunrise,
    route: '/pregnancy-care'
  },
  {
    id: 2,
    title: '新生儿护理',
    description: '专业的新生儿护理知识和实用技巧',
    icon: FirstAidKit,
    route: '/newborn-care'
  },
  {
    id: 3,
    title: '早教指导',
    description: '科学的早期教育方法和互动游戏',
    icon: School,
    route: '/preschool-education'
  }
])

const navigateToService = (route) => {
  router.push(route)
}
</script>