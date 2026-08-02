import type { User } from '~/types'
export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref('')
  const isLoggedIn = computed(() => !!token.value)
  const userName = computed(() => user.value?.full_name || '')
  const setUser = (u: User, t: string) => { user.value = u; token.value = t }
  const clearAuth = () => { user.value = null; token.value = '' }
  return { user, token, isLoggedIn, userName, setUser, clearAuth }
})
