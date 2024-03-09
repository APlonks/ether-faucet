<script setup lang="ts">
import { ref } from 'vue';
import SendEthersToSpecificAddress from '../components/actions/SendEthersToSpecificAddress.vue';

const wallet_to_send = ref(null)

</script>

<template>
<SendEthersToSpecificAddress/>
</template>

<style>

</style>../components/actions/SendEthersToSpecificAddress.vue