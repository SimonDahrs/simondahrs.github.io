<script setup>
defineProps({
    data: {
        type: Object,
        required: true
    }
})
</script>

<template>
    <div>
        <div class="header">
            <h3 class="job-title">{{ data.title }}</h3>
            <span style="text-align: end">{{ data.from }} - {{ data.to }}</span>
        </div>
        <div class="content" v-html="data.text">
        </div>
    </div>
</template>
<style scoped>
.header {
    display: flex;
    justify-content: space-between;
}

.job-title {
    font-weight: bold;
}

.header>* {
    width: 50%;
}

.content {
    margin-top: 1em;
}

@media only screen and (max-width: 768px) {
    .header {
        display: block;
    }

    .header>* {
        width: 100%;
    }
}
</style>