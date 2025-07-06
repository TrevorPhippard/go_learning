<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const auth = useAuthStore()
const router = useRouter()

const username = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const error = ref(null)

async function register() {
  if (password.value !== confirmPassword.value) {
    error.value = 'Passwords do not match'
    return
  }

  await auth.register(username.value, email.value, password.value)

  if (auth.isAuthenticated) {
    router.push('/profile')
  } else {
    error.value = auth.error
  }
}
</script>

<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-50">
    <div class="w-full max-w-md p-8 bg-white rounded-lg shadow-lg">
      <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">
        Create an Account
      </h2>
      <form @submit.prevent="register" class="space-y-4">
        <div class="space-y-4">
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700"
              >username</label
            >
            <input
              v-model="username"
              type="text"
              placeholder="Username"
              class="w-full p-2 border"
              required
            />
          </div>

          <div>
            <label for="email" class="block text-sm font-medium text-gray-700"
              >Email Address</label
            >
            <input
              v-model="email"
              type="email"
              placeholder="Email"
              class="w-full p-2 border"
              required
            />
          </div>

          <div>
            <label
              for="password"
              class="block text-sm font-medium text-gray-700"
              >Password</label
            >
            <input
              v-model="password"
              type="password"
              placeholder="Password"
              class="w-full p-2 border"
              required
            />
          </div>

          <div>
            <label
              for="confirmPassword"
              class="block text-sm font-medium text-gray-700"
              >Confirm Password</label
            >
            <input
              v-model="confirmPassword"
              type="password"
              placeholder="Confirm Password"
              class="w-full p-2 border"
              required
            />
          </div>

          <button
            type="submit"
            class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
          >
            Create Account
          </button>

          <p v-if="error" class="text-red-500 text-sm">{{ error }}</p>

          <p class="mt-4 text-center text-sm text-gray-600">
            Already have an account?
            <router-link to="/" class="text-blue-600 hover:underline"
              >Login</router-link
            >
          </p>
        </div>
      </form>
    </div>
  </div>
</template>

<style></style>
