<template>
  <div class="bg-white rounded-xl p-6 mb-6">
    <!-- 基本信息 -->
    <div class="flex items-start space-x-6 mb-6">
      <el-avatar :src="expert.avatar" :size="80" />
      <div>
        <div class="flex items-center space-x-3 mb-2">
          <h3 class="text-xl font-bold">{{ expert.name }}</h3>
          <el-tag type="success">{{ expert.title }}</el-tag>
        </div>
        <div class="flex items-center space-x-2 mb-2">
          <el-rate v-model="expert.rating" disabled :max="5" />
          <span class="text-gray-600">({{ expert.rating }})</span>
        </div>
        <p class="text-blue-600 font-medium">
          咨询费用：¥{{ expert.price }}/次
        </p>
      </div>
    </div>

    <!-- 专业领域 -->
    <div class="mb-6">
      <h4 class="text-lg font-bold mb-3">专业领域</h4>
      <div class="flex flex-wrap gap-2">
        <el-tag 
          v-for="specialty in expert.specialties" 
          :key="specialty"
          type="info"
        >
          {{ specialty }}
        </el-tag>
      </div>
    </div>

    <!-- 咨询统计 -->
    <div class="grid grid-cols-3 gap-4 text-center">
      <div class="bg-gray-50 rounded-lg p-3">
        <div class="text-2xl font-bold text-blue-600">
          {{ expert.consultationCount }}
        </div>
        <div class="text-gray-600">已服务</div>
      </div>
      <div class="bg-gray-50 rounded-lg p-3">
        <div class="text-2xl font-bold text-blue-600">
          {{ responseRate }}%
        </div>
        <div class="text-gray-600">回复率</div>
      </div>
      <div class="bg-gray-50 rounded-lg p-3">
        <div class="text-2xl font-bold text-blue-600">
          {{ avgResponseTime }}分钟
        </div>
        <div class="text-gray-600">平均响应</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  expert: {
    type: Object,
    required: true,
    validator: (expert) => {
      return expert && 
             typeof expert.name === 'string' &&
             typeof expert.title === 'string' &&
             typeof expert.avatar === 'string' &&
             Array.isArray(expert.specialties) &&
             typeof expert.rating === 'number' &&
             typeof expert.consultationCount === 'number' &&
             typeof expert.price === 'number'
    }
  }
})

const responseRate = computed(() => 98)
const avgResponseTime = computed(() => 15)
</script>