import { defineStore } from 'pinia';
import axios from 'axios';
import { useAuthStore } from './auth';

export const usePostStore = defineStore('posts', {
  state: () => ({
    posts: [],
    error: null,
  }),

  actions: {
    async fetchPosts() {
      try {
        const auth = useAuthStore();
        const res = await axios.get('http://localhost:8080/posts', {
          headers: { Authorization: `Bearer ${auth.token}` },
        });
        this.posts = res.data;
      } catch (err) {
        this.error = 'Failed to fetch posts';
      }
    },

    async createPost(content) {
      try {
        const auth = useAuthStore();
        await axios.post(
          'http://localhost:8080/posts',
          { content },
          { headers: { Authorization: `Bearer ${auth.token}` } }
        );
        await this.fetchPosts();
      } catch (err) {
        this.error = 'Failed to create post';
      }
    },
  },
});
