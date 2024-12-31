<script setup lang="ts">
import { NoteRepository } from '@/repository/note'
import type { Note } from '@/types'
import { computed, onMounted, ref, watch, type ComputedRef, type Ref } from 'vue'
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

const fileName: ComputedRef<string> = computed(() => {
  if (!props.path) {
    return ""
  }
  const parts = props.path.split("/");
  return parts[parts.length - 1]
})

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
  if (!note.value) {
    return
  }
  await noteRepository.saveNote(note.value)
  noteTextMd.value = converter.makeHtml(note.value?.text)
  toggleEditMode()
}
</script>

<template>

  <header>
      <template v-if="isEditMode">
          <button class="btn-primary" @click="saveNote()">Save</button>
          <button class="btn-primary" @click="fetchData()">Cancel</button>
      </template>
      <template v-else>
        <button class="btn-primary" @click="toggleEditMode()">Edit</button>
      </template>
  </header>
  <main v-if="note">
    <div>
      <h2 class="title" style="margin-left: 0; margin-right: auto;">{{ fileName }}</h2>
    </div>
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
  justify-content: space-between;
  margin-bottom: 1rem;
} */

header {
  display: flex;
  align-content: stretch;
}

header > * {
  flex-grow: 1;
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

main {
  padding: 1rem;
}
</style>
