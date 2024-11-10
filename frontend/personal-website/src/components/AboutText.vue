<script setup>
import {onMounted, ref} from 'vue'

const isActive = ref(false)

onMounted(() => {
  setTimeout(() => {
    isActive.value = true
  }, 5000)
})
</script>

<template>
  <div class="intro-container" :class="{ active: isActive }">
    <p>
      My journey in tech began at 13 when I started tinkering with Android app development in Java. This early interest
      led me to pursue a master's degree in Computer Science with AI Techniques specialization at Delft University of
      Technology, which I completed in
      2022. <a href="https://repository.tudelft.nl/record/uuid:44218abd-4b7e-4ae8-8f24-a2432acdacdc" target="_blank">My
      master's thesis</a> explored Graph Neural Networks for pure cold start in recommender systems.
      I am mainly interested in implementing microservice (web) systems of enterprise-grade quality.
      <br><br>
      During my studies, I landed my first programming job, where I gained valuable experience applying my skills in a
      business environment. My main focus was on designing and implementing Python web applications. This role also
      introduced me to the operational side of software development, including CI/CD pipelines and containerization with
      Docker.<br><br>
      Currently, I work as a Software Engineer at Flow Traders, the largest ETF market maker in Europe.
      This role allows me to apply my software engineering skills in a high-paced and technology-driven environment.<br><br>
      Outside of work, I'm expanding my skill set by learning Go(lang). When I'm not coding, you can find me at the gym
      or on the padel court, staying active and enjoying a good game.
    </p>
  </div>
</template>
<style scoped>
#name {
  font-weight: bold;
  position: sticky;
  top: 0;
}
</style>
