<script setup lang="ts">
import { ref } from "vue";
import { RouterLink } from 'vue-router'
import Image from 'primevue/image';
import Button from 'primevue/button';
import { usePrimeVue } from 'primevue/config';


// function toggleTheme(){
//     const PrimeVue = usePrimeVue();
//     PrimeVue.changeTheme('aura-dark-blue', 'aura-light-blue', 'theme-link', () => {});
// }

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
        themeLink.href = currentThemePath.value;
    }
}

</script>


<template>
    <div class="navbar">
        <div>
            <h3 >Ether-faucet</h3>
        </div>
        <div class="ethereum_logo">
            <RouterLink class="home_link" to="/">
                <Image src="/Ethereum_logo_.png" alt="Image" width="40" />
            </RouterLink>
            </div>
        <div>
            <Button icon="pi pi-moon" @click="toggleTheme" label="Dark Mode" />
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

.home_link{
    text-decoration:none;
}
</style>
