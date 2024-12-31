<script setup lang="ts">
import NavigationSidebar from '@/components/NavigationSidebar.vue'
import NoteViewer from '@/components/NoteViewer.vue'
import { DirectoryContentRepository } from '@/repository/directory'
import type { DirectoryContent } from '@/types'
import { computed, onMounted, ref, watch, type ComputedRef, type Ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const currentDirectory: Ref<string> = ref('/')

const directoryContentRepository = new DirectoryContentRepository()

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
    console.log(oldPath, '->', newPath)
    selectedFilePath.value = ''
    updateFromRoutePath(newPath)
  },
)

async function updateFromRoutePath(path: Array<string> | string) {
  if (path && typeof path != 'string') {
    currentDirectory.value = '/' + path.join('/') + '/'
  } else {
    currentDirectory.value = '/'
  }

  console.log('currentDirectory', currentDirectory.value)

  directoryContent.value = await directoryContentRepository.getByPath(currentDirectory.value)
}

onMounted(async () => {
  updateFromRoutePath(route.params.path)
  return
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
          <NavigationSidebar  class="sidebar-content" :directory-content="directoryContent"
            :current-directory="currentDirectory" @file-select="(path) => (selectedFilePath = path)" />

            <template v-if="isSidebarVisible">
              <div class="shrinker" @click="hideSidebar()">-</div>
            </template>
            <template v-else>
              <div class="shrinker" @click="showSidebar()">+</div>
            </template>

        </div>
      </div>

      <div id="note-viewer" v-if="selectedFilePath">
        <NoteViewer :path="selectedFilePath" />
      </div>
      <div v-else style="display: flex; justify-content: center;; padding: 1rem">
        <strong style="font-size: 1.5rem; color: rgb(255 255 255 / 0.5)">Please select a file on the left </strong>
      </div>
    </div>
  </main>
</template>

<style scoped>

.shrinker {
  background-color: var(--color-highlight);
  font-weight: bold;
  color: var(--color-dark);
  text-align: center;
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

header {
  height: 2rem;
  display: flex;
  justify-content: right;
}

header>button {
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

#note-viewer {
  /* padding: 0 0 0 1rem; */
}

h2 {
  color: var(--color-text);
}

.header {
  border-bottom: 2px solid var(--color-light);
  padding-bottom: 1rem;
}

main {
  padding: 1rem 0;
  min-height: 90vh;
}
</style>
