<script setup lang="ts">
</script>

<template>
  <main>
  <h1>Info</h1>
  <h3>
    Reminders
  </h3>
  <p>
    Onyx enables you to improve upon the markdown specs by using specific expressions in the shape of [[type:date]].
    They have to be inside a checkbox list.
  </p>

  <br>

  <h4>
    Birthday reminders
  </h4>
  <p>
    To remind you of birthdays, you use the type <strong>birthday</strong>
  </p>
  <code>
    - [ ] Max Mustermann [[birthday:1990-05-25]]
  </code>

  <br>
  <br>

  <h4>
    Specific date reminders
  </h4>
  <p>
    To remind you on a specific date, you use the type <strong>date</strong>
  </p>
  <code>
    - [ ] Today is new years eve [[date:2025-12-31]]
  </code>

  <br>
  <br>

  <h4>
    Deadline reminders
  </h4>
  <p>
    To remind you on and <strong>after</strong> a specific date is reached, you use the type <strong>deadline</strong>
  </p>
  <code>
    - [ ] Planting a tree [[deadline:2024-01-12]]
  </code>
  <p>
    On and after the specified day, you will receive reminders for unchecked dates.
  </p>

  <br>

  <h4>
    Until reminders
  </h4>
  <p>
    To remind you on and <strong>before</strong> a specific date is reached, you use the type <strong>until</strong>
  </p>
  <code>
    - [ ] Do the dishes [[until:2024-01-12]]
  </code>
  <p>
    Before and on the specified day, you will receive reminders for unchecked dates.
  </p>
</main>
</template>

<style scoped>


main {
  padding: 1rem 0;
  min-height: 90vh;
}
</style>
