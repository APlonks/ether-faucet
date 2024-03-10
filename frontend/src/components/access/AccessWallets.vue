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
        <template #title>Wallets</template>
        <template #content>
            <p class="m-0"> Create Wallets : public key + private key </p>
        </template>
        <template #footer>
            <div class="card_footer flex gap-3 mt-1">
                <RouterLink to="/wallets">
                    <Button label="Create Wallets" class="w-full" />
                </RouterLink>
            </div>
        </template>
    </Card>
</template>

<style>
.p-card-body{
    padding: 0rem;
}
</style>