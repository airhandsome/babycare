<template>
  <div class="space-y-4">
    <div v-for="topic in topics" 
         :key="topic.id"
         class="p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors cursor-pointer"
         @click="$emit('topic-click', topic)">
      <div class="flex justify-between items-start">
        <div>
          <h3 class="text-lg font-semibold mb-2">{{ topic.title }}</h3>
          <p class="text-gray-600">{{ topic.summary }}</p>
        </div>
        <el-tag type="success" size="small">
          {{ topic.views }} 阅读
        </el-tag>
      </div>
    </div>
  </div>
</template>

<script setup>
defineEmits(['topic-click'])
defineProps({
  topics: {
    type: Array,
    required: true
  }
})
</script> 