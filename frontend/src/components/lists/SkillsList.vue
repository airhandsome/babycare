<template>
  <div class="grid md:grid-cols-3 gap-6">
    <div v-for="skill in skills" 
         :key="skill.id"
         class="flex items-start space-x-4">
      <el-icon class="text-3xl text-blue-600">
        <component :is="skill.icon" />
      </el-icon>
      <div>
        <h3 class="text-lg font-bold mb-2">{{ skill.title }}</h3>
        <p class="text-gray-600">{{ skill.description }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  skills: Array
})
</script>