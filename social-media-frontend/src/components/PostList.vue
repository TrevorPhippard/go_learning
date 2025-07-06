<script setup>
defineProps({ posts: Array })
</script>

<template>
  <div class="mt-6">
    <h3 class="text-xl font-semibold mb-2">Posts</h3>
    <ul class="space-y-4">
      <li v-for="post in posts" :key="post.id" class="border p-3 rounded shadow-sm">
        <p>{{ post.content }}</p>
        <small class="text-gray-500">{{ new Date(post.createdAt).toLocaleString() }}</small>
      </li>
    </ul>
  </div>
</template>
