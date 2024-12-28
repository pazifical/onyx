<script setup lang="ts">
import { NoteRepository } from '@/repository/note'
import type { Note } from '@/types'
import { onMounted, ref, type Ref } from 'vue'

const noteRepository = new NoteRepository()

const notes: Ref<Array<Note>> = ref([])

onMounted(async () => {
  notes.value = await noteRepository.getAll()
})
</script>

<template>
  <main>
    <div v-for="note in notes" :key="note.id">
      <a :href="`/note/${note.path}`">{{ note.path }}</a>
    </div>
  </main>
</template>
