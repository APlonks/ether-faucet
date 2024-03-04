<script setup lang="ts">
import { ref } from 'vue';
import SendTransaction from '../components/SendTransaction.vue';

const wallet_to_send = ref(null)

</script>

<template>
<SendTransaction/>
</template>

<style>

</style>