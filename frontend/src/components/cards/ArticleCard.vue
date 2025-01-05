<template>
  <div class="bg-white rounded-xl shadow-md overflow-hidden">
    <div class="md:flex">
      <div class="md:w-1/3">
        <img :src="image" :alt="title" class="h-full w-full object-cover">
      </div>
      <div class="p-6 md:w-2/3">
        <div class="flex items-center mb-2">
          <span class="text-gray-500">{{ date }}</span>
          <span class="mx-2">·</span>
          <span class="text-blue-600">{{ author }}</span>
        </div>
        <h3 class="text-xl font-bold mb-2">{{ title }}</h3>
        <p class="text-gray-600 mb-4">{{ summary }}</p>
        <div class="flex flex-wrap gap-2">
          <el-tag v-for="tag in tags" :key="tag" effect="plain">{{ tag }}</el-tag>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: String,
  author: String,
  date: String,
  summary: String,
  image: String,
  tags: Array
})
</script>