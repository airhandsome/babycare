<template>
  <div class="space-y-4">
    <div v-for="skill in skills" 
         :key="skill.id" 
         class="p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors cursor-pointer"
         @click="$emit('skill-click', skill)">
      <div class="flex items-start space-x-4">
        <el-icon :size="24" class="text-indigo-600">
          <component :is="skill.icon" />
        </el-icon>
        <div>
          <h3 class="text-lg font-semibold mb-2">{{ skill.title }}</h3>
          <p class="text-gray-600">{{ skill.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineEmits(['skill-click'])
defineProps({
  skills: {
    type: Array,
    required: true
  }
})
</script>