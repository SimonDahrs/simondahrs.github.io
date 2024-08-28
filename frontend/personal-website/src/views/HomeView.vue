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
  const container = document.querySelector(".flex-container");
  container?.addEventListener('scroll', (event) => {
    console.log(container.scrollTop, container.scrollHeight)
    const sections = document.querySelectorAll('section');
    const navLinks = document.querySelectorAll('#menu a');

    sections.forEach(section => {
      const sectionTop = section.offsetTop;
      const sectionHeight = section.offsetHeight;
      if (container.scrollTop + 50 >= sectionTop && container.scrollTop + 50 < sectionTop + sectionHeight) {
        navLinks.forEach(link => {

          if (link.getAttribute("data-section") === section.getAttribute("data-section")) {
            // console.log("activate", section.getAttribute("data-section"))
            link.classList.add('active');
          } else {
            // console.log("deactivate", section.getAttribute("data-section"))
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
        <br>
        <div class="divider"></div>
        <br>
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
h1 {
  font-weight: bold
}

.nav-bar {
  margin-top: 2em;
}

.flex-container {
  display: flex;
  flex-direction: row;
  margin-top: 50px;
  box-sizing: border-box;
  justify-content: center;
  height: calc(100vh - 5em);
  overflow-x: auto;
  scroll-behavior: smooth;
}

.left-col {
  max-width: 30rem;
  position: sticky;
  top: 0;
  left: 0;
}

.col {
  padding: 0 2em;
  width: 50%;
  max-width: 640px;
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

.divider {
  width: 100%;
  height: 10px;
  background:
    radial-gradient(circle at 100% 100%, #ffffff 0, #ffffff 3px, transparent 3px) 0% 0%/8px 8px no-repeat,
    radial-gradient(circle at 0 100%, #ffffff 0, #ffffff 3px, transparent 3px) 100% 0%/8px 8px no-repeat,
    radial-gradient(circle at 100% 0, #ffffff 0, #ffffff 3px, transparent 3px) 0% 100%/8px 8px no-repeat,
    radial-gradient(circle at 0 0, #ffffff 0, #ffffff 3px, transparent 3px) 100% 100%/8px 8px no-repeat,
    linear-gradient(#ffffff, #ffffff) 50% 50% / calc(100% - 10px) calc(100% - 16px) no-repeat,
    linear-gradient(#ffffff, #ffffff) 50% 50% / calc(100% - 16px) calc(100% - 10px) no-repeat,
    linear-gradient(145deg, transparent 50%, #dce5f7 100%);
  border-radius: 8px;
}
</style>
