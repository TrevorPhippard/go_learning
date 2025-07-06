import { defineStore } from 'pinia';
import axios from 'axios';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: null,
    error: null,
  }),

  actions: {
    async login(email, password) {
      try {
        const res = await axios.post('http://localhost:8080/login', {
          email,
          password,
        });
        this.token = res.data.token;
        localStorage.setItem('token', this.token);
        this.error = null;
        await this.fetchProfile();
      } catch (err) {
        this.error = 'Invalid credentials';
      }
    },
    async register(username, email, password) {
      try {
        const res = await axios.post(
          'http://localhost:8080/register',
          {
            username,
            email,
            password,
          }
        );
        this.token = res.data.token;
        localStorage.setItem('token', this.token);
        this.error = null;
        await this.fetchProfile();
      } catch (err) {
        this.error =
          err.response?.data?.message || 'Registration failed';
      }
    },

    async fetchProfile() {
      if (!this.token) return;
      try {
        const res = await axios.get('http://localhost:8080/profile', {
          headers: { Authorization: `Bearer ${this.token}` },
        });
        this.user = res.data;
      } catch {
        this.logout();
      }
    },

    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem('token');
    },
  },

  getters: {
    isAuthenticated: (state) => !!state.token,
  },
});
