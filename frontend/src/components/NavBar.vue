<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">
          <h1>O<strong style="font-size: 2rem;">NYX</strong></h1>
        </RouterLink>
      </nav>
    </div>
  </header>
</template>

<style scoped>
h1 {
  font-size: 3rem;
  font-weight: bold;
  margin: 0px;
}
</style>
