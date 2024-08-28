<script setup>
import {ref } from 'vue'
import CVEntry from '@/components/CVEntry.vue';

const entries = ref([
    {
        from: "November 2016",
        to: "February 2023",
        where: "Tetra B.V",
        title: "Software Engineer",
        text: "As a software engineer at Tetra in Utrecht, I was mostly working on automating processes mainly using Python web applications."
    }
])
</script>
<template>
    <div>
        <CVEntry v-for="entry in entries" :data="entry"/>
    </div>
</template>