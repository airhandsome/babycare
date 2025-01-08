<template>
  <div class="bg-white rounded-xl shadow-md p-6">
    <div class="grid md:grid-cols-2 gap-6">
      <!-- 身高体重记录 -->
      <div>
        <h3 class="text-lg font-bold mb-4">身高体重记录</h3>
        <el-form>
          <el-form-item label="日期">
            <el-date-picker v-model="date" type="date" placeholder="选择日期" />
          </el-form-item>
          <el-form-item label="身高(cm)">
            <el-input-number v-model="height" :min="0" :max="200" />
          </el-form-item>
          <el-form-item label="体重(kg)">
            <el-input-number v-model="weight" :min="0" :max="100" :precision="2" />
          </el-form-item>
          <el-button type="primary" @click="saveRecord">记录</el-button>
        </el-form>
      </div>

      <!-- 生长曲线 -->
      <div>
        <h3 class="text-lg font-bold mb-4">生长曲线</h3>
        <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
          生长曲线图表
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const date = ref('')
const height = ref(0)
const weight = ref(0)

const saveRecord = () => {
  // TODO: Implement save logic
  console.log('Saving record:', { date: date.value, height: height.value, weight: weight.value })
}
</script>