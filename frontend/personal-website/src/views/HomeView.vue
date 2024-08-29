<script setup lang="ts">
import AboutText from '@/components/AboutText.vue'
import CurriculumVitae from '@/components/CurriculumVitae.vue';
import NavigationBar from '@/components/NavigationBar.vue'
import PersonalInfo from '@/components/PersonalInfo.vue'
import SkillsOverview from '@/components/SkillsOverview.vue';
import { onBeforeMount, onBeforeUnmount, onMounted, ref } from 'vue'

const isActive = ref(false)

onBeforeMount(() => {
  isActive.value = false
})

onMounted(() => {
  setTimeout(() => {
    isActive.value = true
  }, 0)
  window.addEventListener('scroll', (event) => {
    const sections = document.querySelectorAll('section');
    const navLinks = document.querySelectorAll('#menu a');

    let scrolledSection = sections[0]

    console.log(window.scrollY + window.innerHeight, document.body.scrollHeight)

    if (window.scrollY + window.innerHeight == document.body.scrollHeight) {
      // Bottom reached, activate lowest item
      scrolledSection = sections[sections.length - 1]

    } else {
      sections.forEach((section, index) => {
        const sectionTop = section.offsetTop;
        const sectionHeight = section.offsetHeight;
        if (window.scrollY >= sectionTop && window.scrollY < sectionTop + sectionHeight) {
          scrolledSection = section;
        }
      });
    }
    navLinks.forEach(link => {
      if (link.getAttribute("data-section") === scrolledSection.getAttribute("data-section")) {
        history.replaceState(null, "", "#" + link.getAttribute("data-section"))
        link.classList.add('active');
      } else {
        link.classList.remove('active');
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
      <div class="right-col col" :class="{ active: isActive }">
        <section id="intro" data-section="intro">
          <h1>Introduction</h1>
          <AboutText class="intro" />
        </section>
        <br>
        <div class="divider"></div>
        <br>
        <section id="cv" data-section="cv">
          <h1>Curriculum Vitae</h1>
          <CurriculumVitae />
        </section>
        <br>
        <div class="divider"></div>
        <br>
        <section id="skills" data-section="skills">
          <h1>Skills</h1>
          <SkillsOverview />
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

main {
  width: auto;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.col {
  padding: 0.5rem;
}

.flex-container {
  width: 98rem;
  max-width: 65vw;
  margin: 0 auto;
  padding: 4rem 0 4rem 0;
  scroll-behavior: smooth;
  position: relative;
  display: flex;
}

.left-col {
  width: 30%;
  max-width: 30rem;
  position: fixed;
  margin-bottom: 2rem;
  overflow: hidden;
  overflow-y: auto;
}

.right-col {
  max-width: 55rem;
  width: 50%;
  margin: 0 0 0 auto;
  min-height: 60rem;
  overflow: hidden;
  opacity: 0;
  transform: translateY(-30%);
  transition:
    opacity 0.3s ease-in-out,
    transform 0.3s ease-in-out;
}

.hide {
  opacity: 0;
  transition:
    opacity 0.5s ease-in-out,
    transform 0.5s ease-in-out;
  transform: translateY(-30%);
}

.right-col.active {
  opacity: 1;
  transform: translateY(0%);
}

.divider {
  width: 100%;
  height: 3px;
  background:
    radial-gradient(circle at 100% 100%, #ffffff 0, #ffffff 3px, transparent 3px) 0% 0%/8px 8px no-repeat,
    radial-gradient(circle at 0 100%, #ffffff 0, #ffffff 3px, transparent 3px) 100% 0%/8px 8px no-repeat,
    radial-gradient(circle at 100% 0, #ffffff 0, #ffffff 3px, transparent 3px) 0% 100%/8px 8px no-repeat,
    radial-gradient(circle at 0 0, #ffffff 0, #ffffff 3px, transparent 3px) 100% 100%/8px 8px no-repeat,
    linear-gradient(#ffffff, #ffffff) 50% 50% / calc(100% - 10px) calc(100% - 16px) no-repeat,
    linear-gradient(#ffffff, #ffffff) 50% 50% / calc(100% - 16px) calc(100% - 10px) no-repeat,
    linear-gradient(145deg, transparent 0%, #dce5f7 100%);
  border-radius: 8px;
}


@media only screen and (max-width: 768px) {

  /* For mobile phones: */
  .flex-container {
    flex-direction: column;
    margin-top: 1rem;
    padding-top: 0;
  }

  .left-col {
    position: relative;
    width: 100%;
    margin: 0 auto;
  }

  .right-col {
    width: 100%;
  }

  .nav-bar {
    display: none !important;
  }

}
</style>
