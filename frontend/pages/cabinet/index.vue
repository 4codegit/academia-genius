<script setup lang="ts">
import type { UserStats } from '~/types'
definePageMeta({ middleware: ['auth'] })
const stats = ref<UserStats | null>(null)
const loading = ref(true)
onMounted(async () => {
  try {
    const { get, authHeaders } = useApi()
    stats.value = await get<UserStats>('/user/stats', {}, { headers: authHeaders() })
  } catch { /* redirect handled by middleware */ }
  finally { loading.value = false }
})
const topics = ['Механика','МКТ','Термодинамика','Электростатика','Магнетизм','Оптика','СТО','Квантовая']
const barColor = (p: number) => p >= 70 ? 'bg-green-500' : p >= 40 ? 'bg-[#e2b04a]' : p > 0 ? 'bg-orange-500' : 'bg-gray-600'
</script>
<template>
  <div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <h1 class="text-3xl md:text-4xl font-bold text-white mb-10">Личный кабинет</h1>
    <UiBaseLoader v-if="loading" />
    <template v-else-if="stats">
      <!-- Metrics -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-10">
        <div class="bg-[#16213e] rounded-xl border border-gray-700/50 p-4 text-center"><div class="text-3xl font-bold text-[#e2b04a]">{{stats.total_solved}}</div><div class="text-gray-400 text-sm mt-1">Решено задач</div></div>
        <div class="bg-[#16213e] rounded-xl border border-gray-700/50 p-4 text-center"><div class="text-3xl font-bold text-white">{{stats.streak_days}}</div><div class="text-gray-400 text-sm mt-1">Дней подряд</div></div>
        <div class="bg-[#16213e] rounded-xl border border-gray-700/50 p-4 text-center"><div class="text-3xl font-bold text-white">{{Object.values(stats.by_difficulty).reduce((a,b)=>a+b,0)}}</div><div class="text-gray-400 text-sm mt-1">По сложности</div></div>
        <div class="bg-[#16213e] rounded-xl border border-gray-700/50 p-4 text-center"><div class="text-3xl font-bold text-white">{{Object.values(stats.by_topic).reduce((a,b)=>a+b,0)}}</div><div class="text-gray-400 text-sm mt-1">По темам</div></div>
      </div>
      <!-- Knowledge Map -->
      <h2 class="text-xl font-bold text-white mb-6">Карта знаний</h2>
      <div class="bg-[#16213e] rounded-xl border border-gray-700/50 p-6 space-y-4">
        <div v-for="km in stats.knowledge_map" :key="km.topic" class="flex items-center gap-4">
          <span class="w-28 text-sm text-gray-300 flex-shrink-0">{{km.topic}}</span>
          <div class="flex-1 bg-gray-700 rounded-full h-3 overflow-hidden">
            <div class="h-full rounded-full transition-all duration-500" :class="barColor(km.progress)" :style="{width: km.progress + '%'}" />
          </div>
          <span class="w-10 text-sm text-gray-400 text-right">{{km.progress}}%</span>
        </div>
      </div>
    </template>
  </div>
</template>