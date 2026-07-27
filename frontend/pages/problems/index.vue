<script setup lang="ts">
import { watch } from 'vue'
const store = useProblemsStore()
onMounted(() => store.fetchProblems())
watch(() => store.selectedTopics, () => store.fetchProblems(1))
const topics = ['Механика','МКТ','Термодинамика','Электростатика','Магнетизм','Оптика','СТО','Квантовая']
const diffColor = (d:string) => d==='easy'?'green':d==='medium'?'blue':d==='hard'?'red':'gold'
const diffLabel = (d:string) => ({easy:'Легко',medium:'Средне',hard:'Сложно',olympiad:'Олимпиада'} as any)[d]||d
</script>
<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <h1 class="text-3xl md:text-4xl font-bold text-white mb-2">Архив задач по физике</h1>
    <p class="text-gray-400 mb-8">{{store.total}} задач по всем разделам</p>
    <div class="flex flex-wrap gap-2 mb-8">
      <button v-for='t in topics' :key='t' @click='store.toggleTopic(t)'
        class="px-4 py-2 rounded-lg text-sm font-medium transition-all"
        :class="store.selectedTopics.includes(t)?'bg-[#e2b04a] text-[#1a1a2e]':'bg-[#16213e] text-gray-300 border border-gray-600 hover:border-[#e2b04a]/50'">
        {{t}}
      </button>
      <button v-if="store.selectedTopics.length" @click="{store.clearTopics();store.fetchProblems(1)}" class="px-4 py-2 rounded-lg text-sm text-red-400 hover:text-red-300 transition">Сбросить фильтры</button>
    </div>
    <UiBaseLoader v-if='store.loading' />
    <div v-else-if='store.items.length' class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
      <div v-for='p in store.items' :key='p.id' class="bg-[#16213e] rounded-xl border border-gray-700/50 p-5 hover:border-[#e2b04a]/30 transition-all">
        <div class="flex items-center gap-2 mb-3"><UiBaseBadge :label='p.topic' variant='blue' /><UiBaseBadge :label='diffLabel(p.difficulty)' :variant='diffColor(p.difficulty) as any' /></div>
        <h3 class="text-white font-semibold mb-2">{{p.title}}</h3>
        <p class="text-gray-400 text-sm leading-relaxed">{{p.content.slice(0,120)}}{{p.content.length>120?'...':''}}</p>
      </div>
    </div>
    <div v-else class="text-center py-16 text-gray-500">Задач не найдено.</div>
    <UiBasePagination :page='store.page' :total-pages='store.totalPages' @update:page='(p:number)=>store.fetchProblems(p)' />
  </div>
</template>