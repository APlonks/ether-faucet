<script setup lang="ts">
import { ref } from "vue";
import { RouterLink } from 'vue-router'
import Configuration from './Configuration.vue'
import HelpConfiguration from './HelpConfiguration.vue'
import Image from 'primevue/image';
import Button from 'primevue/button';

// Path for light and dark theme
const darkThemePath = '/themes/aura-dark-blue/theme.css';
const lightThemePath = '/themes/aura-light-blue/theme.css';

// Ref to follow the current theme
const currentThemePath = ref(lightThemePath);

function toggleTheme() {
    // To change theme
    currentThemePath.value = currentThemePath.value === lightThemePath ? darkThemePath : lightThemePath;
  
    // Update <link> element with new item
    const themeLink = document.getElementById('theme-link');
    if (themeLink) {
        (themeLink as HTMLLinkElement).href = currentThemePath.value;
    }
}

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
            <Button @click="toggleTheme" label="Dark Mode" />
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
