<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavBar from '@/components/NavBar.vue'
import AppFooter from './components/AppFooter.vue';
</script>

<template>
  <header>
    <NavBar />
  </header>

  <RouterView />

  <footer>
    <AppFooter />
</footer>
</template>

<style scoped>

</style>
