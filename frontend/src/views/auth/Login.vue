<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8 bg-white p-8 rounded-xl shadow-md">
      <div class="text-center">
        <h2 class="text-3xl font-bold text-gray-900">登录</h2>
        <p class="mt-2 text-sm text-gray-600">
          还没有账号？
          <router-link to="/auth/register" class="text-blue-600 hover:text-blue-800">
            立即注册
          </router-link>
        </p>
      </div>

      <el-form @submit.prevent="handleLogin">
        <el-form-item>
          <el-input
            v-model="form.email"
            type="email"
            placeholder="邮箱"
            :prefix-icon="Message"
          />
        </el-form-item>

        <el-form-item>
          <el-input
            v-model="form.password"
            type="password"
            placeholder="密码"
            :prefix-icon="Lock"
            show-password
          />
        </el-form-item>

        <div class="flex items-center justify-between mb-6">
          <el-checkbox v-model="form.remember">记住我</el-checkbox>
          <router-link to="/auth/forgot-password" class="text-sm text-blue-600 hover:text-blue-800">
            忘记密码？
          </router-link>
        </div>

        <el-button type="primary" class="w-full" @click="handleLogin">
          登录
        </el-button>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Message, Lock } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()

const form = ref({
  email: '',
  password: '',
  remember: false
})

const handleLogin = async () => {
  try {
    // TODO: Implement login logic
    ElMessage.success('登录成功')
    router.push('/')
  } catch (error) {
    ElMessage.error('登录失败')
  }
}
</script>