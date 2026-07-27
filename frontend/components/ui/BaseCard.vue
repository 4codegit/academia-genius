<script setup lang="ts">defineProps<{title?:string}>()</script>
<template>
  <div class="bg-[#16213e] rounded-xl border border-gray-700/50 hover:border-[#e2b04a]/30 transition-all duration-300">
    <h3 v-if="title" class="text-lg font-semibold text-white p-4 pb-0">{{title}}</h3>
    <div class="p-4"><slot/></div>
  </div>
</template>
