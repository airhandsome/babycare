<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">幼儿教育</h1>
    
    <!-- 学前教育 -->
    <section class="mb-12">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">学前教育指南</h2>
        <el-button type="primary" plain @click="loadEducationGuides">
          <el-icon><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
      <div class="grid md:grid-cols-3 gap-6">
        <EducationCard v-for="item in educationGuides" 
                      :key="item.id"
                      :title="item.title"
                      :description="item.summary"
                      :image="item.image"
                      :tags="item.tags?.split(',')"
                      class="cursor-pointer hover:shadow-lg transition-shadow"
                      @click="handleGuideClick(item)" />
      </div>
    </section>

    <!-- 社交技能培养 -->
    <section class="mb-12">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">社交技能培养</h2>
        <el-button type="primary" plain @click="loadSocialSkills">
          <el-icon><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
      <div class="bg-white rounded-xl shadow-md p-6">
        <SkillsList :skills="socialSkills.map(skill => ({
          id: skill.id,
          title: skill.title,
          description: skill.summary,
          icon: skill.icon
        }))"
        @skill-click="handleSkillClick" />
      </div>
    </section>

    <!-- 艺术与创造力 -->
    <section class="mb-12">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">艺术与创造力</h2>
        <el-button type="primary" plain @click="loadArtActivities">
          <el-icon><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
      <div class="grid md:grid-cols-2 gap-8">
        <ActivityCard v-for="activity in artActivities"
                     :key="activity.id"
                     :title="activity.title"
                     :description="activity.summary"
                     :image="activity.image"
                     :duration="activity.duration"
                     :materials="activity.materials?.split(',')"
                     class="cursor-pointer hover:shadow-lg transition-shadow"
                     @click="handleActivityClick(activity)" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import request from '../utils/request'
import EducationCard from '../components/cards/EducationCard.vue'
import SkillsList from '../components/lists/SkillsList.vue'
import ActivityCard from '../components/cards/ActivityCard.vue'

const router = useRouter()
const educationGuides = ref([])
const socialSkills = ref([])
const artActivities = ref([])

// 加载教育指南
const loadEducationGuides = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'education_guide' }
    })
    educationGuides.value = response
  } catch (error) {
    console.error('Failed to load education guides:', error)
    ElMessage.error('加载教育指南失败')
  }
}

// 加载社交技能
const loadSocialSkills = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'social_skills' }
    })
    socialSkills.value = response
  } catch (error) {
    console.error('Failed to load social skills:', error)
    ElMessage.error('加载社交技能失败')
  }
}

// 加载艺术活动
const loadArtActivities = async () => {
  try {
    const response = await request.get('posts', {
      params: { category: 'art_activities' }
    })
    artActivities.value = response
  } catch (error) {
    console.error('Failed to load art activities:', error)
    ElMessage.error('加载艺术活动失败')
  }
}

// 处理指南点击
const handleGuideClick = (guide) => {
  router.push(`/posts/${guide.id}`)
}

// 处理技能点击
const handleSkillClick = (skill) => {
  router.push(`/posts/${skill.id}`)
}

// 处理活动点击
const handleActivityClick = (activity) => {
  router.push(`/posts/${activity.id}`)
}

onMounted(() => {
  loadEducationGuides()
  loadSocialSkills()
  loadArtActivities()
})
</script>