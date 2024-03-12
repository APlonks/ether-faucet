<script setup lang="ts">
import { onMounted, ref } from "vue";
import { RouterLink } from 'vue-router'
import Image from 'primevue/image';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import testingService from "@/service/TestingService";
import { error } from "console";

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


const api_addr = ref<string>('')
const http_endpoint = ref<string>('')
const ws_endpoint = ref<string>('')

const api_connected = ref<string>('Test')
const test_api_button_color = ref<string>('')


function save(){
    // Sauvegarde des valeurs dans localStorage
    localStorage.setItem('api_addr', api_addr.value);
    localStorage.setItem('http_endpoint', http_endpoint.value);
    localStorage.setItem('ws_endpoint', ws_endpoint.value);
}

// Envisager d'utiliser web3 au lieu d'un curl
function TestingConnectionAPI(){
    console.log("API addr:"+api_addr.value)
    testingService.TestingAPI(api_addr.value).then(response => {
        if (response && 'data' in response && response.data.message == "API connected") {
            api_connected.value = "up"
        }else{
            api_connected.value = "down"
        console.log("Can't connect to the API backend")
        }}).catch(error => {
            console.log("Problem with connection to the backend")
        })
        if (api_connected.value == "up"){
            console.log("Entering")
            test_api_button_color.value = "success"
        } else if (api_connected.value == "down") {
            console.log("Entering222")
            test_api_button_color.value = "error"
        } else {
            console.log("TOO FAST")
            console.log("api_connected: "+api_connected.value)
        }
}

// Voir si on peut faire un curl pour du ws sinon utiliser web3
function TestingConnectionHTTPEndpoint(){
}


// Voir si on peut faire un curl pour du ws sinon utiliser web3
function TestingConnectionWSEndpoint(){
}

// Chargement des valeurs sauvegardées au montage du composant
onMounted(() => {
  api_addr.value = localStorage.getItem('api_addr') || '';
  http_endpoint.value = localStorage.getItem('http_endpoint') || '';
  ws_endpoint.value = localStorage.getItem('ws_endpoint') || '';
});

</script>


<template>
    <div class="navbar">
        <div class="pop_up_container">
            <Button class="pop_up_addr" @click="visible = true" label="Configuration"/>
            <Dialog v-model:visible="visible" modal header="Edit configuration" :style="{ width: '40rem' }">
                <span class="p-text-secondary block mb-5">...</span>
                
                <div class="flex align-items-center gap-3 mb-3 label_input_container">
                    <label for="api_addr" class="font-semibold w-6rem">Backend API</label>
                    <div class="input_container">
                        <InputText id="api_addr" class="flex-auto" autocomplete="off" v-model="api_addr" />
                        <Button v-bind:severity="test_api_button_color" v-bind:label="api_connected" v-on:click="TestingConnectionAPI" />
                    </div>
                </div>
                <div class="flex align-items-center gap-3 mb-5 label_input_container">
                    <label for="http_endpoint" class="font-semibold w-6rem">Node HTTP endpoint</label>
                    <div class="input_container">
                        <InputText id="http_endpoint" class="flex-auto" autocomplete="off" v-model="http_endpoint"/>
                        <Button label="Test" />
                    </div>
                </div>
                <div class="flex align-items-center gap-3 mb-5 label_input_container">
                    <label for="ws_endpoint" class="font-semibold w-6rem">Node WS endpoint</label>
                    <div class="input_container">
                        <InputText id="ws_endpoint" class="flex-auto" autocomplete="off" v-model="ws_endpoint"/>
                        <Button label="Test" />
                    </div>
                </div>
                <div class="flex justify-content-end gap-2">
                    <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                    <Button type="button" label="Save" @click="visible = false, save()"></Button>
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

.label_input_container{
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
}

.input_container{
    width: 21rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
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
