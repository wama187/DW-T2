<template>
  <div 
    class="fixed inset-0 bg-black bg-opacity-50 backdrop-blur-sm overflow-y-auto h-full w-full z-50 flex items-center justify-center px-4 animate-fadeIn" 
    @click.self="$emit('close')"
  >
    <div class="relative bg-white rounded-2xl shadow-2xl max-w-md w-full transform transition-all animate-slideUp">
      <!-- Icono de advertencia -->
      <div class="pt-8 pb-4 flex justify-center">
        <div class="bg-red-100 rounded-full p-4 animate-pulse">
          <svg class="w-16 h-16 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
        </div>
      </div>

      <!-- Contenido del Modal -->
      <div class="text-center px-6 pb-6">
        <h3 class="text-2xl font-bold text-gray-900 mb-3">
          ¿Eliminar este libro?
        </h3>
        <p class="text-gray-600 mb-2">
          Esta acción no se puede deshacer. El libro será eliminado permanentemente de tu biblioteca.
        </p>
        <p class="text-sm text-gray-500">
          ¿Estás seguro de que deseas continuar?
        </p>
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
          @click="$emit('confirm')" 
          class="flex-1 px-4 py-3 bg-gradient-to-r from-red-500 to-red-600 hover:from-red-600 hover:to-red-700 text-white font-semibold rounded-lg transition-all duration-200 flex items-center justify-center space-x-2 shadow-md hover:shadow-lg transform hover:scale-105 active:scale-95"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          <span>Eliminar</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineEmits(['close', 'confirm']);
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