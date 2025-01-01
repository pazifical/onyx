<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <div>
          <RouterLink to="/">
            <h1>O<strong style="font-size: 2rem;">NYX</strong></h1>
          </RouterLink>
          <strong>Markdown Manager</strong>
        </div>
        <div>
          <div id="links">
            <RouterLink to="/">Home</RouterLink>
            <RouterLink to="/onyx/reminders">Reminders</RouterLink>
          </div>
        </div>
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

nav {
  display: flex;
  justify-content: space-between;
  align-items: end;
}

#links {
  display: flex;
  gap: 2rem;
  font-weight: bold;
  font-size: 1.2rem;
}
</style>
