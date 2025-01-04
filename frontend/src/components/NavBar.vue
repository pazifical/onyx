<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/" style="border-bottom: 2px solid var(--color-highlight)">
          <div style="display: flex; align-items: center; gap:0.5rem">
            <div>
              <img alt="Onyx logo" class="logo" src="../assets/logo.svg" style="height:3em" />
            </div>
            <div>
              <h1>ONYX</h1>
            </div>
          </div>
        </RouterLink>
        <div>
          <div id="links">
            <RouterLink class="link" to="/">Home</RouterLink>
            <RouterLink class="link" to="/onyx/reminders">Reminders</RouterLink>
          </div>
        </div>
      </nav>
    </div>
  </header>
</template>

<style scoped>

h1 {
  font-size: 2rem;
  font-weight: bold;
  margin: 0px;
  color: var(--color-light);
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

.link {
  border: 2px solid var(--color-dark);
}
.link:hover {
  border-bottom: 2px solid var(--color-light);
}
</style>
