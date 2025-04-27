<script setup lang="ts">
import { ref } from "vue";

const message = ref<string>('PassIt, Buy your tickets!');
const status = ref<string>('active');
const tasks = ref<string[]>(['Task One', 'Task Two', 'Task Three']);

const changeStatus = (): void => {
  if (status.value === 'active') {
    status.value = 'pending';
  } else if (status.value === 'pending') {
    status.value = 'inactive';
  } else {
    status.value = 'active';
  }
};

</script>


<template>
  <h1>{{ message }}</h1>
  <br />
  <h2>The user is {{ status }}</h2>
  <br />
  <h3>Tasks:</h3>
  <br />
  <ul>
    <li v-for="task in tasks">{{ task }}</li>
  </ul>
  <br />
  <button @click="changeStatus">Change Status</button>
</template>