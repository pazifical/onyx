<script setup lang="ts">
import { DirectoryContentRepository } from '@/repository/directory'
import type { DirectoryContent } from '@/types'
import { onMounted, ref, watch, type Ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const currentDirectory: Ref<string> = ref('/')

const directoryContentRepository = new DirectoryContentRepository()

const directoryContent: Ref<DirectoryContent | null> = ref(null)

watch(
  () => route.params.path,
  async (newPath, oldPath) => {
    console.log(oldPath, '->', newPath)
    updateFromRoutePath(newPath)
  },
)

async function updateFromRoutePath(path: Array<string> | string) {
  if (path && typeof path != 'string') {
    currentDirectory.value = '/' + path.join('/') + '/'
  } else {
    currentDirectory.value = '/'
  }

  directoryContent.value = await directoryContentRepository.getByPath(currentDirectory.value)
}

onMounted(async () => {
  updateFromRoutePath(route.params.path)
  return
})
</script>

<template>
  <main v-if="directoryContent">
    <section>
      <nav>
        {{ currentDirectory }}
      </nav>
    </section>

    <section v-if="directoryContent.directories.length > 0">
      <h2>Directories</h2>
      <div v-for="directory in directoryContent.directories" :key="directory">
        <RouterLink :to="`${currentDirectory}${directory}`">{{ directory }}</RouterLink>
      </div>
    </section>

    <section v-if="directoryContent.files.length > 0">
      <h2>Files</h2>
      <div v-for="filename in directoryContent.files" :key="filename">
        <RouterLink :to="`/note${currentDirectory}${filename}`">{{ filename }}</RouterLink>
      </div>
    </section>
  </main>
</template>

<style scoped>
nav {
  font-size: 1rem;
}

section {
  margin-bottom: 2rem;
}

h2 {
  color: var(--color-text);
}
</style>
