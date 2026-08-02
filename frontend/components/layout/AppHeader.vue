<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
const mobileOpen = ref(false)
const authed = ref(false)
const router = useRouter()
const { logout } = useAuth()
const updateAuthState = () => {
  if (import.meta.client) {
    authed.value = !!localStorage.getItem('auth_token')
  }
}
if (import.meta.client) {
  onMounted(updateAuthState)
  window.addEventListener('auth-changed', updateAuthState)
  onBeforeUnmount(() => window.removeEventListener('auth-changed', updateAuthState))
}
const doLogout = async () => {
  await logout()
  authed.value = false
}
const links = [{t:'Главная',p:'/'},{t:'Задачи',p:'/problems'},{t:'Курсы',p:'/courses'},{t:'Книги',p:'/books'},{t:'Выпускники',p:'/alumni'},{t:'Расписание',p:'/schedule'}]
</script>
<template>
  <header class="fixed top-0 left-0 right-0 z-50 bg-[#1a1a2e] border-b border-[#e2b04a]/20">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <NuxtLink to="/" class="text-xl font-bold text-[#e2b04a]">Academy Genius</NuxtLink>
        <nav class="hidden md:flex items-center gap-6">
          <NuxtLink v-for="l in links" :key="l.p" :to="l.p" class="text-gray-300 hover:text-[#e2b04a] transition text-sm">{{l.t}}</NuxtLink>
          <template v-if="authed">
            <NuxtLink to="/cabinet" class="text-[#e2b04a] hover:text-[#f0c75e] text-sm">Кабинет</NuxtLink>
            <button @click="doLogout" class="text-gray-400 hover:text-red-400 text-sm">Выйти</button>
          </template>
          <NuxtLink v-else to="/auth" class="bg-[#e2b04a] hover:bg-[#f0c75e] text-[#1a1a2e] px-4 py-1.5 rounded-lg text-sm font-semibold transition">Войти</NuxtLink>
        </nav>
        <button @click="mobileOpen=!mobileOpen" class="md:hidden text-gray-300">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" d="M4 6h16M4 12h16M4 18h16"/></svg>
        </button>
      </div>
    </div>
    <div v-if="mobileOpen" class="md:hidden bg-[#16213e] border-t border-gray-700 p-4 space-y-3">
      <NuxtLink v-for="l in links" :key="l.p" :to="l.p" @click="mobileOpen=false" class="block text-gray-300 hover:text-[#e2b04a]">{{l.t}}</NuxtLink>
      <NuxtLink v-if="authed" to="/cabinet" @click="mobileOpen=false" class="block text-[#e2b04a]">Кабинет</NuxtLink>
      <button v-if="authed" @click="doLogout" class="block text-red-400">Выйти</button>
      <NuxtLink v-else to="/auth" @click="mobileOpen=false" class="block text-[#e2b04a]">Войти</NuxtLink>
    </div>
  </header>
</template>
