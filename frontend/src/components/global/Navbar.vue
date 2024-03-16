<script setup lang="ts">
import { ref } from "vue";
import { RouterLink } from 'vue-router'
import Configuration from './Configuration.vue'
import HelpConfiguration from './HelpConfiguration.vue'
import Image from 'primevue/image';
import Button from 'primevue/button';

</script>

<template>
    <div class="navbar">
        <div class="configuration_buttons">
            <Configuration/>
            <HelpConfiguration/>
        </div>
        <div class="ethereum_logo">
            <h3 class="application_title">Ether-faucet</h3>
            <RouterLink class="home_link" to="/">
                <Image src="/Ethereum_logo.png" alt="Image" width="40" />
            </RouterLink>
        </div>
        <div class="dark_mode_button_container">
            <Button label="Nothing :)" />
        </div>
    </div>
</template>

<style>
.navbar{
    display: flex;
    /* flex-direction: column; */
    justify-content: space-between;
    margin: 1%;
}

.navbar > div{
    width: 25rem;
}

.configuration_buttons{
    display: flex;
    justify-content: flex-start;
}

.ethereum_logo{
    display: flex;
    flex-direction: column;
    align-items: center;
}

.application_title{
    margin-top: 0rem;
}

.home_link{
    text-decoration:none;
}

.dark_mode_button_container{
    display: flex;
    flex-direction: column;
    align-items: flex-end;
}

</style>
