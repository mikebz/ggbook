<template>
  <div :class="['message', { dark }]">
    <h5>{{ author }}</h5>
    {{ text }}
  </div>
</template>

<script lang="ts">
export default {
  name: 'ChatMessage',
  props: [
    'text', // Content of the message
    'author', // Author of the message
    'dark', // Background variant of the box
  ],
}
</script>

<style scoped>
.message {
  background: #e5e5e5;
  padding: 1rem;
  width: fit-content;
}

.message.dark {
  background: #b1d2f7;
}

h5 {
  margin: 0 0 0.5rem 0;
}
</style>
