<script setup>
import { ref } from 'vue'
import CVEntry from '@/components/CVEntry.vue';

const entries = ref([
    {
        from: "Feb 2023",
        to: "Present",
        where: "Apg",
        title: "DevOps Engineer @ APG Asset Management",
        text: "APG Asset Management is a pension fund investment company. My role as DevOps Engineer is to support our business users by providing them with applications and data solutions they can use. One example is a Python application that uses a Redis task queue to fetch financial data from an external vendor. By doing so, the system can scale well for varying load. Apart from projects like these I am also involved in leading the internal Python community, for example by organizing demos and information sessions."
    },
    {
        from: "Nov 2016",
        to: "Feb 2023",
        where: "Tetra B.V",
        title: "Software Engineer @ Tetra",
        text: "As a software engineer at Tetra in Utrecht, I mostly designed and developed solutions to automate manual processes in pharmaces. Those solutions were often Python web applications or microservices that interact with databases. One specific example I am proud of is a Single Sign-On provider system used to authenticate and authorize users inside the company. Before that, most applications had separate login systems, but with this new system, everyone only has a single account. "
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