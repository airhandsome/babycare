<template>
  <div class="app-container min-h-screen flex flex-col">
    <TheHeader />
    <main class="flex-grow">
      <router-view></router-view>
    </main>
    <TheFooter />
  </div>
</template>

<script setup>
import TheHeader from './components/layout/TheHeader.vue'
import TheFooter from './components/layout/TheFooter.vue'
</script>