<template>
  <form class="chat-box" @submit="onSubmit($event)">
    <input v-model="text" placeholder="Write a message" type="text" />
    <button :disabled="text === ''">Send</button>
  </form>
</template>

<script lang="ts">
export default {
  name: 'ChatBox',
  data: () => ({
    text: '',
  }),
  methods: {
    onSubmit(event: Event) {
      this.$emit('chat-message', event, this.text)
      this.text = ''
    },
  },
}
</script>

<style scoped>
.chat-box {
  width: 100%;
  display: flex;
}

input {
  width: min(100%, 20rem);
  flex-grow: 1;
}

button:disabled {
  opacity: 0.5;
}
</style>
