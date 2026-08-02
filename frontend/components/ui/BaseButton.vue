<script setup lang="ts">
const props = withDefaults(defineProps<{variant?:'primary'|'secondary'|'ghost';size?:'sm'|'md'|'lg';disabled?:boolean;to?:string}>(),{variant:'primary',size:'md',disabled:false})
</script>
<template>
  <NuxtLink v-if="to" :to="to" class="inline-flex items-center justify-center font-semibold rounded-lg transition-all duration-200 focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed"
    :class="[
      size==='sm'?'px-3 py-1.5 text-sm':size==='lg'?'px-6 py-3 text-lg':'px-4 py-2 text-sm',
      variant==='primary'?'bg-[#e2b04a] hover:bg-[#f0c75e] text-[#1a1a2e]':variant==='secondary'?'border border-[#e2b04a] text-[#e2b04a] hover:bg-[#e2b04a]/10':'text-[#e2b04a] hover:text-[#f0c75e]'
    ]" :disabled="disabled">
    <slot/>
  </NuxtLink>
  <button v-else class="inline-flex items-center justify-center font-semibold rounded-lg transition-all duration-200 focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed"
    :class="[
      size==='sm'?'px-3 py-1.5 text-sm':size==='lg'?'px-6 py-3 text-lg':'px-4 py-2 text-sm',
      variant==='primary'?'bg-[#e2b04a] hover:bg-[#f0c75e] text-[#1a1a2e]':variant==='secondary'?'border border-[#e2b04a] text-[#e2b04a] hover:bg-[#e2b04a]/10':'text-[#e2b04a] hover:text-[#f0c75e]'
    ]" :disabled="disabled">
    <slot/>
  </button>
</template>
