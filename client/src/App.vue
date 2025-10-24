<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100">
    <!-- Navbar -->
    <nav class="bg-white shadow-lg border-b border-gray-200">
      <div class="container mx-auto px-4">
        <div class="flex justify-between items-center h-16">
          <!-- Logo y nombre -->
          <router-link 
            to="/" 
            class="flex items-center space-x-2 text-gray-800 hover:text-blue-600 transition-colors duration-200 group"
          >
            <div class="bg-blue-600 p-2 rounded-lg group-hover:bg-blue-700 transition-colors duration-200">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
              </svg>
            </div>
            <span class="text-xl font-bold">BookApp</span>
          </router-link>

          <!-- Menú desktop -->
          <div class="hidden md:flex items-center space-x-1">
            <template v-if="authStore.token">
              <!-- Link a books -->
              <router-link 
                to="/books" 
                class="px-4 py-2 rounded-lg text-gray-700 hover:bg-gray-100 hover:text-blue-600 transition-all duration-200 font-medium"
                active-class="bg-blue-50 text-blue-600"
              >
                <div class="flex items-center space-x-2">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                  </svg>
                  <span>Libros</span>
                </div>
              </router-link>

              <!-- Botón logout -->
              <button 
                @click="logout" 
                class="ml-2 px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-all duration-200 font-medium flex items-center space-x-2 transform hover:scale-105 active:scale-95"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span>Cerrar sesión</span>
              </button>
            </template>

            <template v-else>
              <!-- Link Login -->
              <router-link 
                to="/login" 
                class="px-4 py-2 rounded-lg text-gray-700 hover:bg-gray-100 hover:text-blue-600 transition-all duration-200 font-medium"
                active-class="bg-blue-50 text-blue-600"
              >
                <div class="flex items-center space-x-2">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
                  </svg>
                  <span>Iniciar sesión</span>
                </div>
              </router-link>

              <!-- Link Register -->
              <router-link 
                to="/register" 
                class="ml-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-all duration-200 font-medium flex items-center space-x-2 transform hover:scale-105 active:scale-95"
                active-class="bg-blue-700"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
                </svg>
                <span>Registrarse</span>
              </router-link>
            </template>
          </div>

          <!-- Botón menú Mobile -->
          <button 
            @click="mobileMenuOpen = !mobileMenuOpen"
            class="md:hidden p-2 rounded-lg text-gray-600 hover:bg-gray-100 transition-colors duration-200"
          >
            <svg v-if="!mobileMenuOpen" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
            <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Menú mobile -->
        <div 
          v-if="mobileMenuOpen" 
          class="md:hidden pb-4 border-t border-gray-200 mt-2 pt-4 space-y-2"
        >
          <template v-if="authStore.token">
            <router-link 
              to="/books" 
              @click="mobileMenuOpen = false"
              class="block px-4 py-3 rounded-lg text-gray-700 hover:bg-gray-100 hover:text-blue-600 transition-all duration-200 font-medium"
              active-class="bg-blue-50 text-blue-600"
            >
              <div class="flex items-center space-x-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                </svg>
                <span>Libros</span>
              </div>
            </router-link>

            <button 
              @click="logout" 
              class="w-full text-left px-4 py-3 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-all duration-200 font-medium"
            >
              <div class="flex items-center space-x-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span>Cerrar sesión</span>
              </div>
            </button>
          </template>

          <template v-else>
            <router-link 
              to="/login" 
              @click="mobileMenuOpen = false"
              class="block px-4 py-3 rounded-lg text-gray-700 hover:bg-gray-100 hover:text-blue-600 transition-all duration-200 font-medium"
              active-class="bg-blue-50 text-blue-600"
            >
              <div class="flex items-center space-x-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
                </svg>
                <span>Iniciar sesión</span>
              </div>
            </router-link>

            <router-link 
              to="/register" 
              @click="mobileMenuOpen = false"
              class="block px-4 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-all duration-200 font-medium"
              active-class="bg-blue-700"
            >
              <div class="flex items-center space-x-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
                </svg>
                <span>Registrarse</span>
              </div>
            </router-link>
          </template>
        </div>
      </div>
    </nav>

    <!-- Contenido principal -->
    <main class="container mx-auto px-4 py-6">
      <router-view />
    </main>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();
const router = useRouter();
const mobileMenuOpen = ref(false);

const logout = () => {
  authStore.logout();
  mobileMenuOpen.value = false;
  router.push('/login');
};
</script>