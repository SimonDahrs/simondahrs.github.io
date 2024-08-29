<script setup>
import { ref } from 'vue';

const skills = ref(
    {
        "Languages": ["Python", "SQL", "Javascript", "HTML", "CSS"],
        "Web tech": ["FastAPI", "Flask", "Web2Py", "Vue.js"],
        "Data tech": ["MySQL", "PostgreSQL", "SQL Server", "Redis", "MongoDB", "Celery", "Pandas", "RabbitMQ"],
        "DevOps": ["Docker", "CI/CD", "Kubernetes", "Ansible"],
        "Cloud": ["Azure"]
    }
)
</script>

<template>
    <div class="all-skills">
        <template v-for="(value, key) in skills">
            <span class="category">{{ key }}</span>
            <div>
                <span class="skill" v-for="skill in value">{{ skill }}</span>
            </div>
        </template>
    </div>
</template>
<style scoped>
.all-skills {
    display: grid;
    grid-template-columns: 125px 75%;
    row-gap: 2rem;
}

.category {
    text-transform: uppercase;
    font-weight: bold;
}

.skill {
    display: inline-block;
    padding: 0.5rem;
    margin-left: 2px;
    margin-right: 2px;
    /* border: 1px solid transparent; */
    border-radius: 9999px;
    background: linear-gradient(135deg, var(--vt-c-pastel-blue-dark) 5%, var(--vt-c-pastel-blue-light) 80%, var(--vt-c-pastel-blue) 100%);
    font-size: 16px;
    font-weight: bold;
    margin-bottom: 3px;
}

@media only screen and (max-width: 768px) {
    .all-skills {
        grid-template-columns: 100%;
        row-gap: 0.8rem;
    }

    .category:not(:first-child) {
        margin-top: 2rem;
    }
}
</style>