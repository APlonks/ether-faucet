<script setup lang="ts">

import {ref} from "vue"
import Button from 'primevue/button'
import InputNumber from "primevue/inputnumber";
import SimulationService from "@/service/SimulationService";
import Toast from "primevue/toast";
import { useToast } from 'primevue/usetoast';

const blocked = ref(false)
const accounts_per_wallet = ref<number>()
const ethers_per_wallet = ref<number>()
const ethers_per_transaction = ref<number>()
const transactions_per_block = ref<number>()
const reqReturn = ref("")
const toast = useToast();


function startSimulation(){
    SimulationService.StartSimulation(accounts_per_wallet.value ?? 1, ethers_per_wallet.value ?? 1, ethers_per_transaction.value ?? 0, transactions_per_block.value ?? 2)
    .then(response => {
        if (response && "data" in response) {
            console.log(response.data)
            reqReturn.value = response.data.message
            if (reqReturn.value == "Simulation already started"){
                toast.add({ severity: 'warn', summary: 'Warning', detail: 'Simulation already started', life: 3000});
            } else if (reqReturn.value == "Simulation started"){
                toast.add({ severity: 'info', summary: 'Success', detail: 'Simulation started', life: 3000});
            } else if (reqReturn.value == "Cannot start Simulation"){
                toast.add({ severity: 'info', summary: 'Success', detail: 'Cannot start Simulation, look at backend logs', life: 3000});
            }
        } else {
            console.error("Response is undefined or not in expected format.");
            console.error("The Backend is propably not running");
            toast.add({ severity: 'error', summary: 'Error', detail: 'Request not sent to the backend (check console)\n the backend is probably not running', life: 3000});
        }
    })
}

function stopSimulation(){
    SimulationService.StopSimulation().then(response =>{
        if (response && "data" in response){
            console.log(response.data)
            reqReturn.value = response.data.message
            if(reqReturn.value == "Simulation already stopped"){
                toast.add({ severity: 'warn', summary: 'Warning', detail: 'Simulation already stopped', life: 3000});
            }else if (reqReturn.value == "Simulation stopped"){
                toast.add({ severity: 'info', summary: 'Warning', detail: 'Simulation stopped', life: 3000});
            }
        }
    })
}
</script>

<template>
    <div class="container_simu">
        <div class="faucet_form">
            <div class="simu_form_item">
                <label for="stacked-buttons" class="font-bold block mb-2"> Wallets number per group </label>
                <InputNumber v-model="accounts_per_wallet" inputId="wallets-number-per-group" mode="decimal" showButtons :min="1" :max="100" />
            </div>
            <div class="simu_form_item">
                <label for="stacked-buttons" class="font-bold block mb-2"> Ethers per wallet </label>
                <InputNumber v-model="ethers_per_wallet" inputId="ethers-per-wallet" mode="decimal" showButtons :min="1" :max="100" :step="1" />
            </div>
            <div class="simu_form_item">
                <label for="stacked-buttons" class="font-bold block mb-2"> Ethers per transaction </label>
                <InputNumber v-model="ethers_per_transaction" inputId="ethers-per-transaction" mode="decimal" showButtons :min="0" :max="100" :step="0.001" />
            </div>
            <div class="simu_form_item">
                <label for="stacked-buttons" class="font-bold block mb-2"> Transactions per block </label>
                <InputNumber v-model="transactions_per_block" inputId="transactions-per-block" mode="decimal" showButtons :min="2" :max="100" :step="1"/>
            </div>
            <br>
            <div>
                <Button class="simu_button" label="Start Simulaion" outlined @click="startSimulation"/>
                <Button class="simu_button" label="Stop Simulation" outlined @click="stopSimulation"/>
            </div>
            <Toast/>
        </div>
    </div>
</template>

<style>
.simu_form_item{
    display: flex;
    flex-direction: column;
}

.simu_button{
    margin-inline: 1rem;
}

.container_simu{
    display: flex;
    flex-direction: column;
}

.faucet_form{
    display: flex;
    flex-direction: column;
}

.simu_help_button{
    margin-top: 1rem;
    width: 40%;
}

</style>