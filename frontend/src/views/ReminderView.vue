<script setup lang="ts">
import ReminderBox from '@/components/ReminderBox.vue';
import { ReminderRepository } from '@/repository/reminder';
import type { Reminder } from '@/types';
import { onMounted, ref, type Ref } from 'vue';


const reminderRepository = new ReminderRepository()

const reminders: Ref<Array<Reminder>> = ref([])

onMounted(async() => {
  reminders.value = await reminderRepository.getAll()
})
</script>

<template>
  <main>
  <h1>Reminders</h1>
  <div id="reminder-area">
  <template v-for="reminder in reminders" :key="reminder.date+reminder.todo+reminder.source">
    <ReminderBox :reminder="reminder" />
  </template>
</div>
</main>
</template>

<style scoped>
#reminder-area {
  display: flex;
  gap: 1rem;
}

main {
  padding: 1rem 0;
  min-height: 90vh;
}
</style>
