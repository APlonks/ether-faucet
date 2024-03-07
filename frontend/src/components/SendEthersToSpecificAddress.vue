<script setup lang="ts">
import { ref } from 'vue';
import InputText from 'primevue/inputtext';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import faucetService from '../service/FaucetService';
import FloatLabel from 'primevue/floatlabel';
import Button from 'primevue/button';

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

<div class="faucet_form">
    <form @submit.prevent="submitForm" class="form-send-transaction">
        <FloatLabel class="faucet_float_label">
            <label for="wallet_to_send">Choose the public address</label>
            <InputText class="faucet_input" id="wallet_to_send" v-model="wallet_to_send" aria-describedby="username-help" />
        </FloatLabel>
        <br>
        <Button class="faucet_button" label="Send Transaction" outlined type="submit"/>
    </form>
    <Toast/>
</div>

</template>

<style>
.form-send-transaction{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: -2px;
    width: 500px;
}

.faucet_float_label{
    min-width: 90%;
}

.faucet_input{
    min-width: 100%;
}

.faucet_button{
    max-width: 50%;
}

</style>