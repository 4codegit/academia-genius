<script setup lang="ts">
withDefaults(defineProps<{label:string;variant?:'gold'|'green'|'red'|'blue'}>(),{variant:'gold'})
</script>
<template>
  <span class="inline-block px-2 py-0.5 text-xs font-medium rounded-full"
    :class="variant==='gold'?'bg-[#e2b04a]/20 text-[#e2b04a]':variant==='green'?'bg-green-500/20 text-green-400':variant==='red'?'bg-red-500/20 text-red-400':'bg-blue-500/20 text-blue-400'">
    {{label}}
  </span>
</template>
