<template>
  <div class="mt-6">
    <el-form :model="formData" @submit.prevent="handleSubmit">
      <!-- 咨询类型 -->
      <el-form-item label="咨询类型">
        <el-select v-model="formData.type" class="w-full">
          <el-option
            v-for="type in consultationTypes"
            :key="type.value"
            :label="type.label"
            :value="type.value"
          />
        </el-select>
      </el-form-item>

      <!-- 咨询问题 -->
      <el-form-item label="咨询问题">
        <el-input
          v-model="formData.question"
          type="textarea"
          :rows="4"
          placeholder="请详细描述您的问题..."
        />
      </el-form-item>

      <!-- 图片上传 -->
      <el-form-item label="相关图片">
        <el-upload
          action="#"
          list-type="picture-card"
          :auto-upload="false"
          :on-change="handleImageChange"
        >
          <el-icon><Plus /></el-icon>
        </el-upload>
      </el-form-item>

      <!-- 预约时间 -->
      <el-form-item label="期望咨询时间">
        <el-date-picker
          v-model="formData.preferredTime"
          type="datetime"
          placeholder="选择期望的咨询时间"
          :disabled-date="disabledDate"
          class="w-full"
        />
      </el-form-item>

      <div class="flex justify-between items-center mt-6">
        <div class="text-blue-600 font-bold">
          费用：¥{{ expert.price }}
        </div>
        <el-button type="primary" @click="handleSubmit">
          提交咨询
        </el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Plus } from '@element-plus/icons-vue'

const props = defineProps({
  expert: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['submit'])

const formData = ref({
  type: '',
  question: '',
  images: [],
  preferredTime: null
})

const consultationTypes = [
  { value: 'growth', label: '生长发育' },
  { value: 'nutrition', label: '营养饮食' },
  { value: 'health', label: '疾病防治' },
  { value: 'psychology', label: '心理行为' },
  { value: 'education', label: '教育指导' }
]

const handleImageChange = (file) => {
  formData.value.images.push(file)
}

const disabledDate = (time) => {
  return time.getTime() < Date.now()
}

const handleSubmit = () => {
  emit('submit', {
    ...formData.value,
    expertId: props.expert.id
  })
}
</script>