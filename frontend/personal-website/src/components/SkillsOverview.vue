<script setup>
import { ref } from 'vue';

const skills = ref(
    {
        "Languages": ["Python", "SQL", "Java", "Javascript", "HTML", "CSS"],
        "Web tech": ["FastAPI", "Flask", "Web2Py", "Vue.js"],
        "AI tech": ["PyTorch", "scikit-learn", "LangChain", "SciPy", "NumPy"],
        "Data tech": ["MySQL", "PostgreSQL", "SQL Server", "Redis", "MongoDB", "Celery", "Pandas", "RabbitMQ"],
        "DevOps": ["Docker", "CI/CD", "Kubernetes", "Ansible"],
        "Cloud": ["Azure"]
    }
)
</script>

<template>
    <div class="all-skills">
        <template v-for="(value, key) in skills" :key="key">
            <span class="category">{{ key }}</span>
            <div>
                <span class="skill" v-for="skill in value" :key="skill">{{ skill }}</span>
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
    padding: 8px 16px;
    font-size: 14px;
    background: linear-gradient(135deg, rgb(var(--vt-c-accent-light)) 25%, rgb(var(--vt-c-accent-extra-light)) 80%);
    border: 1px solid rgb(var(--vt-c-accent-extra-light));
    border-radius: 50px;
    box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);
    text-align: center;
    transition: all 0.3s ease;
    margin: 0 5px 7px 5px;
    cursor: pointer;
}

.skill:hover {
    background: linear-gradient(135deg, rgb(182, 192, 237) 25%, rgb(var(--vt-c-accent-light)) 80%);
  color: #FFFFFF; /* Change text color to white for contrast */
  box-shadow: 2px 4px 12px rgba(0, 0, 0, 0.15); /* Slightly deeper shadow on hover */
  transform: translateY(-2px); /* Subtle lift effect */
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