<script setup>
import { ref } from 'vue'
import CVEntry from '@/components/CVEntry.vue';

const entries = ref([
    {
        from: "Jan 2024",
        to: "April 2024",
        where: "Apg",
        title: "AI for bears",
        text: "<p>Participating in the 'AI for Bears' challenge organized by FruitPunch AI. Goal of the project is twofold: helping ecologists doing research on bears, and preventing human-bear conflicts. Lead the team working on (re)identification of specific bears using deep (metric) learning, preferably on edge hardware. </p>"
    }
])
</script>
<template>
    <div>
        <div v-for="entry in entries" :key="entry.title">
            <CVEntry :data="entry" />
            <br><br>
        </div>
    </div>
</template>