<script setup>
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router';

const route = useRoute();

const observer = ref(null)

const menuItems = ref([
    {
        to: '/#intro',
        title: 'Introduction',
        data: "intro"
    },
    {
        to: '/#cv',
        title: 'Curriculum vitae',
        data: "cv"
    }
])

onMounted(() => {
    // const sections = this.$refs.sectionsContainer.querySelectorAll('section');

    // observer = new IntersectionObserver(this.handleScroll, {
    //     root: null,
    //     threshold: 0.5, // Adjust threshold to determine visibility
    // });
    // sections.forEach((section) => this.observer.observe(section));
})
</script>

<template>
    <nav id="menu">
        <a v-for="item in menuItems" :key="item.to" :href="item.to" :data-section="item.data" class="menu-item" :class="{
            'active': route.fullPath === item.to
        }">
            {{ item.title }}
        </a>
    </nav>
</template>
<style scoped>
#menu {
    display: flex;
    flex-direction: column;
    gap: 1.5em;
}
.menu-item{
    width: 12rem;
}
.menu-item:after {
    display: block;
    content: '';
    border-bottom: solid 3px var(--vt-c-dark-pastel-blue);  
    transform: scaleX(0);  
    transition: transform 250ms ease-in-out;
}

.menu-item:not(.active):hover:after { transform: scaleX(1); }

.menu-item.active:after {
    border-bottom: 3px solid var(--vt-c-dark-pastel-blue);
    transform: scaleX(1);
}

</style>
