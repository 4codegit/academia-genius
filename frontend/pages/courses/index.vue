<script setup lang="ts">
import type { Course } from '~/types'
const courses = ref<Course[]>([])
const loading = ref(true)
const { get } = useApi()
onMounted(async () => { try { courses.value = await get<Course[]>('/courses') } finally { loading.value = false } })
const fmtPrice = (p: number) => p === 0 ? 'Бесплатно' : new Intl.NumberFormat('ru-RU').format(p) + ' сом'
</script>
<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <h1 class="text-3xl md:text-4xl font-bold text-white mb-2">Наши курсы</h1>
    <p class="text-gray-400 mb-10">Выберите программу, которая подходит вам</p>
    <UiBaseLoader v-if="loading" />
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="c in courses" :key="c.id" class="bg-[#16213e] rounded-xl border border-gray-700/50 overflow-hidden hover:border-[#e2b04a]/30 transition-all">
        <div class="aspect-video bg-gradient-to-br from-gray-800 to-gray-900" />
        <div class="p-5">
          <h3 class="text-white font-semibold mb-1">{{c.title}}</h3>
          <p class="text-gray-400 text-sm mb-1">{{c.instructor}}</p>
          <p class="text-gray-500 text-xs mb-3">{{c.duration}}</p>
          <span class="inline-block px-3 py-1 rounded-full text-sm font-medium" :class="c.price===0?'bg-green-500/20 text-green-400':'bg-[#e2b04a]/20 text-[#e2b04a]'">{{fmtPrice(c.price)}}</span>
        </div>
      </div>
    </div>
  </div>
</template>