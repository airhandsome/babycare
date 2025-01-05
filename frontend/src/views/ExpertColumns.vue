<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">专家专栏</h1>

    <!-- 专家推荐 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">特邀专家</h2>
      <div class="grid md:grid-cols-3 gap-6">
        <ExpertCard v-for="expert in experts"
                   :key="expert.id"
                   :name="expert.name"
                   :title="expert.title"
                   :avatar="expert.avatar"
                   :specialties="expert.specialties"
                   :description="expert.description" />
      </div>
    </section>

    <!-- 最新文章 -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">最新专栏文章</h2>
      <div class="space-y-6">
        <ArticleCard v-for="article in articles"
                    :key="article.id"
                    :title="article.title"
                    :author="article.author"
                    :date="article.date"
                    :summary="article.summary"
                    :image="article.image"
                    :tags="article.tags" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ExpertCard from '../components/cards/ExpertCard.vue'
import ArticleCard from '../components/cards/ArticleCard.vue'

const experts = ref([
  {
    id: 1,
    name: '王医生',
    title: '儿科主任医师',
    avatar: 'https://images.unsplash.com/photo-1559839734-2b71ea197ec2',
    specialties: ['儿童营养', '疾病预防'],
    description: '从事儿科临床工作20年，专注儿童营养与健康研究。'
  },
  {
    id: 2,
    name: '李教授',
    title: '儿童心理学专家',
    avatar: 'https://images.unsplash.com/photo-1580489944761-15a19d654956',
    specialties: ['心理发展', '行为矫正'],
    description: '著名儿童心理学专家，擅长儿童心理健康指导。'
  },
  {
    id: 3,
    name: '张老师',
    title: '幼教专家',
    avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80',
    specialties: ['早期教育', '亲子关系'],
    description: '资深幼教工作者，专注幼儿教育研究与指导。'
  }
])

const articles = ref([
  {
    id: 1,
    title: '婴儿期营养补充的关键点',
    author: '王医生',
    date: '2024-01-15',
    summary: '详细解析0-1岁婴儿的营养需求，科学指导营养补充。',
    image: 'https://images.unsplash.com/photo-1544829099-b9a0c07fad1a',
    tags: ['营养', '婴儿护理']
  },
  {
    id: 2,
    title: '如何培养孩子的专注力',
    author: '李教授',
    date: '2024-01-10',
    summary: '通过游戏和日常活动培养孩子的注意力和专注力。',
    image: 'https://images.unsplash.com/photo-1503676260728-1c00da094a0b',
    tags: ['心理', '教育']
  }
])
</script>