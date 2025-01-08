<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">在线咨询</h1>

    <div class="grid md:grid-cols-3 gap-8">
      <!-- 专家列表 -->
      <div class="md:col-span-1">
        <ExpertList 
          :experts="experts"
          v-model:selected="selectedExpert"
        />
      </div>

      <!-- 咨询区域 -->
      <div class="md:col-span-2">
        <div v-if="selectedExpert" class="bg-white rounded-xl shadow-md p-6">
          <ExpertProfile :expert="selectedExpert" />
          <ConsultationForm 
            :expert="selectedExpert"
            @submit="handleConsultation"
          />
        </div>
        <div v-else class="bg-white rounded-xl shadow-md p-6 text-center">
          请选择一位专家进行咨询
        </div>
      </div>
    </div>

    <!-- 咨询记录 -->
    <section class="mt-8">
      <h2 class="text-2xl font-bold mb-6">咨询记录</h2>
      <ConsultationHistory :records="consultationRecords" />
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ExpertList from '../components/consultation/ExpertList.vue'
import ExpertProfile from '../components/consultation/ExpertProfile.vue'
import ConsultationForm from '../components/consultation/ConsultationForm.vue'
import ConsultationHistory from '../components/consultation/ConsultationHistory.vue'

const selectedExpert = ref(null)

const experts = ref([
  {
    id: 1,
    name: '王医生',
    title: '儿科主任医师',
    avatar: 'https://images.unsplash.com/photo-1559839734-2b71ea197ec2',
    specialties: ['儿童营养', '疾病预防'],
    rating: 4.9,
    consultationCount: 1000,
    price: 100
  },
  {
    id: 2,
    name: '李教授',
    title: '儿童心理学专家',
    avatar: 'https://images.unsplash.com/photo-1580489944761-15a19d654956',
    specialties: ['心理发展', '行为矫正'],
    rating: 4.8,
    consultationCount: 800,
    price: 150
  }
])

const consultationRecords = ref([
  {
    id: 1,
    expert: '王医生',
    date: '2024-01-20',
    topic: '宝宝发烧护理',
    status: '已完成'
  }
])

const handleConsultation = (data) => {
  // 处理咨询提交
  console.log('New consultation:', data)
}
</script>