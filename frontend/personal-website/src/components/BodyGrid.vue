<script setup>
import BodyItem from '@/components/BodyItem.vue'

import { defineProps, onMounted } from 'vue'

defineProps({
  items: Array
})

onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add('fade-animation')
        // scrollDown.classList.value.add('hide');
      } else {
        // We're not intersecting, so remove the class!
        entry.target.classList.remove('fade-animation')
      }
    })
  })
  const bodyItems = document.querySelectorAll('.item')
  bodyItems.forEach((element) => observer.observe(element))
})
</script>

<template>
  <div class="item-container">
    <BodyItem v-for="item in items" :src="item.src" :title="item.title" :key="item.id" />
  </div>
</template>

<style scoped>
.item-container {
  display: grid;
  grid-template-columns: 50% 50%;
}
</style>
