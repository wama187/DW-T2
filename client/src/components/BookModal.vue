<template>
  <div 
    class="fixed inset-0 bg-black bg-opacity-50 backdrop-blur-sm overflow-y-auto h-full w-full z-50 flex items-center justify-center px-4 animate-fadeIn" 
    @click.self="$emit('close')"
  >
    <div class="relative bg-white rounded-2xl shadow-2xl max-w-md w-full transform transition-all animate-slideUp">
      <!-- Header del Modal -->
      <div class="bg-gradient-to-r from-blue-500 to-purple-600 rounded-t-2xl p-6 text-white">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-3">
            <div class="bg-white bg-opacity-20 p-2 rounded-lg">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
              </svg>
            </div>
            <h3 class="text-2xl font-bold">{{ book ? 'Editar' : 'Agregar' }} Libro</h3>
          </div>
          <button 
            @click="$emit('close')" 
            class="text-white hover:bg-white hover:bg-opacity-20 rounded-lg p-2 transition-all duration-200"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Contenido del Modal -->
      <div class="p-6 space-y-4">
        <!-- Campo Título -->
        <div>
          <label for="title" class="block text-sm font-semibold text-gray-700 mb-2">
            Título del libro
          </label>
          <div class="relative">
            <input 
              id="title"
              type="text" 
              v-model="form.title" 
              placeholder="Ej: Cien años de soledad" 
              class="w-full px-4 py-3 text-gray-700 border-2 border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 transition-colors duration-200 pl-11"
              required
            />
            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Campo Autor -->
        <div>
          <label for="author" class="block text-sm font-semibold text-gray-700 mb-2">
            Autor
          </label>
          <div class="relative">
            <input 
              id="author"
              type="text" 
              v-model="form.author" 
              placeholder="Ej: Gabriel García Márquez" 
              class="w-full px-4 py-3 text-gray-700 border-2 border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 transition-colors duration-200 pl-11"
              required
            />
            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer con botones -->
      <div class="bg-gray-50 px-6 py-4 rounded-b-2xl flex gap-3">
        <button 
          @click="$emit('close')" 
          class="flex-1 px-4 py-3 bg-gray-200 hover:bg-gray-300 text-gray-700 font-semibold rounded-lg transition-all duration-200 flex items-center justify-center space-x-2 transform hover:scale-105 active:scale-95"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
          <span>Cancelar</span>
        </button>
        <button 
          @click="saveBook" 
          :disabled="!form.title || !form.author"
          class="flex-1 px-4 py-3 bg-gradient-to-r from-green-500 to-green-600 hover:from-green-600 hover:to-green-700 text-white font-semibold rounded-lg transition-all duration-200 flex items-center justify-center space-x-2 shadow-md hover:shadow-lg disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none transform hover:scale-105 active:scale-95"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          <span>Guardar</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import type { BookV1, BookV2 } from '@/types';

const props = defineProps<{ book: BookV1 | BookV2 | null }>();
const emit = defineEmits(['close', 'save']);

const form = ref({ title: '', author: '' });

watch(() => props.book, (newBook) => {
  if (newBook) {
    form.value.title = newBook.title;
    form.value.author = newBook.author;
  } else {
    form.value.title = '';
    form.value.author = '';
  }
}, { deep: true, immediate: true });

const saveBook = () => {
  if (form.value.title && form.value.author) {
    emit('save', form.value);
  }
};
</script>

<style scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fadeIn {
  animation: fadeIn 0.2s ease-out;
}

.animate-slideUp {
  animation: slideUp 0.3s ease-out;
}
</style>