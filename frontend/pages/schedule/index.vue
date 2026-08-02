<script setup lang="ts">
import type { Webinar } from '~/types'
const webinars = ref<Webinar[]>([])
const loading = ref(true)
const { get } = useApi()
onMounted(async () => { try { webinars.value = await get<Webinar[]>('/schedule') } finally { loading.value = false } })
const fmtDate = (iso: string) => { try { const d = new Date(iso); return new Intl.DateTimeFormat('ru-RU',{day:'numeric',month:'long',year:'numeric'}).format(d) } catch { return iso } }
const fmtTime = (iso: string) => { try { return new Date(iso).toLocaleTimeString('ru-RU',{hour:'2-digit',minute:'2-digit'}) } catch { return '' } }
</script>
<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <h1 class="text-3xl md:text-4xl font-bold text-white mb-2">Расписание вебинаров</h1>
    <p class="text-gray-400 mb-10">Ближайшие онлайн-мероприятия</p>
    <UiBaseLoader v-if="loading" />
    <div v-else-if="webinars.length" class="space-y-4">
      <div v-for="w in webinars" :key="w.id" class="bg-[#16213e] rounded-xl border border-gray-700/50 p-5 flex flex-col sm:flex-row gap-5 hover:border-[#e2b04a]/30 transition">
        <div class="flex-shrink-0 text-center sm:text-left bg-[#0f0f23] rounded-lg px-4 py-3 min-w-[90px]">
          <div class="text-2xl font-bold text-[#e2b04a]">{{new Date(w.event_date).getDate()}}</div>
          <div class="text-xs text-gray-400 uppercase">{{new Date(w.event_date).toLocaleDateString('ru-RU',{month:'short'})}}</div>
          <div class="text-xs text-gray-500 mt-1">{{fmtTime(w.event_date)}}</div>
        </div>
        <div class="flex-1">
          <h3 class="text-white font-semibold mb-1">{{w.title}}</h3>
          <p class="text-gray-400 text-sm mb-2">Спикер: {{w.speaker}} · {{w.duration_min}} мин</p>
          <p class="text-gray-500 text-sm">{{w.description}}</p>
        </div>
        <div class="flex-shrink-0 self-start">
          <a v-if="w.platform_url" :href="w.platform_url" target="_blank" class="inline-block bg-[#e2b04a] hover:bg-[#f0c75e] text-[#1a1a2e] px-4 py-2 rounded-lg text-sm font-semibold transition">Подключиться</a>
        </div>
      </div>
    </div>
    <div v-else class="text-center py-16 text-gray-500">Нет запланированных вебинаров</div>
  </div>
</template>