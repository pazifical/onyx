<script setup lang="ts">
import { NoteRepository } from '@/repository/note'
import type { Note } from '@/types'
import { onMounted, ref, watch, type Ref } from 'vue'
import * as showdown from 'showdown'

const props = defineProps<{
  path?: string
}>()

const converter = new showdown.Converter({
  tasklists: true,
})

const noteRepository = new NoteRepository()

const note: Ref<Note | null> = ref(null)

const noteTextMd: Ref<string> = ref('')

const isEditMode: Ref<boolean> = ref(false)

watch(
  () => props.path,
  async (newPath, oldPath) => {
    console.log('note', oldPath, '->', newPath)

    if (newPath) {
      await fetchData()
    } else {
      note.value = null
    }
  },
)

onMounted(async () => {
  await fetchData()
})

async function fetchData() {
  if (!props.path) {
    return
  }

  isEditMode.value = false
  note.value = await noteRepository.getByPath(props.path)
  updateMarkdown()
}

function toggleEditMode() {
  isEditMode.value = !isEditMode.value
}

function updateMarkdown() {
  if (!note.value) {
    return
  }
  noteTextMd.value = converter.makeHtml(note.value?.text)
}

async function saveNote() {
  // TODO
}
</script>

<template>
  <main v-if="note">
    <header>
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

    <h2 class="title">/{{ note?.path }}</h2>
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
/* header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1rem;
} */
header {
  height: 2rem;
  display: flex;
  justify-content: right;
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
