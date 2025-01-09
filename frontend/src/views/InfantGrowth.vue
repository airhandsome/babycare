<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">婴幼儿成长</h1>
    
    <!-- 成长里程碑 -->
    <section class="mb-12">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">成长里程碑</h2>
        <el-button type="primary" plain @click="loadMilestones">
          <el-icon><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
      <MilestoneTracker 
        :milestones="milestones" 
        @update="handleMilestoneUpdate"
        @create="handleMilestoneCreate"
      />
    </section>

    <!-- 成长记录 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">成长记录</h2>
      <GrowthRecorder 
        @record-added="handleRecordAdded"
        :records="records"
      />
    </section>

    <!-- 发展指南 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">发展指南</h2>
      <DevelopmentGuides :guides="developmentGuides" />
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import request from '../utils/request'
import MilestoneTracker from '../components/growth/MilestoneTracker.vue'
import GrowthRecorder from '../components/growth/GrowthRecorder.vue'
import DevelopmentGuides from '../components/growth/DevelopmentGuides.vue'

const milestones = ref([])
const records = ref([])

const developmentGuides = ref([
  {
    id: 1,
    title: '语言发展',
    icon: 'ChatRound',
    content: '通过日常交谈和阅读培养语言能力',
    image: 'https://images.unsplash.com/photo-1596526131083-e8c633c948d2'
  },
  // ... 其他指南
])

// 加载里程碑数据
const loadMilestones = async () => {
  try {
    const response = await request.get('growth/milestones')
    milestones.value = response
  } catch (error) {
    console.error('Failed to load milestones:', error)
    ElMessage.error('加载里程碑失败')
  }
}

// 加载成长记录
const loadRecords = async () => {
  try {
    const response = await request.get('growth/records')
    records.value = response
  } catch (error) {
    console.error('Failed to load records:', error)
    ElMessage.error('加载成长记录失败')
  }
}

// 处理里程碑更新
const handleMilestoneUpdate = async (milestone) => {
  try {
    await request.put(`growth/milestones/${milestone.id}`, milestone)
    ElMessage.success('更新成功')
    await loadMilestones()
  } catch (error) {
    console.error('Failed to update milestone:', error)
    ElMessage.error('更新失败')
  }
}

// 处理里程碑创建
const handleMilestoneCreate = async (milestone) => {
  try {
    await request.post('growth/milestones', milestone)
    ElMessage.success('创建成功')
    await loadMilestones()
  } catch (error) {
    console.error('Failed to create milestone:', error)
    ElMessage.error('创建失败')
  }
}

// 处理新增记录
const handleRecordAdded = async (record) => {
  try {
    await request.post('growth/records', record)
    ElMessage.success('记录添加成功')
    await loadRecords()
  } catch (error) {
    console.error('Failed to add record:', error)
    ElMessage.error('添加记录失败')
  }
}

onMounted(() => {
  loadMilestones()
  loadRecords()
})
</script>