<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8 bg-white p-8 rounded-xl shadow-md">
      <div class="text-center">
        <h2 class="text-3xl font-bold text-gray-900">找回密码</h2>
        <p class="mt-2 text-sm text-gray-600">
          我们将向您的邮箱发送重置密码的链接
        </p>
      </div>

      <el-form @submit.prevent="handleSubmit">
        <el-form-item>
          <el-input
            v-model="email"
            type="email"
            placeholder="请输入您的邮箱地址"
            :prefix-icon="Message"
          />
        </el-form-item>

        <el-button type="primary" class="w-full" @click="handleSubmit">
          发送重置链接
        </el-button>

        <div class="text-center mt-4">
          <router-link to="/auth/login" class="text-sm text-blue-600 hover:text-blue-800">
            返回登录
          </router-link>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Message } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const email = ref('')

const handleSubmit = async () => {
  if (!email.value) {
    ElMessage.warning('请输入邮箱地址')
    return
  }

  try {
    // TODO: Implement password reset logic
    ElMessage.success('重置链接已发送到您的邮箱')
  } catch (error) {
    ElMessage.error('发送失败，请稍后重试')
  }
}
</script>