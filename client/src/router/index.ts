import { createRouter, createWebHistory } from 'vue-router';
import Login from '@/views/Login.vue';
import Register from '@/views/Register.vue';
import BookList from '@/views/BookList.vue';
import { useAuthStore } from '@/stores/auth';

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/books',
    name: 'BookList',
    component: BookList,
    meta: { requiresAuth: true },
  },
  { path: '/:pathMatch(.*)*', redirect: '/books' },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();
  if (to.meta.requiresAuth && !authStore.token) {
    next('/login');
  } else if (authStore.token && to.name === 'Login') {
    next('/books');
  } else {
    next();
  }
});

export default router;
