<script setup>
import { onMounted, ref } from 'vue'

const isActive = ref(false)

onMounted(() => {
  setTimeout(() => {
    isActive.value = true
  }, 5000)
})
</script>

<template>
  <div class="intro-container" :class="{ active: isActive }">
    <div>
      When I was thirteen years old, I started programming Android apps in Java. This inspired me to
      pursue a master's degree in Computer Science at Delft University of Technology which I
      obtained in 2022.
    </div>
    <br />
    <div>
      During my studies, I had my first programming job where I learned how to apply my skills in a
      business setting. Mostly, I was concerned with designing and implementing Python web
      applications. Also this is where I first I gained experience in the operations side of things,
      such as CI/CD pipelines, Docker, etc.
    </div>
    <br />
    <div>
      Currently, I am employed as DevOps Engineer at APG Asset Management, which is a company that
      invests for pension funds.
    </div>
  </div>
</template>
<style scoped>
#name {
  font-weight: bold;
  position: sticky;
  top: 0;
}
</style>
