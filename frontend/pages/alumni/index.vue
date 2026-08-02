<script setup lang="ts">
import type { FeaturedAlumnus, Alumnus } from '~/types'
const featured = ref<FeaturedAlumnus | null>(null)
const others = ref<Alumnus[]>([])
const loading = ref(true)
const { get } = useApi()
onMounted(async () => {
  try {
    const r = await get<{featured: FeaturedAlumnus; others: Alumnus[]}>('/alumni')
    featured.value = r.featured; others.value = r.others
  } finally { loading.value = false }
})
</script>
<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <h1 class="text-3xl md:text-4xl font-bold text-white mb-10">Наши выпускники</h1>
    <UiBaseLoader v-if="loading" />
    <template v-else-if="featured">
      <!-- Featured -->
      <div class="bg-[#16213e] rounded-2xl border border-[#e2b04a]/30 p-6 md:p-8 mb-12">
        <div class="flex flex-col md:flex-row gap-6">
          <div class="w-32 h-32 rounded-full bg-gradient-to-br from-[#e2b04a] to-[#a67c2e] flex-shrink-0 flex items-center justify-center text-4xl font-bold text-[#1a1a2e]">{{featured.full_name.split(' ').map(n=>n[0]).join('')}}</div>
          <div class="flex-1">
            <span class="text-xs text-[#e2b04a] font-medium tracking-wider uppercase">Главный выпускник</span>
            <h2 class="text-2xl md:text-3xl font-bold text-white mt-1 mb-2">{{featured.full_name}}</h2>
            <p class="text-[#e2b04a] text-sm mb-3">{{featured.university}} · {{featured.graduation_year}}</p>
            <p class="text-gray-300 leading-relaxed">{{featured.bio}}</p>
          </div>
        </div>
        <!-- Timeline -->
        <div v-if="featured.awards.length" class="mt-8 pt-6 border-t border-gray-700">
          <h3 class="text-lg font-semibold text-white mb-6">Хронология наград</h3>
          <div class="relative pl-6">
            <div class="absolute left-[7px] top-2 bottom-2 w-0.5 bg-[#e2b04a]/40" />
            <div v-for="a in featured.awards" :key="a.id" class="relative mb-6 last:mb-0">
              <div class="absolute -left-6 top-1.5 w-4 h-4 rounded-full bg-[#e2b04a] border-2 border-[#16213e]" />
              <div class="text-[#e2b04a] text-sm font-medium">{{a.year}} — {{a.competition}}</div>
              <div class="text-white font-semibold mt-0.5">{{a.award_title}}</div>
              <div v-if="a.description" class="text-gray-400 text-sm mt-1">{{a.description}}</div>
            </div>
          </div>
        </div>
      </div>
      <!-- Others -->
      <h2 v-if="others.length" class="text-2xl font-bold text-white mb-6">Другие выпускники</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
        <div v-for="a in others" :key="a.id" class="bg-[#16213e] rounded-xl border border-gray-700/50 p-5 hover:border-[#e2b04a]/30 transition">
          <div class="text-white font-semibold mb-1">{{a.full_name}}</div>
          <p class="text-[#e2b04a] text-sm mb-2">{{a.university}} · {{a.graduation_year}}</p>
          <p class="text-gray-400 text-sm">{{a.bio}}</p>
        </div>
      </div>
    </template>
  </div>
</template>