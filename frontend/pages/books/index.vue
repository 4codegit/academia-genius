<script setup lang="ts">
const store = useBooksStore()
const cats = ['','Учебники','Задачники','Справочники','Монографии','Подготовка к олимпиадам']
const catLabels: Record<string,string> = {'':'Все','Учебники':'Учебники','Задачники':'Задачники','Справочники':'Справочники','Монографии':'Монографии','Подготовка к олимпиадам':'Олимпиады'}
onMounted(() => store.fetchBooks())
const changeCat = (c: string) => { store.setCategory(c); store.fetchBooks(c, 1) }
</script>
<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <h1 class="text-3xl md:text-4xl font-bold text-white mb-2">Библиотека</h1>
    <p class="text-gray-400 mb-8">Учебники, задачники и справочники по физике</p>
    <div class="flex flex-wrap gap-1 mb-8 border-b border-gray-700 pb-1">
      <button v-for="c in cats" :key="c" @click="changeCat(c)"
        class="px-4 py-2 text-sm font-medium transition-colors"
        :class="store.activeCategory===c?'text-[#e2b04a] border-b-2 border-[#e2b04a]':'text-gray-400 hover:text-white'">
        {{catLabels[c]}}
      </button>
    </div>
    <UiBaseLoader v-if="store.loading" />
    <div v-else-if="store.items.length" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-5">
      <div v-for="b in store.items" :key="b.id" class="bg-[#16213e] rounded-xl border border-gray-700/50 overflow-hidden hover:border-[#e2b04a]/30 transition-all group">
        <div class="aspect-[3/4] bg-gradient-to-b from-gray-700 to-gray-800 flex items-center justify-center">
          <span class="text-gray-600 text-4xl font-bold group-hover:text-[#e2b04a]/30 transition">📖</span>
        </div>
        <div class="p-4">
          <h3 class="text-white font-semibold text-sm mb-1 line-clamp-2">{{b.title}}</h3>
          <p class="text-gray-400 text-xs mb-2">{{b.author}}</p>
          <UiBaseBadge :label="String(b.year)" variant="gold" />
        </div>
      </div>
    </div>
    <div v-else class="text-center py-16 text-gray-500">Книг не найдено.</div>
    <UiBasePagination :page="store.page" :total-pages="store.totalPages" @update:page="(p:number)=>store.fetchBooks(store.activeCategory,p)" />
  </div>
</template>