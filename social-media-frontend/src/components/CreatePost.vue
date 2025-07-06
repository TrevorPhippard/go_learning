<script setup>
import { ref } from 'vue'
import { usePostStore } from '../stores/posts'

const content = ref('')
const posts = usePostStore()

async function submit() {
  if (content.value.trim()) {
    await posts.createPost(content.value)
    content.value = ''
  }
}
</script>

<template>
  <form @submit.prevent="submit" class="mt-4 space-y-2">
    <textarea v-model="content" class="w-full p-2 border" placeholder="Write a post..." required></textarea>
    <button class="bg-green-600 text-white px-4 py-2">Post</button>
    <p v-if="posts.error" class="text-red-600">{{ posts.error }}</p>
  </form>
</template>
