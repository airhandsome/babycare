<template>
  <div class="space-y-4">
    <div v-for="expert in experts" 
         :key="expert.id"
         class="bg-white rounded-xl shadow-md p-4 cursor-pointer"
         :class="{ 'border-2 border-blue-500': modelValue?.id === expert.id }"
         @click="$emit('update:modelValue', expert)">
      <div class="flex items-center space-x-4">
        <el-avatar :src="expert.avatar" :size="50" />
        <div>
          <h3 class="font-bold">{{ expert.name }}</h3>
          <p class="text-gray-600 text-sm">{{ expert.title }}</p>
          <div class="flex items-center mt-1">
            <el-rate v-model="expert.rating" disabled :max="5" />
            <span class="ml-2 text-sm text-gray-500">
              已服务 {{ expert.consultationCount }} 人
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  experts: {
    type: Array,
    required: true
  },
  modelValue: {
    type: Object,
    default: null
  }
})

defineEmits(['update:modelValue'])
</script>