import { defineStore } from 'pinia';
import axios from 'axios';

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token') || null,
  }),
  actions: {
    async login(email, password) {
      const res = await axios.post('http://localhost:8080/login', {
        email,
        password,
      });
      this.token = res.data.token;
      localStorage.setItem('token', this.token);
      await this.fetchUser();
    },
    async fetchUser() {
      if (!this.token) return;
      const res = await axios.get('http://localhost:8080/profile', {
        headers: { Authorization: `Bearer ${this.token}` },
      });
      this.user = res.data;
    },
    logout() {
      this.user = null;
      this.token = null;
      localStorage.removeItem('token');
    },
  },
});
