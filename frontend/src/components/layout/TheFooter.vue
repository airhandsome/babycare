<template>
  <footer class="bg-gray-100 mt-12">
    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
        <div>
          <h3 class="text-lg font-semibold mb-4">关于我们</h3>
          <p class="text-gray-600">BabyCare致力于为每个家庭提供专业的育儿指导和支持。</p>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">快速链接</h3>
          <ul class="space-y-2">
            <li><router-link to="/about" class="text-gray-600 hover:text-blue-600">关于我们</router-link></li>
            <li><router-link to="/contact" class="text-gray-600 hover:text-blue-600">联系我们</router-link></li>
            <li><router-link to="/privacy" class="text-gray-600 hover:text-blue-600">隐私政策</router-link></li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">联系方式</h3>
          <ul class="space-y-2 text-gray-600">
            <li>邮箱：contact@babycare.com</li>
            <li>电话：400-123-4567</li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">关注我们</h3>
          <div class="flex space-x-4">
            <a href="#" class="text-gray-600 hover:text-blue-600"><el-icon><Position /></el-icon></a>
            <a href="#" class="text-gray-600 hover:text-blue-600"><el-icon><ChatDotRound /></el-icon></a>
          </div>
        </div>
      </div>
      <div class="text-center text-gray-600 mt-8">
        © 2024 BabyCare. All rights reserved.
      </div>
    </div>
  </footer>
</template>

<script setup>
import { Position, ChatDotRound } from '@element-plus/icons-vue'
</script>