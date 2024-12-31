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
        <strong>Markdown Manager</strong>
      </nav>
    </div>
  </header>
</template>

<style scoped>
h1 {
  font-size: 3rem;
  font-weight: bold;
  margin: 0px;
  text-shadow: 2px 2px 0 var(--color-light);
  letter-spacing: 0.4rem;
}

a {
  line-height: initial;
}

</style>
