<script setup lang="ts">
import AboutText from '@/components/AboutText.vue'
import CurriculumVitae from '@/components/CurriculumVitae.vue';
import NavigationBar from '@/components/NavigationBar.vue'
import PersonalInfo from '@/components/PersonalInfo.vue'
import SkillsOverview from '@/components/SkillsOverview.vue';
import { onBeforeMount, onBeforeUnmount, onMounted, ref } from 'vue'

const isActive = ref(false)
const menuOpened = ref(false)

onBeforeMount(() => {
  isActive.value = false
})

onMounted(() => {
  setTimeout(() => {
    isActive.value = true
  }, 0)
  window.addEventListener('scroll', (event) => {
    const sections = document.querySelectorAll('section');
    const navLinks = document.querySelectorAll('.menu a');

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
      <div id="hamburger-menu">
        <div id="hamburger-toggle" @click="menuOpened = !menuOpened">
          <span class="button-icon" :class="{ opened: menuOpened }" id="openMenu">☰</span>
          <span class="button-icon" :class="{ opened: menuOpened }" id="closeMenu">✕</span>
        </div>
        <div id="hamburger-menu-item-container" :class="{ opened: menuOpened }">
          <NavigationBar />
        </div>
      </div>
      <div class="left-col col">
        <PersonalInfo />
        <NavigationBar id="nav-bar" />
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
#openMenu {
  opacity: 1;
  transition: transform .2s ease-in-out;
  transform: scale(1) rotate(0);
}

#openMenu.opened {
  transform: scale(0) rotate(90deg);
}

#closeMenu {
  transition: transform .2s ease-in-out;
  transform: scale(0) rotate(-90deg);
}

.button-icon {
  position: absolute;
  width: 50px;
  text-align: center;
}

#closeMenu.opened {
  opacity: 1;
  transform: scale(1) rotate(0);
}

#hamburger-menu {
  display: none;
  position: fixed;
  top: 2rem;
  left: 1rem;
  z-index: 5;
  display: none;
  row-gap: 1em;
}

#hamburger-toggle {
  font-weight: bold;
  font-size: 24px;
  box-sizing: border-box;
  background: rgba(199, 217, 254, 0.5);
  border-radius: 50%;
  display: flex;
  justify-content: center;
  height: 50px;
  width: 50px;
  text-align: center;
  align-items: center;
  cursor: pointer;
  transition: transform .3s ease-out;
  box-shadow:
    0.1px 0.1px 0.3px rgba(0, 0, 0, 0.02),
    0.1px 0.1px 0.8px rgba(0, 0, 0, 0.028),
    0.3px 0.3px 1.5px rgba(0, 0, 0, 0.035),
    0.4px 0.4px 2.7px rgba(0, 0, 0, 0.042),
    0.8px 0.8px 5px rgba(0, 0, 0, 0.05),
    2px 2px 12px rgba(0, 0, 0, 0.07);
}

#hamburger-toggle:hover {
  transform: scale(1.1);
  box-shadow:
    0.4px 0.4px 0.3px rgba(0, 0, 0, 0.02),
    1px 1px 0.8px rgba(0, 0, 0, 0.028),
    1.9px 1.9px 1.5px rgba(0, 0, 0, 0.035),
    3.4px 3.4px 2.7px rgba(0, 0, 0, 0.042),
    6.3px 6.3px 5px rgba(0, 0, 0, 0.05),
    15px 15px 12px rgba(0, 0, 0, 0.07);
}

#hamburger-menu-item-container {
  transform: scaleY(0);
  background-color: white;
  z-index: 5;
  box-shadow: 0 3px 10px rgb(0 1 1 / 0.3);
  border: 1px solid rgba(200, 200, 200, 0.3);
  border-radius: 7px;
  transition: transform .15s;
  transform-origin:  top;
}

#hamburger-menu-item-container.opened {
  transform: scaleY(1);
}

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

#nav-bar{
  margin-top: 3rem;
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
    float: right;
  }

  .right-col {
    width: 100%;
  }

  #nav-bar {
    display: none !important;
  }

  #hamburger-menu {
    display: grid;
  }

}
</style>
