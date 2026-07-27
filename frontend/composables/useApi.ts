export const useApi = () => {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase as string

  const authHeaders = (): Record<string, string> => {
    if (import.meta.client) {
      const t = localStorage.getItem('auth_token')
      if (t) return { Authorization: `Bearer ${t}` }
    }
    return {}
  }

  const get = async <T>(path: string, params?: Record<string, any>): Promise<T> =>
    await $fetch<T>(path, { baseURL, params, headers: authHeaders() })
  const post = async <T>(path: string, body: any): Promise<T> =>
    await $fetch<T>(path, { method: 'POST', baseURL, body, headers: authHeaders() })

  return { get, post, authHeaders }
}
