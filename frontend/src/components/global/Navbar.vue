<script setup lang="ts">
import { ref } from "vue";
import { RouterLink } from 'vue-router'
import Image from 'primevue/image';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'

// Path for light and dark theme
const darkThemePath = '/themes/aura-dark-blue/theme.css';
const lightThemePath = '/themes/aura-light-blue/theme.css';

// Ref to follow the current theme
const currentThemePath = ref(lightThemePath);

const visible = ref<boolean>(false)

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
        <div class="pop_up_container">
            <Button class="pop_up_addr" @click="visible = true" label="Configuration"/>
            <Dialog v-model:visible="visible" modal header="Edit configuration" :style="{ width: '30rem' }">
                <span class="p-text-secondary block mb-5">...</span>
                <div class="flex align-items-center space-around gap-3 mb-3">
                    <label for="username" class="font-semibold w-6rem">Backend API</label>
                    <InputText id="username" class="flex-auto" autocomplete="off" />
                </div>
                <div class="flex align-items-center gap-3 mb-5">
                    <label for="email" class="font-semibold w-6rem">Node HTTP endpoint</label>
                    <InputText id="email" class="flex-auto" autocomplete="off" />
                </div>
                <div class="flex align-items-center gap-3 mb-5">
                    <label for="email" class="font-semibold w-6rem">Node WS endpoint</label>
                    <InputText id="email" class="flex-auto" autocomplete="off" />
                </div>
                <div class="flex justify-content-end gap-2">
                    <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                    <Button type="button" label="Save" @click="visible = false"></Button>
                </div>
            </Dialog>
        </div>
        <div class="ethereum_logo">
            <h3 class="application_title">Ether-faucet</h3>
            <RouterLink class="home_link" to="/">
                <Image src="/Ethereum_logo.png" alt="Image" width="40" />
            </RouterLink>
        </div>
        <div class="dark_mode_button">
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

.pop_up_container{
    width: 10rem;
}

.pop_up_addr{
    margin-top: 0rem;
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

.dark_mode_button{
    width: 10rem;
    display: flex;
    flex-direction: column;
    align-items: center;
}

</style>
