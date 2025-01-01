<script setup lang="ts">
import type { Reminder } from '@/types';

const props = defineProps<{ reminder: Reminder }>()
</script>

<template>
  <section>
    <header>
      {{ new Date(Date.parse(props.reminder.date)).toUTCString() }}
    </header>
    <main>
      {{ props.reminder.todo }}
    </main>
    <footer>
      {{ props.reminder.source }}
    </footer>
  </section>
</template>

<style scoped>
section {
  border: 1px solid var(--color-highlight);
  max-width: 20rem;
  border-radius: 1rem;
}

header, footer, main {
  padding: 0.2rem 1rem;
}

main {
  font-size: 1.2rem;
}

footer {
  color: rgb(255 255 255 / 0.5);
  border-top: 1px solid var(--color-highlight);
}

header {
  background-color: var(--color-highlight);
  font-weight: bold;
  color: var(--color-dark);
  border-radius: 1rem 1rem 0 0;
}
</style>
