<template>
  <div class="grid md:grid-cols-2 gap-6">
    <div v-for="stage in milestones" :key="stage.age" 
         class="bg-white rounded-xl shadow-md p-6">
      <h3 class="text-xl font-bold mb-4">{{ stage.age }}</h3>
      <ul class="space-y-3">
        <li v-for="item in stage.items" 
            :key="item.title"
            class="flex items-center">
          <el-checkbox v-model="item.completed" disabled>
            {{ item.title }}
          </el-checkbox>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
defineProps({
  milestones: {
    type: Array,
    required: true
  }
})
</script>