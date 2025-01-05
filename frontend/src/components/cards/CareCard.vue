<template>
  <div class="bg-white rounded-xl shadow-md overflow-hidden transform transition duration-300 hover:shadow-lg">
    <img :src="image" :alt="title" class="w-full h-48 object-cover">
    <div class="p-6">
      <div class="flex items-center mb-4">
        <el-icon class="text-2xl text-blue-600 mr-3">
          <component :is="icon" />
        </el-icon>
        <h3 class="text-xl font-bold">{{ title }}</h3>
      </div>
      <p class="text-gray-600 mb-4">{{ description }}</p>
      <el-button type="primary" plain>了解更多</el-button>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    required: true
  },
  description: {
    type: String,
    required: true
  },
  icon: {
    type: Object,
    required: true
  },
  image: {
    type: String,
    required: true
  }
})
</script>