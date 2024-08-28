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
        <h3>{{ data.title }}</h3>
        <span style="text-align: end">{{ data.from }} - {{ data.to }}</span>
    </div>
    <div class="content">
        {{ data.text }}
    </div>
</div>
</template>
<style scoped>
.header {
    display: flex;
    justify-content: space-between;
    grid-template-columns: 50% 50%;
}
.header > * {
    width: 50%;
}
.content {
    margin-top: 1em;
}
</style>