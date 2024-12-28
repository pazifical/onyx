<script setup lang="ts">
import { NoteRepository } from '@/repository/note'
import type { Note } from '@/types'
import { onMounted, ref, type Ref } from 'vue'

const noteRepository = new NoteRepository()

const notes: Ref<Array<Note>> = ref([])

onMounted(async () => {
  notes.value = await noteRepository.getAllNotes()
})
</script>

<template>
  <main>{{ notes }}</main>
</template>
