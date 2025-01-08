<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">婴幼儿成长</h1>
    
    <!-- 成长里程碑 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">成长里程碑</h2>
      <MilestoneTracker :milestones="milestones" />
    </section>

    <!-- 成长记录 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">成长记录</h2>
      <GrowthRecorder />
    </section>

    <!-- 发展指南 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">发展指南</h2>
      <DevelopmentGuides :guides="developmentGuides" />
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import MilestoneTracker from '../components/growth/MilestoneTracker.vue'
import GrowthRecorder from '../components/growth/GrowthRecorder.vue'
import DevelopmentGuides from '../components/growth/DevelopmentGuides.vue'

const milestones = ref([
  {
    age: '0-3个月',
    items: [
      { title: '社交微笑', completed: true },
      { title: '抬头', completed: true },
      { title: '追视物体', completed: false }
    ]
  },
  {
    age: '4-6个月',
    items: [
      { title: '翻身', completed: false },
      { title: '坐立', completed: false },
      { title: '发出声音', completed: true }
    ]
  }
])

const developmentGuides = ref([
  {
    id: 1,
    title: '语言发展',
    icon: 'ChatRound',
    content: '通过日常交谈和阅读培养语言能力',
    image: 'https://images.unsplash.com/photo-1596526131083-e8c633c948d2'
  },
  {
    id: 2,
    title: '运动技能',
    icon: 'Position',
    content: '鼓励爬行和行走，增强体能',
    image: 'https://images.unsplash.com/photo-1590650153855-d9e808231d41'
  },
  {
    id: 3,
    title: '认知能力',
    icon: 'Light',
    content: '通过游戏和互动提升认知水平',
    image: 'https://images.unsplash.com/photo-1596815064285-45ed8a9c0463'
  }
])
</script>