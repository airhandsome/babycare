<template>
  <div class="bg-white rounded-xl shadow-md p-6">
    <div class="flex items-center mb-4">
      <el-avatar :size="64" :src="avatar" />
      <div class="ml-4">
        <h3 class="text-xl font-bold">{{ name }}</h3>
        <p class="text-gray-600">{{ title }}</p>
      </div>
    </div>
    <p class="text-gray-600 mb-4">{{ description }}</p>
    <div class="flex flex-wrap gap-2">
      <el-tag v-for="specialty in specialties" 
              :key="specialty" 
              type="info">
        {{ specialty }}
      </el-tag>
    </div>
  </div>
</template>

<script setup>
defineProps({
  name: String,
  title: String,
  avatar: String,
  specialties: Array,
  description: String
})
</script>