<template>
  <div class="flex flex-wrap gap-2">
    <el-radio-group v-model="selected" @change="handleChange">
      <el-radio-button 
        v-for="category in categories" 
        :key="category.id" 
        :label="category.id"
      >
        {{ category.name }}
      </el-radio-button>
    </el-radio-group>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  categories: {
    type: Array,
    required: true
  },
  modelValue: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['update:modelValue'])

const selected = ref(props.modelValue)

watch(() => props.modelValue, (newValue) => {
  selected.value = newValue
})

const handleChange = (value) => {
  emit('update:modelValue', value)
}
</script>