<template>
  <div class="bg-white rounded-xl shadow-md p-6">
    <img :src="image" :alt="title" class="w-full h-48 object-cover rounded-lg mb-4">
    <h3 class="text-xl font-bold mb-2">{{ title }}</h3>
    <p class="text-gray-600">{{ description }}</p>
    <el-button type="primary" class="mt-4" @click="$emit('click')">查看详情</el-button>
  </div>
</template>

<script>
export default {
  name: 'CareCard',
  props: {
    title: {
      type: String,
      required: true
    },
    description: {
      type: String,
      required: true
    },
    image: {
      type: String,
      required: true
    },
    icon: {
      type: String,
      default: ''
    }
  }
}
</script>