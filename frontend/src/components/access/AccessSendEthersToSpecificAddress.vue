<script setup lang="ts">
import Button from 'primevue/button';
import Card from 'primevue/card';
import Image from 'primevue/image';

</script>


<template>
    <Card class="card_simu" style="width: 15rem; overflow: hidden">
        <template #header>
            <Image src="/Ethereum_logo.png" alt="Image" width="30" />
        </template>
        <template #title>Faucet</template>
        <template #content>
            <p class="m-0">
                Send Ether 1 by 1 to a specific wallet</p>
        </template>
        <template #footer>
            <div class="flex gap-3 mt-1">
                <RouterLink to="/SendEthersToSpecificAddress">
                    <Button label="Go to Simu" class="w-full" />
                </RouterLink>
            </div>
        </template>
    </Card>
</template>

<style>

</style>