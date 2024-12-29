<script setup lang="ts">
import { NoteRepository } from '@/repository/note'
import type { Note } from '@/types'
import { onMounted, ref, type Ref } from 'vue'

import { useRoute } from 'vue-router'

import * as showdown from 'showdown'

const converter = new showdown.Converter({
  tasklists: true,
})

const baseURL: string = `${window.location.origin}/note`

const route = useRoute()

const noteRepository = new NoteRepository()

const note: Ref<Note | null> = ref(null)

const noteTextMd: Ref<string> = ref('')

const isEditMode: Ref<boolean> = ref(false)

onMounted(async () => {
  await fetchData()
})

async function fetchData() {
  isEditMode.value = false
  const filePath = route.params.path.join('/')
  note.value = await noteRepository.getByPath(filePath)
  updateMarkdown()
}

function toggleEditMode() {
  isEditMode.value = !isEditMode.value
}

function updateMarkdown() {
  noteTextMd.value = converter.makeHtml(note.value?.text.replace('<BASEURL>', baseURL))
}

async function saveNote() {}
</script>

<template>
  <main v-if="note">
    <header>
      <h2 class="title">Displaying {{ note?.path }}</h2>
      <template v-if="isEditMode">
        <div class="button-row">
          <button @click="saveNote()">Save</button>
          <button @click="fetchData()">Cancel</button>
        </div>
      </template>
      <template v-else>
        <button @click="toggleEditMode()">Edit</button>
      </template>
    </header>

    <div class="markdown">
      <template v-if="isEditMode">
        <textarea v-model="note.text"></textarea>
      </template>
      <template v-else>
        <div v-html="noteTextMd"></div>
      </template>
    </div>
  </main>
</template>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  border-bottom: 2px solid var(--color-light);
  padding: 0.5rem 0;
  margin-bottom: 1rem;
}

textarea {
  background-color: var(--color-background);
  width: 100%;
  height: 100vh;
  color: var(--color-text);
}

.title {
  color: var(--color-light);
  font-weight: bold;
}
</style>
