<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">幼儿教育</h1>
    
    <!-- 学前教育 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">学前教育指南</h2>
      <div class="grid md:grid-cols-3 gap-6">
        <EducationCard v-for="item in educationGuides" 
                      :key="item.id"
                      :title="item.title"
                      :description="item.description"
                      :image="item.image"
                      :tags="item.tags" />
      </div>
    </section>

    <!-- 社交技能培养 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">社交技能培养</h2>
      <div class="bg-white rounded-xl shadow-md p-6">
        <SkillsList :skills="socialSkills" />
      </div>
    </section>

    <!-- 艺术与创造力 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">艺术与创造力</h2>
      <div class="grid md:grid-cols-2 gap-8">
        <ActivityCard v-for="activity in artActivities"
                     :key="activity.id"
                     :title="activity.title"
                     :description="activity.description"
                     :image="activity.image"
                     :duration="activity.duration"
                     :materials="activity.materials" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import EducationCard from '../components/cards/EducationCard.vue'
import SkillsList from '../components/lists/SkillsList.vue'
import ActivityCard from '../components/cards/ActivityCard.vue'

const educationGuides = ref([
  {
    id: 1,
    title: '选择合适的幼儿园',
    description: '如何根据孩子特点选择合适的幼儿园环境',
    image: 'https://images.unsplash.com/photo-1503454537195-1dcabb73ffb9',
    tags: ['教育指南', '环境选择']
  },
  {
    id: 2,
    title: '早教课程推荐',
    description: '适合3-6岁儿童的早教课程介绍',
    image: 'https://images.unsplash.com/photo-1503676260728-1c00da094a0b',
    tags: ['早教', '课程']
  },
  {
    id: 3,
    title: '入园准备',
    description: '帮助孩子顺利适应幼儿园生活',
    image: 'https://images.unsplash.com/photo-1544507888-56d73eb6046e',
    tags: ['入园适应', '生活技能']
  }
])

const socialSkills = ref([
  {
    id: 1,
    title: '分享与合作',
    description: '培养孩子与他人分享和合作的能力',
    icon: 'Share'
  },
  {
    id: 2,
    title: '情绪管理',
    description: '帮助孩子认识和管理自己的情绪',
    icon: 'Sunny'
  },
  {
    id: 3,
    title: '沟通表达',
    description: '提升语言表达和倾听能力',
    icon: 'ChatDotRound'
  }
])

const artActivities = ref([
  {
    id: 1,
    title: '创意绘画工作坊',
    description: '通过有趣的绘画活动激发孩子的想象力',
    image: 'https://images.unsplash.com/photo-1513364776144-60967b0f800f',
    duration: '30分钟',
    materials: ['画纸', '蜡笔', '水彩']
  },
  {
    id: 2,
    title: '音乐律动课',
    description: '结合音乐和肢体活动，培养艺术感知',
    image: 'https://images.unsplash.com/photo-1445633629932-0029acc44e88',
    duration: '45分钟',
    materials: ['乐器', '音乐播放器']
  }
])
</script>