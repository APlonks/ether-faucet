<script setup lang="ts">
import { onMounted, ref } from "vue";
import Button from 'primevue/button';
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import InlineMessage from 'primevue/inlinemessage';
import testingService from "@/service/TestingService";


const visible = ref<boolean>(false)

const api_addr = ref<string>('')
const http_endpoint = ref<string>('')
const ws_endpoint = ref<string>('')

const api_addr_status = ref<string>('')
const http_endpoint_status = ref<string>('')
const ws_endpoint_status = ref<string>('')

function reset_status(){
    api_addr_status.value = "info"
    http_endpoint_status.value = "info"
    ws_endpoint_status.value = "info"
}

function reset_variables_addr(){
    api_addr.value =''
    http_endpoint.value=''
    ws_endpoint.value=''
    save()
}

function save(){
    // Save values in localStorage
    localStorage.setItem('api_addr', api_addr.value);
    localStorage.setItem('http_endpoint', http_endpoint.value);
    localStorage.setItem('ws_endpoint', ws_endpoint.value);
}

// Test connection to API
function TestingConnectionAPI(){
    if (api_addr.value == ""){
        api_addr_status.value = "error"
        console.log("No value has been entered")
    } else {
        testingService.TestingAPI(api_addr.value).then(async response => {
            if (response && 'data' in response && response.data.message == "API connected") {
                api_addr_status.value = "success"
            }else{
                api_addr_status.value = "error"
                console.log("Can't connect to the API backend")
            }
        }).catch(error => {
                console.log("Problem with connection to the backend")
        })
    }
}

// Test connection to the Node ethereum http endpoint
function TestingConnectionHTTPEndpoint(){
    const httpEndpointRegex = /^http:\/\/[0-9]+(?:\.[0-9]+){3}:[0-9]+$/;

    if (http_endpoint.value == ""){
        http_endpoint_status.value = "error"
        console.log("No value has been entered")
    } else if (!httpEndpointRegex.test(http_endpoint.value)) {
        http_endpoint_status.value = "error";
        console.log("HTTP endpoint is not in the correct format"); 
    } else {
        testingService.TestingHTTPEndpoint(http_endpoint.value).then(async response => {
        if (response && 'data' in response && response.status == 200) {
            console.log("response:",response)
            http_endpoint_status.value = "success"
        } else {
            http_endpoint_status.value = "error"
            console.log("Can't connect to the node HTTP endpoint")
        }}).catch(error => {
            http_endpoint_status.value = "error"
            console.log("Problem with connection to the HTTP endpoint")
        })
    }
}

// Load saved values when component mounted
onMounted(() => {
    api_addr.value = localStorage.getItem('api_addr') || '';
    http_endpoint.value = localStorage.getItem('http_endpoint') || '';
    ws_endpoint.value = localStorage.getItem('ws_endpoint') || '';
});

// Test connection to the Node ethereum ws endpoint
function TestingConnectionWSEndpoint(){
    const wsEndpointRegex = /^ws:\/\/[0-9]+(?:\.[0-9]+){3}:[0-9]+$/;

    if (ws_endpoint.value == ""){
            ws_endpoint_status.value = "error"
            console.log("No value has been entered")
    } else if (!wsEndpointRegex.test(ws_endpoint.value)) {
        ws_endpoint_status.value = "error";
        console.log("WS endpoint is not in the correct format"); 
    } else {
        testingService.TestingWSEndpoint(ws_endpoint.value).then(async response => {
        if (response == true){
            ws_endpoint_status.value = "success"
        } else {
            ws_endpoint_status.value = "error"
            console.log("Can't connect to the node WS endpoint")
        }}).catch(error => {
            ws_endpoint_status.value = "error"
            console.log("Problem with connection to the WS endpoint")
        })
    }
}

</script>

<template>
    <div class="pop_up_container">
        <Button class="pop_up_addr" @click="visible = true, reset_status()" label="Configuration"/>
        <Dialog v-model:visible="visible" modal header="Edit configuration" :style="{ width: '42rem' }">
            <span class="p-text-secondary block mb-5">...</span>
            
            <div class="flex align-items-center gap-3 mb-5 label_input_container">
                <label for="api_addr" class="font-semibold w-6rem">Backend API Address</label>
                <div class="input_container">
                    <InputText id="api_addr" class="flex-auto" autocomplete="off" v-model="api_addr" v-bind:placeholder="api_addr"/>
                    <InlineMessage v-bind:severity="api_addr_status"></InlineMessage>
                    <Button outlined label="Test" v-on:click="TestingConnectionAPI" />
                </div>
            </div>
            <div class="flex align-items-center gap-3 mb-5 label_input_container">
                <label for="http_endpoint" class="font-semibold w-6rem">Node HTTP endpoint</label>
                <div class="input_container">
                    <InputText id="http_endpoint" class="flex-auto" autocomplete="off" v-model="http_endpoint" v-bind:placeholder="http_endpoint"/>
                    <InlineMessage v-bind:severity="http_endpoint_status"></InlineMessage>
                    <Button outlined label="Test" v-on:click="TestingConnectionHTTPEndpoint"/>
                </div>
            </div>
            <div class="flex align-items-center gap-3 mb-5 label_input_container">
                <label for="ws_endpoint" class="font-semibold w-6rem">Node WS endpoint</label>
                <div class="input_container">
                    <InputText id="ws_endpoint" class="flex-auto" autocomplete="off" v-model="ws_endpoint" v-bind:placeholder="ws_endpoint"/>
                    <InlineMessage v-bind:severity="ws_endpoint_status"></InlineMessage>
                    <Button outlined label="Test" v-on:click="TestingConnectionWSEndpoint"/>
                </div>
            </div>
            <div class="buttons_cancel_reset_save flex justify-content-end gap-2">
                <div class="buttons_cancel_reset">
                    <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                    <Button type="button" label="Reset" @click="reset_variables_addr(), reset_status()"></Button>
                </div>
                <div class="button_save">
                    <Button type="button" label="Save" @click="visible = false, save()"></Button>
                </div>
            </div>
        </Dialog>
    </div>
</template>

<style>

.label_input_container{
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
}

.input_container{
    width: 25rem;
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

.buttons_cancel_reset_save{
    display: flex;
    justify-content: space-between;
}

.buttons_cancel_reset{
    width: 12rem;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.buttons_save{
    width: 10rem;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

</style>