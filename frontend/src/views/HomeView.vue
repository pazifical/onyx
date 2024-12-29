<script setup lang="ts">
import { DirectoryContentRepository } from '@/repository/directory'
import { NoteRepository } from '@/repository/note'
import type { DirectoryContent, Note } from '@/types'
import { onMounted, ref, type Ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const currentDirectory: Ref<string> = ref('/')

const noteRepository = new NoteRepository()
const directoryContentRepository = new DirectoryContentRepository()

const notes: Ref<Array<Note>> = ref([])
const directoryContent: Ref<DirectoryContent | null> = ref(null)

onMounted(async () => {
  notes.value = await noteRepository.getAll()

  if (route.params.path) {
    currentDirectory.value = '/' + route.params.path.join('/') + '/'
  }
  directoryContent.value = await directoryContentRepository.getByPath(currentDirectory.value)
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
        <a :href="`${currentDirectory}${directory}`">{{ directory }}</a>
      </div>
    </section>

    <section v-if="directoryContent.files.length > 0">
      <h2>Files</h2>
      <div v-for="filename in directoryContent.files" :key="filename">
        <a :href="`/note${currentDirectory}${filename}`">{{ filename }}</a>
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
</style>
