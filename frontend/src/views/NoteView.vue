<script setup lang="ts">
import { NoteRepository } from '@/repository/note'
import type { Note } from '@/types'
import { onMounted, ref, type Ref } from 'vue'

import { useRoute } from 'vue-router'

const route = useRoute()

const noteRepository = new NoteRepository()

const note: Ref<Note | null> = ref(null)

const noteText: Ref<string> = ref('')

onMounted(async () => {
  const filePath = route.params.path.join('/')
  note.value = await noteRepository.getByPath(filePath)

  const pattern = /\[.*\]/
  const matches = note.value.text.match(pattern)
  console.log(matches)

  noteText.value = `${note.value.text}`
  matches?.forEach((m) => {
    console.log(m)
    const replacement = m.replace('[[', '').replace(']]', '')
    const parts = replacement.split('/')
    const fileName = parts[parts.length - 1].replace('.md', '')
    console.log(fileName)
    console.log(replacement)
    noteText.value = noteText.value.replace(m, `<a href="/note/${replacement}">${fileName}</a>`)
  })
})
</script>

<template>
  <main v-if="note">
    <h2>
      {{ note?.path }}
    </h2>
    <p v-html="noteText"></p>
  </main>
</template>
