<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <div style="display: flex; align-items: end; gap:0.5rem">
          <div>
            <img alt="Onyx logo" class="logo" src="../assets/logo.svg" style="height:3em" />
          </div>
          <div>
            <RouterLink to="/">
              <h1>O<strong style="font-size: 2rem;">NYX</strong></h1>
            </RouterLink>
          </div>
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
.logo {
}

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
