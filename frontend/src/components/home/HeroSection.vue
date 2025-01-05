<template>
  <div class="relative bg-gradient-to-r from-indigo-500 to-purple-600 text-white py-20" style="background-image: url('/images/home/index.png'); background-size: 100% 100%; ">
    <div class="max-w-7xl mx-auto px-4">
      <div class="md:w-2/3">
        <h1 class="text-4xl md:text-5xl font-bold mb-6">
          科学育儿，快乐成长
        </h1>
        <p class="text-xl mb-8">
          提供专业的母婴护理知识和建议，陪伴您的育儿之旅
        </p>
        <div class="space-x-4">
          <el-button 
            type="primary" 
            size="large" 
            @click="router.push('/pregnancy-care')"
          >
            孕期指导
          </el-button>
          <el-button 
            type="primary" 
            size="large" 
            @click="router.push('/newborn-care')"
          >
            新生儿护理
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()
</script>