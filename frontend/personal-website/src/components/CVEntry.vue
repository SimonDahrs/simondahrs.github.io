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
    <h3>{{ data.title }}</h3>
</div>
</template>