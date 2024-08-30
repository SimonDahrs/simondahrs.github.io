<script setup lang="ts">
import { ref } from 'vue'
import CVEntry from '@/components/CVEntry.vue';

const entries = ref([
    {
        from: "Feb 2023",
        to: "Present",
        where: "Apg",
        title: "DevOps Engineer @ APG Asset Management",
        text: "<p>As a DevOps Engineer at APG Asset Management, a pension fund investment company, I:</p> <ul> <li>Develop and maintain applications and data solutions to support business users</li> <li>Design scalable systems, such as a Python application utilizing Redis task queues for efficient financial data retrieval</li> <li>Take lead in the internal Python community, organizing knowledge-sharing sessions and demos</li> <li>Collaborate with cross-functional teams to optimize operational workflows and implement best practices</li> </ul>"
    },
    {
        from: "Nov 2016",
        to: "Feb 2023",
        where: "Tetra B.V",
        title: "Software Engineer @ Tetra",
        text: "<p>At Tetra, I focused on designing and developing solutions to automate manual processes in pharmacies, primarily using Python web applications and microservices. Key achievements include:</p> <ul> <li>Designed and implemented a Single Sign-On (SSO) system for affiliated pharmacists using Flask, Vue.js, and MySQL, enhancing security and standardizing user authentication across multiple applications</li> <li>Built a flexible report generation system allowing parametric reports to be generated in Docker containers, improving data analysis capabilities</li> <li>Developed automated systems for medicine inventory management in the wholesale division, optimizing stock control and distribution</li> <li>Created Python web applications and microservices that interact with databases to streamline pharmacy operations</li></ul> "
    }
])
</script>
<template>
    <div>
        <div v-for="entry in entries">
            <CVEntry :data="entry" />
            <br><br>
        </div>
    </div>
</template>