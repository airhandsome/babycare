<template>
  <div class="bg-white rounded-xl shadow-md overflow-hidden">
    <div class="p-6">
      <div class="flex items-center mb-4">
        <img :src="avatar" :alt="name" class="w-16 h-16 rounded-full object-cover">
        <div class="ml-4">
          <h3 class="text-xl font-bold">{{ name }}</h3>
          <p class="text-gray-600">{{ title }}</p>
        </div>
      </div>
      <p class="text-gray-700 mb-4">{{ description }}</p>
      <div class="flex flex-wrap gap-2">
        <el-tag v-for="tag in tags" :key="tag" size="small">
          {{ tag }}
        </el-tag>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  name: String,
  title: String,
  avatar: String,
  description: String,
  tags: Array
})
</script>