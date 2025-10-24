<template>
  <div class="min-h-[calc(100vh-200px)] flex items-center justify-center px-4 py-12">
    <div class="w-full max-w-md">
      <!-- Encabezado -->
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Crear cuenta</h1>
        <p class="text-gray-600">Crea una nueva cuenta para comenzar</p>
      </div>

      <!-- Tarjeta-->
      <div class="bg-white rounded-lg shadow-lg p-8">
        <form @submit.prevent="register" class="space-y-6">
          <!-- Nombre -->
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700 mb-2">
              Nombre completo
            </label>
            <div class="relative">
              <input
                id="name"
                type="text"
                v-model="name"
                placeholder="Juan Pérez"
                required
                :disabled="isLoading"
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 disabled:bg-gray-100 disabled:cursor-not-allowed"
                :class="{ 'border-red-500 focus:ring-red-500': error }"
              />
              <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
            </div>
          </div>

          <!-- Email -->
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-2">
              Correo electrónico
            </label>
            <div class="relative">
              <input
                id="email"
                type="email"
                v-model="email"
                placeholder="tu@email.com"
                required
                :disabled="isLoading"
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 disabled:bg-gray-100 disabled:cursor-not-allowed"
                :class="{ 'border-red-500 focus:ring-red-500': error }"
              />
              <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
                </svg>
              </div>
            </div>
          </div>

          <!-- Password -->
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
              Contraseña
            </label>
            <div class="relative">
              <input
                id="password"
                :type="showPassword ? 'text' : 'password'"
                v-model="password"
                placeholder="••••••••"
                required
                minlength="6"
                :disabled="isLoading"
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 disabled:bg-gray-100 disabled:cursor-not-allowed"
                :class="{ 'border-red-500 focus:ring-red-500': error }"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-400 hover:text-gray-600"
                :disabled="isLoading"
              >
                <svg v-if="!showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
                </svg>
              </button>
            </div>
            <p class="mt-2 text-xs text-gray-500">Mínimo 6 caracteres</p>
          </div>

          <!-- Error -->
          <div v-if="error" class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg flex items-start">
            <svg class="w-5 h-5 mr-2 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
            </svg>
            <span class="text-sm">{{ error }}</span>
          </div>

          <!-- Éxito -->
          <div v-if="success" class="bg-green-50 border border-green-200 text-green-700 px-4 py-3 rounded-lg flex items-start">
            <svg class="w-5 h-5 mr-2 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
            </svg>
            <span class="text-sm">{{ success }}</span>
          </div>

          <!-- Registro -->
          <button
            type="submit"
            :disabled="isLoading"
            class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-4 rounded-lg transition-all duration-200 flex items-center justify-center disabled:bg-gray-400 disabled:cursor-not-allowed transform active:scale-95"
          >
            <svg v-if="isLoading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span v-if="isLoading">Creando cuenta...</span>
            <span v-else>Crear cuenta</span>
          </button>
        </form>

        <!-- Divisor -->
        <div class="mt-6 text-center">
          <p class="text-sm text-gray-600">
            ¿Ya tienes una cuenta? 
            <router-link to="/login" class="text-blue-600 hover:text-blue-700 font-semibold hover:underline">
              Inicia sesión aquí
            </router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const name = ref('');
const email = ref('');
const password = ref('');
const showPassword = ref(false);
const isLoading = ref(false);
const error = ref('');
const success = ref('');

const authStore = useAuthStore();
const router = useRouter();

const register = async () => {
  error.value = '';
  success.value = '';
  isLoading.value = true;
  
  try {
    await authStore.register(name.value, email.value, password.value);
    success.value = '¡Registro exitoso! Redirigiendo al inicio de sesión...';
    
    setTimeout(() => {
      router.push('/login');
    }, 2000);
  } catch (err: any) {
    error.value = err.response?.data?.message || 'Error al registrarse. Por favor, intenta de nuevo.';
    console.error(err);
  } finally {
    isLoading.value = false;
  }
};
</script>