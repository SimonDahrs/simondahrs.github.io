<script setup>
import { defineProps } from 'vue'

defineProps({
  src: String,
  title: String
})
</script>

<template>
  <div class="item">
    <img :src="src" class="item-img" />
    <h2 class="item-title">{{ title }}</h2>
  </div>
</template>

<style scoped>
@keyframes wipe-enter {
  0% {
    opacity: 0;
    transform: translateY(-30%);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}
.item {
  padding: 16px;
  opacity: 0;
}

.item-img {
  width: 100%;
}

.item-title {
  font-weight: bold;
  text-align: center;
}

.fade-animation {
  animation: wipe-enter 0.75s 1;
  opacity: 1;
}
</style>
