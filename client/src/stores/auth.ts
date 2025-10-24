import { defineStore } from 'pinia';
import api from '@/services/api';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    user: null,
  }),
  actions: {
    async login(email: string, password: string) {
      const response = await api.post('/login', { email, password });
      this.token = response.data.access_token;
      localStorage.setItem('token', this.token);
      api.defaults.headers.common['Authorization'] = `Bearer ${this.token}`;
    },
    async register(name: string, email: string, password: string) {
      await api.post('/register', { name, email, password });
    },
    logout() {
      this.token = '';
      this.user = null;
      localStorage.removeItem('token');
      api.defaults.headers.common['Authorization'] = '';
    },
  },
});
