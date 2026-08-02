import type { AuthResponse } from '~/types'
import { computed } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { useRouter } from 'vue-router'
export const useAuth = () => {
  const { get, post } = useApi()
  const store = useAuthStore()
  const router = useRouter()
  const login = async (email: string, password: string) => {
    const r = await post<AuthResponse>('/auth/login', { email, password })
    if (import.meta.client) {
      localStorage.setItem('auth_token', r.token)
      window.dispatchEvent(new Event('auth-changed'))
    }
    store.setUser(r.user, r.token)
  }
  const register = async (username: string, email: string, password: string, fullName: string) => {
    const r = await post<AuthResponse>('/auth/register', { username, email, password, full_name: fullName })
    if (import.meta.client) {
      localStorage.setItem('auth_token', r.token)
      window.dispatchEvent(new Event('auth-changed'))
    }
    store.setUser(r.user, r.token)
  }
  const forgotPassword = async (email: string, password: string) => {
    await post('/auth/forgot-password', { email, password })
  }
  const logout = async () => {
    try {
      await post('/auth/logout', {})
    } catch (error) {
      // ignore logout errors; still clear local auth state
    }
    if (import.meta.client) {
      localStorage.removeItem('auth_token')
      window.dispatchEvent(new Event('auth-changed'))
    }
    store.clearAuth()
    router.push('/')
  }
  const isAuthenticated = computed(() => {
    if (import.meta.client) return !!localStorage.getItem('auth_token')
    return false
  })
  return { login, register, forgotPassword, logout, isAuthenticated }
}
