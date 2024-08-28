<script setup lang="ts">
import AboutText from '@/components/AboutText.vue'
import CurriculumVitae from '@/components/CurriculumVitae.vue';
import NavigationBar from '@/components/NavigationBar.vue'
import PersonalInfo from '@/components/PersonalInfo.vue'
import { onBeforeMount, onBeforeUnmount, onMounted, ref } from 'vue'

const isActive = ref(false)

onBeforeMount(() => {
  isActive.value = false
})

onMounted(() => {
  setTimeout(() => {
    isActive.value = true
  }, 0)
  window.addEventListener('scroll', function () {
    const sections = document.querySelectorAll('section');
    const navLinks = document.querySelectorAll('#menu a');

    sections.forEach(section => {
      const sectionTop = section.offsetTop;
      const sectionHeight = section.offsetHeight;

      

      if (window.scrollY >= sectionTop && window.scrollY < sectionTop + sectionHeight) {
        navLinks.forEach(link => {
          console.log(window.scrollY, sectionTop, sectionTop + sectionHeight, true)
          if (link.getAttribute("data-section") === section.getAttribute("data-section")) {
            console.log("activate", section.getAttribute("data-section"))
            link.classList.add('active');
          } else {
            console.log("deactivate", section.getAttribute("data-section"))
            link.classList.remove('active');
          }
        });
      }
    });
  });


})

onBeforeUnmount(() => {
  isActive.value = false
})
</script>

<template>
  <main>
    <div class="flex-container">
      <div class="left-col col">
        <PersonalInfo />
        <NavigationBar class="nav-bar" />
      </div>
      <div class="right-col col">
        <section id="intro" data-section="intro">
          <AboutText class="intro" :class="{ active: isActive }" />
        </section>
        <br><br>
        <br><br>
        <br><br>
        <br><br>
        <br><br>
        <br><br>
        <section id="cv" data-section="cv">
          <h1>Curriculum Vitae</h1>
          <CurriculumVitae />
        </section>
      </div>
    </div>
    <!-- <div class="intro-box">
      <IntroductionBanner class="intro" :class="{ active: isActive }" />
    </div> -->
    <!-- <BodyGrid :items="bodyItems" class="bodyGrid" /> -->
  </main>
</template>

<style scoped>
.nav-bar {
  margin-top: 2em;
}

.flex-container {
  display: flex;
  flex-direction: row;
  margin-top: 3em;
  justify-content: space-between;
  height: 100vh;
}

.left-col {
  max-width: 30rem;
  position: sticky;
  top: 1em;
  left: 0;
}

.col {
  padding: 2em;
  width: 50%;
}

main {
  width: 100%;
}

.hide {
  opacity: 0;
  transition:
    opacity 0.5s ease-in-out,
    transform 0.5s ease-in-out;
  transform: translateY(-30%);
}

.intro {
  opacity: 0;
  transform: translateY(-30%);
  transition:
    opacity 0.3s ease-in-out,
    transform 0.3s ease-in-out;
}

.intro.active {
  opacity: 1;
  transform: translateY(0%);
}

.intro-box {
  display: flex;
  place-items: center;
  /* min-height: calc(100vh - 5em); */
}

.down-nav-container {
  width: 100%;
  height: 5em;
  text-align: center;
}
</style>
