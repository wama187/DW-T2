import { ref } from 'vue';
import api from '@/services/api';
import type { BookV1, BookV2 } from '@/types';

export function useBooks() {
  const books = ref<(BookV1 | BookV2)[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const fetchBooks = async (version: 'v1' | 'v2') => {
    loading.value = true;
    error.value = null;
    try {
      const response = await api.get(`/${version}/books`);
      books.value = response.data;
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const createBook = async (version: 'v1' | 'v2', book: { title: string; author: string }) => {
    loading.value = true;
    error.value = null;
    try {
      await api.post(`/${version}/books`, book);
      await fetchBooks(version);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const updateBook = async (version: 'v1' | 'v2', book: BookV1 | BookV2) => {
    loading.value = true;
    error.value = null;
    try {
      await api.put(`/${version}/books`, book);
      await fetchBooks(version);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const deleteBook = async (version: 'v1' | 'v2', id: number | string) => {
    loading.value = true;
    error.value = null;
    try {
      await api.delete(`/${version}/books/${id}`);
      await fetchBooks(version);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };


  return { books, loading, error, fetchBooks, createBook, updateBook, deleteBook };
}
