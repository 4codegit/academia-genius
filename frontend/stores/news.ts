import type { News, PaginatedResponse } from '~/types'
export const useNewsStore = defineStore('news', () => {
  const items = ref<News[]>([])
  const total = ref(0)
  const page = ref(1)
  const totalPages = ref(0)
  const loading = ref(false)
  const { get } = useApi()
  const fetchNews = async (p = 1, limit = 10) => {
    loading.value = true
    try {
      const r = await get<PaginatedResponse<News>>('/news', { page: p, limit })
      items.value = r.data; total.value = r.total; page.value = r.page; totalPages.value = r.total_pages
    } finally { loading.value = false }
  }
  return { items, total, page, totalPages, loading, fetchNews }
})
