<script setup lang="ts">
import { ref } from 'vue';
import InputText from 'primevue/inputtext';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import faucetService from '../service/FaucetService';

const wallet_to_send = ref("")
const toast = useToast();
const reqReturn = ref("")

function submitForm (){
    faucetService.SendEthersToSpecificAddress(wallet_to_send.value).then(response => {
        console.log('Transaction successful', response);
        reqReturn.value = response.data.status
        console.log('Request return : '+ reqReturn.value)
    })
    .catch(error => {
        console.error('Error sending transaction', error);
    });
    toast.add({ severity: 'info', summary: 'Info', detail: '1 ETH sent', life: 3000 });
}

</script>

<template>

<div class="">
    <form @submit.prevent="submitForm" class="form-send-transaction">
        <label for="wallet_to_send">Choose the public address</label>
        <InputText id="wallet_to_send" v-model="wallet_to_send" aria-describedby="username-help" />
        <small id="username-help">Need to be in type common.address</small>
        <button type="submit">Send Transaction</button>
    </form>
    <Toast/>
</div>

</template>

<style>
.form-send-transaction{
    display: flex;
    flex-direction: column;
    gap: -2px;
    width: 500px;
}
</style>