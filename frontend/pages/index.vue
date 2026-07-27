<script setup lang='ts'>
const newsStore = useNewsStore()
const fmt = (iso: string) => { try { return new Intl.DateTimeFormat('ru-RU',{day:'numeric',month:'long',year:'numeric'}).format(new Date(iso)) } catch { return iso } }
onMounted(() => newsStore.fetchNews())
</script>
<template>
  <div>
    <section class='relative overflow-hidden border-b border-white/5'>
      <div class='absolute inset-0 bg-gradient-to-br from-[#16213e] via-[#0f0f23] to-[#0f0f23]' />
      <div class='absolute -top-24 -right-24 w-96 h-96 rounded-full bg-[#e2b04a]/10 blur-3xl' />
      <div class='relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 md:py-32 text-center'>
        <span class='inline-block px-3 py-1 mb-6 text-xs font-medium tracking-wider uppercase text-[#e2b04a] border border-[#e2b04a]/30 rounded-full'>Физика · Олимпиады · Победа</span>
        <h1 class='text-5xl md:text-7xl font-extrabold tracking-tight mb-6'><span class='bg-gradient-to-r from-[#e2b04a] via-[#f0c75e] to-[#e2b04a] bg-clip-text text-transparent'>Academy Genius</span></h1>
        <p class='max-w-2xl mx-auto text-lg md:text-xl text-gray-300 leading-relaxed mb-10'>Академия физики для будущих олимпиадников. Готовим к Всероссийской олимпиаде и поступлению в ведущие вузы.</p>
        <div class='flex flex-wrap items-center justify-center gap-4'>
          <UiBaseButton to='/problems' size='lg'>Перейти к задачам</UiBaseButton>
          <UiBaseButton to='/courses' size='lg' variant='secondary'>Наши курсы</UiBaseButton>
        </div>
      </div>
    </section>
    <section class='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16'>
      <h2 class='text-3xl md:text-4xl font-bold text-white mb-2'>Последние новости</h2>
      <p class='text-gray-400 mb-10'>События, достижения и анонсы академии</p>
      <UiBaseLoader v-if='newsStore.loading' />
      <div v-else class='grid grid-cols-1 md:grid-cols-2 gap-6'>
        <article v-for='n in newsStore.items' :key='n.id' class='group bg-[#16213e] rounded-xl border border-gray-700/50 overflow-hidden hover:border-[#e2b04a]/30 transition-all duration-300 flex flex-col'>
          <div class='relative aspect-video bg-gradient-to-br from-gray-800 to-gray-900 flex items-center justify-center'><span class='absolute top-3 left-3 px-2.5 py-0.5 text-xs font-medium text-[#0f0f23] bg-[#e2b04a] rounded-full'>Новость</span></div>
          <div class='p-6 flex flex-col flex-1'>
            <div class='text-xs text-gray-500 mb-2'>{{fmt(n.published_at)}}</div>
            <h3 class='text-lg font-semibold text-white mb-2 group-hover:text-[#e2b04a] transition-colors'>{{n.title}}</h3>
            <p class='text-sm text-gray-400 leading-relaxed flex-1'>{{n.summary.slice(0,140)}}{{n.summary.length>140?'...':''}}</p>
          </div>
        </article>
      </div>
      <UiBasePagination :page='newsStore.page' :total-pages='newsStore.totalPages' @update:page='(p:number)=>newsStore.fetchNews(p)' />
    </section>
  </div>
</template>