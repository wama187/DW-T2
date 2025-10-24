<template>
  <div class="container mx-auto px-4 py-8">
    <!-- Header con título y botón de versión -->
    <div class="bg-white rounded-lg shadow-md p-6 mb-6">
      <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div>
          <h1 class="text-3xl font-bold text-gray-900 mb-1">Biblioteca</h1>
          <p class="text-gray-600 flex items-center">
            <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-blue-100 text-blue-800">
              {{ bookVersion.toUpperCase() }}
            </span>
            <span class="ml-2">{{ books.length }} {{ books.length === 1 ? 'libro' : 'libros' }}</span>
          </p>
        </div>
        
        <div class="flex gap-2">
          <button 
            @click="toggleVersion" 
            class="px-4 py-2 bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 text-white rounded-lg transition-all duration-200 flex items-center space-x-2 font-medium shadow-md hover:shadow-lg transform hover:scale-105 active:scale-95"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4" />
            </svg>
            <span>Cambiar a {{ bookVersion === 'v1' ? 'V2' : 'V1' }}</span>
          </button>
          
          <button 
            @click="openAddModal" 
            class="px-4 py-2 bg-gradient-to-r from-green-500 to-green-600 hover:from-green-600 hover:to-green-700 text-white rounded-lg transition-all duration-200 flex items-center space-x-2 font-medium shadow-md hover:shadow-lg transform hover:scale-105 active:scale-95"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            <span>Agregar Libro</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Estado de carga -->
    <div v-if="loading" class="flex justify-center items-center py-12">
      <div class="text-center">
        <svg class="animate-spin h-12 w-12 text-blue-600 mx-auto mb-4" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <p class="text-gray-600 font-medium">Cargando libros...</p>
      </div>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
      <div class="flex items-start">
        <svg class="w-6 h-6 text-red-600 mr-3 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
        </svg>
        <div>
          <h3 class="text-red-800 font-semibold mb-1">Error al cargar los libros</h3>
          <p class="text-red-700 text-sm">{{ error }}</p>
        </div>
      </div>
    </div>

    <!-- Lista de libros -->
    <div v-else-if="books.length" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="book in books" 
        :key="getBookId(book)"
        class="bg-white rounded-lg shadow-md hover:shadow-xl transition-all duration-300 overflow-hidden group transform hover:-translate-y-1"
      >
        <!-- Encabezado de la tarjeta con gradiente -->
        <div class="bg-gradient-to-r from-blue-500 to-purple-600 p-4 h-32 flex items-center justify-center relative overflow-hidden">
          <div class="absolute inset-0 bg-black opacity-0 group-hover:opacity-10 transition-opacity duration-300"></div>
          <svg class="w-16 h-16 text-white opacity-90" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
          </svg>
        </div>

        <!-- Contenido del libro -->
        <div class="p-5">
          <h2 class="text-xl font-bold text-gray-900 mb-2 line-clamp-2">{{ book.title }}</h2>
          <p class="text-gray-600 mb-4 flex items-center">
            <svg class="w-4 h-4 mr-2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
            {{ book.author }}
          </p>

          <!-- Botones de acción -->
          <div class="flex gap-2 pt-4 border-t border-gray-100">
            <button 
              @click="openEditModal(book)" 
              class="flex-1 px-3 py-2 bg-yellow-500 hover:bg-yellow-600 text-white rounded-lg transition-all duration-200 flex items-center justify-center space-x-2 font-medium transform hover:scale-105 active:scale-95"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
              <span>Editar</span>
            </button>
            
            <button 
              @click="openDeleteModal(book)" 
              class="flex-1 px-3 py-2 bg-red-500 hover:bg-red-600 text-white rounded-lg transition-all duration-200 flex items-center justify-center space-x-2 font-medium transform hover:scale-105 active:scale-95"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
              <span>Eliminar</span>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Estado vacío -->
    <div v-else class="bg-white rounded-lg shadow-md p-12 text-center">
      <svg class="w-24 h-24 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
      </svg>
      <h3 class="text-xl font-semibold text-gray-700 mb-2">No hay libros aún</h3>
      <p class="text-gray-500 mb-6">Comienza agregando tu primer libro a la biblioteca</p>
      <button 
        @click="openAddModal" 
        class="px-6 py-3 bg-gradient-to-r from-green-500 to-green-600 hover:from-green-600 hover:to-green-700 text-white rounded-lg transition-all duration-200 inline-flex items-center space-x-2 font-medium shadow-md hover:shadow-lg transform hover:scale-105"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
        <span>Agregar Primer Libro</span>
      </button>
    </div>

    <!-- Modales -->
    <BookModal
      v-if="showBookModal"
      :book="editingBook"
      @close="closeBookModal"
      @save="handleSaveBook"
    />
    
    <DeleteConfirmationModal
      v-if="showDeleteModal"
      @close="closeDeleteModal"
      @confirm="handleDeleteBook"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useBooks } from '@/composables/useBooks';
import BookModal from '@/components/BookModal.vue';
import DeleteConfirmationModal from '@/components/DeleteConfirmationModal.vue';
import type { BookV1, BookV2 } from '@/types';

const { books, loading, error, fetchBooks, createBook, updateBook, deleteBook } = useBooks();

const bookVersion = ref<'v1' | 'v2'>('v1');
const showBookModal = ref(false);
const editingBook = ref<BookV1 | BookV2 | null>(null);
const showDeleteModal = ref(false);
const bookToDelete = ref<BookV1 | BookV2 | null>(null);

const toggleVersion = () => {
  bookVersion.value = bookVersion.value === 'v1' ? 'v2' : 'v1';
  fetchBooks(bookVersion.value);
};

const openAddModal = () => {
  editingBook.value = null;
  showBookModal.value = true;
};

const openEditModal = (book: BookV1 | BookV2) => {
  editingBook.value = book;
  showBookModal.value = true;
};

const closeBookModal = () => {
  showBookModal.value = false;
};

const handleSaveBook = async (bookData: { title: string; author: string }) => {
  if (editingBook.value) {
    const updatedBook = { ...editingBook.value, ...bookData };
    await updateBook(bookVersion.value, updatedBook);
  } else {
    await createBook(bookVersion.value, bookData);
  }
  closeBookModal();
};

const openDeleteModal = (book: BookV1 | BookV2) => {
  bookToDelete.value = book;
  showDeleteModal.value = true;
};

const closeDeleteModal = () => {
  showDeleteModal.value = false;
  bookToDelete.value = null;
};

const handleDeleteBook = async () => {
  if (bookToDelete.value) {
    const id = 'id' in bookToDelete.value ? bookToDelete.value.id : bookToDelete.value._id;
    await deleteBook(bookVersion.value, id);
  }
  closeDeleteModal();
};

onMounted(() => {
  fetchBooks(bookVersion.value);
});

function getBookId(book: BookV1 | BookV2): string | number {
  return 'id' in book ? book.id : book._id;
}

</script>