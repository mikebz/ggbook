<script lang="ts">
import Message from './components/Message.vue'
import ChatBox from './components/ChatBox.vue'

interface Message {
  text: string
  id: number
  isMine: boolean
  author: string
}

export default {
  name: 'App',

  components: {
    ChatBox,
    Message,
  },

  created() { },
  methods: {
    onSubmit(event: SubmitEvent, text: string) {
      event.preventDefault()

      this.messages.push({
        text: text,
        id: this.messages.length + 1,
        isMine: true,
        author: this.user?.name,
      })
      fetch('http://localhost:8080/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'text/plain',
        },
        body: text,
      })
        .then((response) => response.text())
        .then((data) => {
          this.messages.push({
            text: data,
            id: this.messages.length + 1,
            isMine: false,
            author: 'ai',
          })
        })
        .catch((error) => {
          console.error('Error:', error)
          // Add error message to chat
          this.messages.push({
            text: 'Sorry, there was an error processing your message',
            id: this.messages.length + 1,
            isMine: false,
            author: 'System',
          })
        })
    },
  },
  data: () => ({
    user: { id: '1', name: 'human' },
    messages: [] as Message[],
  }),
}
</script>

<template>
  <div class="app">
    <div ref="messages" class="messages">
      <Message v-for="message in messages" :key="message.id" :class="['message', { right: message.isMine }]"
        :dark="message.isMine" :text="message.text" :author="message.author" />
    </div>

    <ChatBox class="chat-box" @chat-message="onSubmit" />
  </div>
</template>

<style>
* {
  box-sizing: border-box;
}

html {
  font-family: 'Tahoma', sans-serif;
}

body {
  margin: 0;
}

button {
  border: 0;
  background: #004899;
  color: white;
  cursor: pointer;
  padding: 1rem;
}

input {
  border: 0;
  padding: 1rem;
  background: rgba(0, 0, 0, 0.1);
}
</style>

<style scoped>
.app {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.messages {
  flex-grow: 1;
  overflow: auto;
  padding: 1rem;
}

.message+.message {
  margin-top: 1rem;
}

.message.right {
  margin-left: auto;
}
</style>
