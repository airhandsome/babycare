<template>
  <div class="grid md:grid-cols-3 gap-6">
    <div v-for="guide in guides" 
         :key="guide.id"
         class="bg-white rounded-xl shadow-md overflow-hidden">
      <img :src="guide.image" 
           :alt="guide.title"
           class="w-full h-48 object-cover">
      <div class="p-6">
        <div class="flex items-center mb-4">
          <el-icon class="text-2xl text-blue-600 mr-3">
            <component :is="guide.icon" />
          </el-icon>
          <h3 class="text-xl font-bold">{{ guide.title }}</h3>
        </div>
        <p class="text-gray-600">{{ guide.content }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ChatRound, Position } from '@element-plus/icons-vue'

defineProps({
  guides: {
    type: Array,
    required: true
  }
})
</script>