definePageMeta({ middleware: ['auth'] })
export default defineNuxtRouteMiddleware((to) => {
  if (import.meta.client && !localStorage.getItem('auth_token')) {
    return navigateTo('/')
  }
})
