<template>
  <div class="space-y-6">
    <!-- 记录表单 -->
    <div class="bg-white rounded-lg shadow p-6">
      <h3 class="text-lg font-semibold mb-4">添加记录</h3>
      <el-form :model="form" label-width="80px">
        <el-form-item label="记录类型">
          <el-select v-model="form.type" placeholder="选择记录类型">
            <el-option label="身高" value="身高" />
            <el-option label="体重" value="体重" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="数值">
          <el-input-number 
            v-model="form.value" 
            :precision="1" 
            :step="0.1"
            :min="0"
          />
          <span class="ml-2">{{ form.type === '身高' ? 'cm' : 'kg' }}</span>
        </el-form-item>

        <el-form-item label="日期">
          <el-date-picker
            v-model="form.date"
            type="date"
            placeholder="选择日期"
          />
        </el-form-item>

        <el-form-item label="备注">
          <el-input v-model="form.note" type="textarea" :rows="2" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="handleSubmit">记录</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 成长曲线 -->
    <div class="grid md:grid-cols-2 gap-6">
      <GrowthChart
        title="身高曲线"
        :records="heightRecords"
        unit="cm"
      />
      <GrowthChart
        title="体重曲线"
        :records="weightRecords"
        unit="kg"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import GrowthChart from './GrowthChart.vue'

const props = defineProps({
  records: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['record-added'])

const form = ref({
  type: '身高',
  value: 0,
  date: new Date(),
  note: ''
})

// 分类记录
const heightRecords = computed(() => 
  props.records.filter(r => r.type === '身高')
)

const weightRecords = computed(() => 
  props.records.filter(r => r.type === '体重')
)

const handleSubmit = () => {
  const record = {
    ...form.value,
    unit: form.value.type === '身高' ? 'cm' : 'kg'
  }
  emit('record-added', record)
  
  // 重置表单
  form.value = {
    type: '身高',
    value: 0,
    date: new Date(),
    note: ''
  }
}
</script>