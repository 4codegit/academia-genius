<script setup lang="ts">
const props = defineProps<{page:number;totalPages:number}>()
const emit = defineEmits<{(e:'update:page',p:number):void}>()
const pages = computed(()=>{const t=props.totalPages;const c=props.page;if(t<=7)return Array.from({length:t},(_,i)=>i+1)
const r=[];r.push(1);if(c>3)r.push(-1);for(let i=Math.max(2,c-1);i<=Math.min(t-1,c+1);i++)r.push(i);if(c<t-2)r.push(-1);r.push(t);return r})
</script>
<template>
  <div v-if="totalPages>1" class="flex items-center justify-center gap-1 mt-8">
    <button :disabled="page<=1" @click="emit('update:page',page-1)" class="px-3 py-1.5 rounded-lg text-sm border border-gray-600 text-gray-300 hover:border-[#e2b04a] disabled:opacity-30 transition">Назад</button>
    <template v-for="p in pages" :key="p">
      <span v-if="p===-1" class="px-2 text-gray-500">...</span>
      <button v-else @click="emit('update:page',p)" class="w-9 h-9 rounded-lg text-sm font-medium transition" :class="p===page?'bg-[#e2b04a] text-[#1a1a2e]':'text-gray-300 hover:bg-gray-700'">{{p}}</button>
    </template>
    <button :disabled="page>=totalPages" @click="emit('update:page',page+1)" class="px-3 py-1.5 rounded-lg text-sm border border-gray-600 text-gray-300 hover:border-[#e2b04a] disabled:opacity-30 transition">Далее</button>
  </div>
</template>
