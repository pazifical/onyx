<script setup lang="ts">
import { DirectoryContentRepository } from '@/repository/directory';
import { NoteRepository } from '@/repository/note';
import type { DirectoryContent } from '@/types'
import { ref, type Ref } from 'vue';

const props = defineProps<{
  currentDirectory: string
  directoryContent: DirectoryContent
}>()


const emit = defineEmits(["refresh"])

const directoryRepository = new DirectoryContentRepository()
const noteRepository = new NoteRepository()

const isNoteCreation: Ref<boolean> = ref(false)
const newNoteName: Ref<string> = ref("")
const isDirectoryCreation: Ref<boolean> = ref(false)
const newDirectoryName: Ref<string> = ref("")

function startNoteCreation() {
  isNoteCreation.value = true
}

function startDirectoryCreation() {
  isDirectoryCreation.value = true
}

function cancelDirectoryCreation() {
  newDirectoryName.value = ""
  isDirectoryCreation.value = false
}

async function submitDirectoryCreation() {
  const directoryPath = `${props.currentDirectory}/${newDirectoryName.value}`
  console.log("new dir", directoryPath)
  await directoryRepository.createNew(directoryPath)
  newDirectoryName.value = ""
  isDirectoryCreation.value = false
  emit("refresh")
}

async function submitFileCreation() {
  const notePath = `${props.currentDirectory}/${newNoteName.value}`
  console.log("new note", notePath)
  await noteRepository.saveNote({
    path: notePath,
    text: "",
  })
  newNoteName.value = ""
  isNoteCreation.value = false
  emit("refresh")
}

function cancelFileCreation() {
  newNoteName.value = ""
  isNoteCreation.value = false
}
</script>

<template>
  <main v-if="props.directoryContent">
    <section>
      <h2>Directories</h2>
      <ul>
        <li v-for="directory in props.directoryContent.directories" :key="directory">
          <RouterLink class="btn-secondary" style="border: none;" :to="`${props.currentDirectory}/${directory}`">
            <p class="bullet">⧐</p> {{ directory }}
          </RouterLink>
        </li>

        <li>
          <template v-if="isDirectoryCreation">
            <input type="text" v-model="newDirectoryName">
            <div style="display: grid; gap: 0.5rem; grid-template-columns: 1fr 1fr;">
              <button @click="submitDirectoryCreation()" class="btn-primary" style="text-align: center;">✔</button>
              <button @click="cancelDirectoryCreation()" class="btn-primary" style="text-align: center;">✘</button>
            </div>
          </template>
          <template v-else>
            <button @click="startDirectoryCreation()" class="btn-secondary" style="color: var(--color-light);">
              <p class="bullet">⧏</p> Create new
            </button>
          </template>
        </li>
      </ul>
    </section>

    <section>
      <div style="display: flex; justify-content: space-between; align-items: center;">
        <h2>Files</h2>
      </div>
      <ul>
        <li v-for="filename in props.directoryContent.files" :key="filename">
          <!-- <button class="btn-secondary" @click="$emit('file-select', `${props.currentDirectory}/${filename}`)">
            <p class="bullet">⧐</p> {{ filename }}
          </button> -->
          <RouterLink class="btn-secondary" style="border: none;" :to="`${props.currentDirectory}/${filename}`">
            <p class="bullet">⧐</p> {{ filename }}
          </RouterLink>
        </li>
        <li>
          <template v-if="isNoteCreation">
            <input type="text" v-model="newNoteName">
            <div style="display: grid; gap: 0.5rem; grid-template-columns: 1fr 1fr;">
              <button @click="submitFileCreation()" class="btn-primary" style="text-align: center;">✔</button>
              <button @click="cancelFileCreation()" class="btn-primary" style="text-align: center;">✘</button>
            </div>
          </template>
          <template v-else>
            <button @click="startNoteCreation()" class="btn-secondary" style="color: var(--color-light);">
              <p class="bullet">⧏</p> Create new
            </button>
          </template>
        </li>
      </ul>
    </section>
  </main>
</template>

<style scoped>
input[type="text"] {
  background-color: var(--color-black);
  color: var(--color-light);
  border: 1px solid var(--color-highlight);
  width: 100%;
}

.bullet {
  color: var(--color-light);
  display: inline-block;
}

nav {
  font-size: 1rem;
}

header>button {
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

.icon-button {
  background-color: none;
  border: none;
  padding: none;
  background: none;
  color: var(--color-light);
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
