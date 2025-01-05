<template>
  <div class="about-page">
    <div class="max-w-4xl mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold mb-8 text-center">关于我们</h1>
      
      <!-- 公司简介 -->
      <div class="bg-white rounded-lg shadow-md p-6 mb-8">
        <h2 class="text-xl font-semibold mb-4 text-indigo-700">公司简介</h2>
        <p class="text-gray-700 leading-relaxed mb-4">
          BabyCare是一个专注于母婴健康的综合服务平台。我们致力于为准妈妈和新手父母提供专业、可靠的育儿知识和建议。
        </p>
        <p class="text-gray-700 leading-relaxed">
          我们的团队由经验丰富的儿科医生、营养师和育儿专家组成，为用户提供全方位的母婴护理指导和建议。
        </p>
      </div>

      <!-- 我们的使命 -->
      <div class="bg-white rounded-lg shadow-md p-6 mb-8">
        <h2 class="text-xl font-semibold mb-4 text-indigo-700">我们的使命</h2>
        <div class="grid md:grid-cols-3 gap-6">
          <div class="text-center p-4">
            <i class="fas fa-heart text-3xl text-indigo-600 mb-3"></i>
            <h3 class="font-semibold mb-2">关爱母婴</h3>
            <p class="text-gray-600">为母婴提供全面的健康关怀</p>
          </div>
          <div class="text-center p-4">
            <i class="fas fa-book-reader text-3xl text-indigo-600 mb-3"></i>
            <h3 class="font-semibold mb-2">知识传播</h3>
            <p class="text-gray-600">传播科学的育儿知识</p>
          </div>
          <div class="text-center p-4">
            <i class="fas fa-users text-3xl text-indigo-600 mb-3"></i>
            <h3 class="font-semibold mb-2">社区互助</h3>
            <p class="text-gray-600">建立温暖的育儿社区</p>
          </div>
        </div>
      </div>

      <!-- 我们的团队 -->
      <div class="bg-white rounded-lg shadow-md p-6">
        <h2 class="text-xl font-semibold mb-6 text-indigo-700">专家团队</h2>
        <div class="grid md:grid-cols-3 gap-6">
          <div v-for="expert in experts" :key="expert.id" class="text-center">
            <img :src="expert.avatar" :alt="expert.name" class="w-24 h-24 rounded-full mx-auto mb-3">
            <h3 class="font-semibold">{{ expert.name }}</h3>
            <p class="text-gray-600 text-sm">{{ expert.title }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'About',
  data() {
    return {
      experts: [
        {
          id: 1,
          name: '张医生',
          title: '主任医师',
          avatar: '/images/avatars/expert1.jpg'
        },
        {
          id: 2,
          name: '李医生',
          title: '营养学专家',
          avatar: '/images/avatars/expert2.jpg'
        },
        {
          id: 3,
          name: '王医生',
          title: '儿科专家',
          avatar: '/images/avatars/expert3.jpg'
        }
      ]
    }
  }
}
</script>

<style scoped>
.about-page {
  background-color: #f3f4f6;
  min-height: 100vh;
}
</style> 