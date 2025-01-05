<script setup lang="ts">
import SquareXIcon from '@/components/icons/SquareXIcon.vue'
import NavigationSidebar from '@/components/NavigationSidebar.vue'
import NoteViewer from '@/components/NoteViewer.vue'
import { DirectoryContentRepository } from '@/repository/directory'
import { NoteRepository } from '@/repository/note'
import type { DirectoryContent, Note } from '@/types'
import { computed, onMounted, ref, watch, type ComputedRef, type Ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const currentDirectory: Ref<string> = ref('/')

const errorDialog: Ref<HTMLDialogElement | null> = ref(null)
const errorMessage: Ref<string> = ref("")

const directoryContentRepository = new DirectoryContentRepository()
const noteRepository = new NoteRepository()

const selectedNote: Ref<Note | null> = ref(null)

const directoryContent: Ref<DirectoryContent | null> = ref(null)

const selectedFilePath: Ref<string> = ref('')

const isSidebarVisible: Ref<boolean> = ref(true)
const selectedGridLayout: Ref<string> = ref('20ch 1fr')

const parentDirectories: ComputedRef<Array<Array<string>>> = computed(() => {
  const directories: Array<Array<string>> = []

  const parts = currentDirectory.value.split('/').filter((d) => d !== '')
  console.log(parts)

  let path: string = ''
  parts.forEach((p) => {
    path += `/${p}`
    directories.push([p, `${path}`])
  })

  return directories
})

watch(
  () => route.params.path,
  async (newPath, oldPath) => {
    console.log("oldPath", oldPath, '->', "newPath", newPath)
    updateFromRoutePath()
  },
)

async function updateFromRoutePath() {
  const path = route.params.path
  if (path && typeof path != 'string') {
    const parts = [...path];
    if (parts[parts.length - 1].endsWith(".md")) {
      currentDirectory.value = "/" + parts.slice(0, parts.length - 1).join("/")
      selectedFilePath.value = parts[parts.length - 1]

      try {
        selectedNote.value = await noteRepository.getByPath(path.join("/"))
      } catch (e) {
        console.log(`${e}`)
        if (e instanceof Error) {
          errorMessage.value = e.message
          errorDialog.value?.showModal()
        }
        return
      }
      console.log("updating note with note", selectedNote.value)

      if (parts.length == 1) {
        currentDirectory.value = ""
      }

    } else {
      currentDirectory.value = "/" + path.join('/')
    }
  } else {
    currentDirectory.value = ''
  }

  console.log('currentDirectory', currentDirectory.value)

  try {
    directoryContent.value = await directoryContentRepository.getByPath(currentDirectory.value)
  } catch (e) {
    console.log(`${e}`)
    if (e instanceof Error) {
      errorMessage.value = e.message
      errorDialog.value?.showModal()
    }
    return
  }
}

function closeErrorDialog() {
  errorMessage.value = ""
  errorDialog.value?.close()
}

onMounted(async () => {
  updateFromRoutePath()
})

function hideSidebar() {
  isSidebarVisible.value = false
  selectedGridLayout.value = '1rem 1fr'
}

function showSidebar() {
  isSidebarVisible.value = true
  selectedGridLayout.value = '20ch 1fr'
}
</script>

<template>
  <main v-if="directoryContent">
    <section class="header">
      <nav v-if="parentDirectories.length > 0">
        <RouterLink :to="`${pd[1]}`" v-for="pd in parentDirectories" :key="pd[1]">
          {{ '/' + pd[0] }}
        </RouterLink>
      </nav>
      <nav v-else>
        <RouterLink to="/">/</RouterLink>
      </nav>
    </section>

    <div class="content" :style="{ 'grid-template-columns': selectedGridLayout }">
      <div id="sidebar">
        <div id="nav-area">
          <NavigationSidebar class="sidebar-content" :directory-content="directoryContent"
            :current-directory="currentDirectory"
            @refresh="updateFromRoutePath()"/>

          <template v-if="isSidebarVisible">
            <button class="shrinker" @click="hideSidebar()">◀</button>
          </template>
          <template v-else>
            <button class="shrinker" @click="showSidebar()">▶</button>
          </template>

        </div>
      </div>

      <div id="note-viewer" v-if="selectedFilePath">
        <NoteViewer :note="selectedNote" :key="selectedNote?.path" />
      </div>
      <div v-else style="display: flex; justify-content: center;; padding: 1rem">
        <strong style="font-size: 1.5rem; color: rgb(255 255 255 / 0.5)">Please select a file on the left </strong>
      </div>
    </div>
  </main>

  <dialog ref="errorDialog" id="error-dialog">
    <header>
      <h1>Error</h1>
      <button @click="closeErrorDialog()" class="btn-primary">
        <SquareXIcon />
      </button>
    </header>
    <br>
    <p>
      {{ errorMessage }}
    </p>
  </dialog>
</template>

<style scoped>
#error-dialog {
  background-color: var(--color-dark);
  border: 4px solid var(--color-highlight);
  margin: auto;
}

#error-dialog::backdrop {
  background-color: rgb(0 0 0 /0.5);
  backdrop-filter: blur(4px);
}


#error-dialog>p {
  color: var(--color-light);
  font-size: 1.2rem;
}

#error-dialog>header {
  display: flex;
  justify-content: space-between;
}

#error-dialog>header>button {
  border: none;
  background-color: var(--color-dark);
  color: var(--color-highlight);
}

.shrinker {
  background-color: var(--color-highlight);
  font-weight: bold;
  color: var(--color-light);
  text-align: center;
  padding: 0rem;
}

.shrinker:hover {
  background-color: var(--color-dark);
  color: var(--color-highlight);
  border: 1px solid var(--color-highlight);
}

#nav-area {
  display: grid;
  grid-template-columns: 1fr 1rem;
}

main header {
  height: 2rem;
  display: flex;
  justify-content: right;
}

main header>button {
  padding: none;
  border: 1px solid var(--color-highlight);
  border-radius: 0;
  margin: 0;
  line-height: 0;
}

nav {
  font-size: 1rem;
}

.content {
  display: grid;
  grid-template-columns: 20ch 1fr;
}

#sidebar {
  border-right: 2px solid var(--color-light);
  display: grid;
}

#sidebar-content {
  padding: 0 1rem 0 0;
}

h2 {
  color: var(--color-text);
}

.header {
  border-bottom: 2px solid var(--color-light);
  padding-bottom: 1rem;
}

main {
  /* padding: 1rem 0 0 0; */
  min-height: 90vh;
}
</style>
