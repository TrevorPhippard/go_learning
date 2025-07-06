<script setup>
import { onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import { usePostStore } from '../stores/posts'
import PostList from './PostList.vue'
import CreatePost from './CreatePost.vue'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const posts = usePostStore()
const router = useRouter()

function logout() {
  auth.logout()
  router.push('/')
}

onMounted(() => {
  auth.fetchProfile()
  posts.fetchPosts()
})
</script>

<template>
  <div class="p-4 max-w-2xl mx-auto">
    <h2 class="text-2xl font-bold mb-4">Welcome, {{ auth.user?.username }}</h2>
    <p>Email: {{ auth.user?.email }}</p>

    <CreatePost />
    <PostList :posts="posts.posts" />

    <button @click="logout" class="mt-4 bg-red-600 text-white px-4 py-2">Logout</button>
  </div>
</template>
