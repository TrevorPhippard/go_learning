<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const email = ref('')
const password = ref('')
const router = useRouter()
const auth = useAuthStore()

async function submit() {
  await auth.login(email.value, password.value)
  if (auth.isAuthenticated) {
    router.push('/profile')
  }
}
</script>

<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-50">
    <div class="w-full max-w-md p-8 bg-white rounded-lg shadow-lg">
      <div class="p-4 max-w-md mx-auto">
        <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">Login</h2>
        <form @submit.prevent="submit" class="space-y-4">
          <input
            v-model="email"
            type="email"
            placeholder="Email"
            class="w-full p-2 border"
            required
          />
          <input
            v-model="password"
            type="password"
            placeholder="Password"
            class="w-full p-2 border"
            required
          />
          <button class="bg-blue-600 text-white px-4 py-2">Login</button>
        </form>
        <p v-if="auth.error" class="text-red-600 mt-2">{{ auth.error }}</p>
      </div>
      <div class="p-4 max-w-md mx-auto">
        <!-- form here -->
        <p class="mt-2">
          Don't have an account?
          <router-link to="/register" class="text-blue-600 hover:underline"
            >Register</router-link
          >
        </p>
      </div>
    </div>
  </div>
</template>
