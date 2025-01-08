<template>
  <div class="space-y-6">
    <div v-for="comment in comments" 
         :key="comment.id"
         class="bg-white rounded-lg p-4">
      <div class="flex justify-between mb-2">
        <span class="font-bold">{{ comment.author }}</span>
        <span class="text-gray-500">{{ formatDate(comment.date) }}</span>
      </div>
      <p class="text-gray-700">{{ comment.content }}</p>
      <div class="mt-2">
        <el-button text size="small" @click="$emit('reply', comment.id)">
          回复
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { formatDate } from '../../utils/dateFormat'

defineProps({
  comments: {
    type: Array,
    required: true
  }
})

defineEmits(['reply'])
</script>