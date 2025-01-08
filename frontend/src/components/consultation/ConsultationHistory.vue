<template>
  <div class="space-y-4">
    <div v-for="record in records" 
         :key="record.id"
         class="bg-white rounded-xl shadow-md p-6">
      <div class="flex justify-between items-start mb-4">
        <div>
          <h3 class="font-bold text-lg mb-2">{{ record.expert }}</h3>
          <p class="text-gray-600">{{ record.topic }}</p>
        </div>
        <el-tag :type="getStatusType(record.status)">
          {{ record.status }}
        </el-tag>
      </div>
      
      <div class="flex justify-between items-center text-sm text-gray-500">
        <span>咨询时间：{{ record.date }}</span>
        <div class="space-x-2">
          <el-button 
            v-if="record.status === '已完成'"
            type="primary" 
            link
            @click="viewDetail(record.id)"
          >
            查看记录
          </el-button>
          <el-button 
            v-if="record.status === '待咨询'"
            type="danger" 
            link
            @click="cancelConsultation(record.id)"
          >
            取消预约
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  records: {
    type: Array,
    required: true
  }
})

const getStatusType = (status) => {
  const types = {
    '待咨询': 'warning',
    '进行中': 'primary',
    '已完成': 'success',
    '已取消': 'info'
  }
  return types[status] || 'info'
}

const viewDetail = (id) => {
  console.log('View consultation detail:', id)
}

const cancelConsultation = (id) => {
  console.log('Cancel consultation:', id)
}
</script>