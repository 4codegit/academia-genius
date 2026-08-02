import type { Book, PaginatedResponse } from '~/types'
export const useBooksStore = defineStore('books', () => {
  const items = ref<Book[]>([])
  const total = ref(0); const page = ref(1); const totalPages = ref(0); const loading = ref(false)
  const activeCategory = ref('')
  const { get } = useApi()
  const fetchBooks = async (category = '', p = 1, limit = 12) => {
    loading.value = true
    try {
      const params: any = { page: p, limit }
      if (category) params.category = category
      const r = await get<PaginatedResponse<Book>>('/books', params)
      items.value = r.data; total.value = r.total; page.value = r.page; totalPages.value = r.total_pages
    } finally { loading.value = false }
  }
  const setCategory = (c: string) => { activeCategory.value = c }
  return { items, total, page, totalPages, loading, activeCategory, fetchBooks, setCategory }
})
