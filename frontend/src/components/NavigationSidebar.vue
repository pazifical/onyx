<script setup lang="ts">
import type { DirectoryContent } from '@/types'

const props = defineProps<{
  currentDirectory: string
  directoryContent: DirectoryContent
}>()
</script>

<template>
  <main v-if="props.directoryContent">
    <section v-if="props.directoryContent.directories.length > 0">
      <h2>Directories</h2>
      <ul>
        <li v-for="directory in props.directoryContent.directories" :key="directory">
          <RouterLink class="btn-secondary" style="border: none;" :to="`${props.currentDirectory}${directory}`">{{ directory }}</RouterLink>
        </li>
      </ul>
    </section>

    <section v-if="props.directoryContent.files.length > 0">
      <h2>Files</h2>
      <ul>
        <li v-for="filename in props.directoryContent.files" :key="filename">
          <button class="btn-secondary" @click="$emit('file-select', `${props.currentDirectory}/${filename}`)">
            {{ filename }}
          </button>
        </li>
      </ul>
    </section>
  </main>
</template>

<style scoped>
nav {
  font-size: 1rem;
}

header > button {
  color: var(--color-light);
  font-weight: bold;
  border: 1px solid var(--color-light);
  width: 1rem;
}

button {
  border: none;
  text-align: left;
  padding: 0;
  border-radius: 0;
  font-weight: normal;
  display: inline-block;
  white-space: nowrap;
}

section {
  margin-bottom: 2rem;
}

ul {
  list-style-position: inside;
  padding-left: 0;
  list-style-type: none;
}

h2 {
  color: var(--color-text);
}

.header {
  border-bottom: 2px solid var(--color-light);
  padding-bottom: 1rem;
}

main {
  /* padding: 1rem 0; */
  overflow: auto;
}
</style>
