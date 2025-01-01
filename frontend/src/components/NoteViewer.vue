<script setup lang="ts">
import { NoteRepository } from '@/repository/note'
import type { Note } from '@/types'
import { onMounted, ref, watch, type Ref } from 'vue'
import * as showdown from 'showdown'
import PencilFillIcon from './icons/PencilFillIcon.vue';
import FloppyFillIcon from './icons/FloppyFillIcon.vue';
import SquareXIcon from './icons/SquareXIcon.vue';
import { useRoute } from 'vue-router';

// const props = defineProps<{ path?: string }>()

const converter = new showdown.Converter({
  tasklists: true,
})

const noteRepository = new NoteRepository()

const note: Ref<Note | null> = ref(null)

const noteTextMd: Ref<string> = ref('')

const isEditMode: Ref<boolean> = ref(false)

const route = useRoute()

const fileName: Ref<string> = ref("")

// watch(() => props.path, async (oldPath, newPath) => {
//   console.log("from", oldPath, "to", newPath)
// })

watch(
  () => route.query.file,
  async (newFilename, oldFilename) => {
    if (newFilename === oldFilename) {
      return
    }

    console.log('note', oldFilename, '->', newFilename)

    if (newFilename) {
      await fetchData(`${route.path}/${newFilename}`)
    } else {
      note.value = null
    }

  },
)

onMounted(async () => {
  const url = `${route.path}/${route.query.file}`;
  console.log("fetching from", url)
  await fetchData(url)
})

async function fetchData(filepath?: string) {
  if (!filepath) {
    return
  }

  const parts = filepath.split("/")
  fileName.value = parts[parts.length - 1]

  isEditMode.value = false
  note.value = await noteRepository.getByPath(filepath)
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

async function cancelEdit() {
  await fetchData(`${route.path}/${route.query.file}`)
}
</script>

<template>

  <header>
    <template v-if="isEditMode">
      <button class="btn-primary" @click="saveNote()">
        <FloppyFillIcon />
        <strong>Save</strong>
      </button>
      <button class="btn-primary" @click="cancelEdit()">
        <SquareXIcon />
        <strong>Cancel</strong>
      </button>
    </template>
    <template v-else>
      <button class="btn-primary" @click="toggleEditMode()">
        <PencilFillIcon />
        <strong>Edit</strong>
      </button>
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

button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  justify-content: center;
}

header {
  display: flex;
  align-content: stretch;
  position: sticky;
  top: 0;
}

header>* {
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
  border-bottom: 1px solid var(--color-light);
}

main {
  padding: 1rem;
}

.markdown {
  font-size: 1.2rem;
}

</style>
