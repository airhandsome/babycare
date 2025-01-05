<template>
  <div class="contact-page">
    <div class="max-w-4xl mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold mb-8 text-center">联系我们</h1>
      
      <div class="grid md:grid-cols-2 gap-8">
        <!-- 联系信息 -->
        <div class="bg-white rounded-lg shadow-md p-6">
          <h2 class="text-xl font-semibold mb-6 text-indigo-700">联系方式</h2>
          
          <div class="space-y-4">
            <div class="flex items-center">
              <i class="fas fa-phone text-indigo-600 w-6"></i>
              <span class="ml-3">400-123-4567</span>
            </div>
            <div class="flex items-center">
              <i class="fas fa-envelope text-indigo-600 w-6"></i>
              <span class="ml-3">support@babycare.com</span>
            </div>
            <div class="flex items-center">
              <i class="fas fa-map-marker-alt text-indigo-600 w-6"></i>
              <span class="ml-3">北京市朝阳区xxx街xxx号</span>
            </div>
            <div class="flex items-center">
              <i class="fas fa-clock text-indigo-600 w-6"></i>
              <span class="ml-3">周一至周五 9:00-18:00</span>
            </div>
          </div>

          <div class="mt-8">
            <h3 class="font-semibold mb-3">关注我们</h3>
            <div class="flex space-x-4">
              <a href="#" class="text-gray-600 hover:text-indigo-600">
                <i class="fab fa-weixin text-2xl"></i>
              </a>
              <a href="#" class="text-gray-600 hover:text-indigo-600">
                <i class="fab fa-weibo text-2xl"></i>
              </a>
            </div>
          </div>
        </div>

        <!-- 联系表单 -->
        <div class="bg-white rounded-lg shadow-md p-6">
          <h2 class="text-xl font-semibold mb-6 text-indigo-700">发送消息</h2>
          
          <form @submit.prevent="submitForm" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">姓名</label>
              <input 
                type="text" 
                v-model="form.name"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                required
              >
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">邮箱</label>
              <input 
                type="email" 
                v-model="form.email"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                required
              >
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">消息内容</label>
              <textarea 
                v-model="form.message"
                rows="4"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                required
              ></textarea>
            </div>
            
            <button 
              type="submit"
              class="w-full bg-indigo-600 text-white py-2 px-4 rounded-md hover:bg-indigo-700 transition-colors"
            >
              发送消息
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Contact',
  data() {
    return {
      form: {
        name: '',
        email: '',
        message: ''
      }
    }
  },
  methods: {
    submitForm() {
      // TODO: 实现表单提交逻辑
      console.log('Form submitted:', this.form)
      alert('消息已发送！我们会尽快回复您。')
      this.form = {
        name: '',
        email: '',
        message: ''
      }
    }
  }
}
</script>

<style scoped>
.contact-page {
  background-color: #f3f4f6;
  min-height: 100vh;
}
</style> 