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
    },
    {
        to: '/#skills',
        title: 'Skills',
        data: "skills"
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
    <nav class="menu">
        <a v-for="(item, index) in menuItems" :key="item.to" :href="item.to" :data-section="item.data" class="menu-item" :class="{
            'active': route.fullPath === item.to || index == 0
        }">
            {{ item.title }}
        </a>
    </nav>
</template>
<style scoped>
.menu {
    padding: 1rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1rem;
}
.menu-item{
    width: 12rem;
    text-align:center;
}
.menu-item:after {
    margin: 0 auto;
    display: block;
    content: '•';
    float:left;
    font-weight: bold;
    /* width: 5px; */
    /* border-radius: 10px; */
    /* border-bottom: solid 3px var(--vt-c-pastel-blue-dark);   */
    opacity: 0;
    transition: opacity 350ms ease-out;
}

.menu-item:not(.active):hover:after { opacity: 1; }

.menu-item.active:after {
    /* border-bottom: 3px solid var(--vt-c-pastel-blue-dark); */
    opacity: 1;
}

</style>
