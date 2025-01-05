<template>
  <div class="bg-white rounded-xl shadow-md overflow-hidden">
    <img :src="image" :alt="title" class="w-full h-48 object-cover">
    <div class="p-6">
      <h3 class="text-xl font-bold mb-2">{{ title }}</h3>
      <p class="text-gray-600 mb-4">{{ description }}</p>
      <div class="flex items-center mb-2">
        <el-icon class="mr-2"><Timer /></el-icon>
        <span>{{ duration }}</span>
      </div>
      <div class="flex flex-wrap gap-2">
        <el-tag v-for="material in materials" :key="material" type="success">
          {{ material }}
        </el-tag>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Timer } from '@element-plus/icons-vue'

defineProps({
  title: String,
  description: String,
  image: String,
  duration: String,
  materials: Array
})
</script>