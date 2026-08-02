import type { Problem, PaginatedResponse } from '~/types'
export const useProblemsStore = defineStore('problems', () => {
  const items = ref<Problem[]>([])
  const total = ref(0); const page = ref(1); const totalPages = ref(0); const loading = ref(false)
  const selectedTopics = ref<string[]>([])
  const { get } = useApi()
  const fetchProblems = async (p = 1, limit = 20) => {
    loading.value = true
    try {
      const params: any = { page: p, limit }
      if (selectedTopics.value.length) params.topics = selectedTopics.value.join(',')
      const r = await get<PaginatedResponse<Problem>>('/problems', params)
      items.value = r.data; total.value = r.total; page.value = r.page; totalPages.value = r.total_pages
    } finally { loading.value = false }
  }
  const toggleTopic = (t: string) => {
    const i = selectedTopics.value.indexOf(t)
    if (i >= 0) {
      selectedTopics.value.splice(i, 1)
    } else {
      selectedTopics.value.push(t)
    }
  }
  const clearTopics = () => { selectedTopics.value = [] }
  return { items, total, page, totalPages, loading, selectedTopics, fetchProblems, toggleTopic, clearTopics }
})
