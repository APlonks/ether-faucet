<script setup lang="ts">
import { ref } from 'vue';
import InputText from 'primevue/inputtext';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import faucetService from '../../service/FaucetService';
import FloatLabel from 'primevue/floatlabel';
import Button from 'primevue/button';

const wallet_to_send = ref("")
const toast = useToast();
const reqReturn = ref("")

function submitForm (){
    faucetService.SendEthersToSpecificAddress(wallet_to_send.value).then(response => {
        if (response && 'data' in response) {
            console.log("The response:",response)
            reqReturn.value = response.data.message
            if (reqReturn.value == "Request sent to the backend") {
                toast.add({ severity: 'info', summary: 'Info', detail: '1 ETH sent', life: 3000});
                console.log("Request sent to the backend")
            } else if (reqReturn.value == "Public address format is not valid")  {
                toast.add({ severity: 'warn', summary: 'Error', detail: 'Public address format is not valid', life: 3000});
            } else {
                console.error("Problem with the http response from the backend")
                toast.add({ severity: 'error', summary: 'Error', detail: 'Problem with the http response from the backend', life: 3000});
            }
        } else {
            console.error("Response is undefined or not in expected format.");
            console.error("The Backend is propably not running");
            toast.add({ severity: 'error', summary: 'Error', detail: 'Request not sent to the backend (check console)', life: 3000});
        }
    }).catch(error => {
        console.error('Error sending transaction : check the file .env of the frontend to configure the backend ip');
        console.error(error)
        toast.add({ severity: 'error', summary: 'Error', detail: 'Not sent to the backend (check console)', life: 3000});
    });
}

</script>

<template>

<div class="faucet_form">
    <form @submit.prevent="submitForm" class="form-send-transaction">
        <FloatLabel class="faucet_float_label">
            <label for="wallet_to_send">Choose the public address</label>
            <InputText class="faucet_input" id="wallet_to_send" v-model="wallet_to_send" aria-describedby="username-help" />
        </FloatLabel>
        <Button class="faucet_button" label="Send Transaction" type="submit"/>
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
    margin-top: 2rem;
    max-width: 50%;
}

</style>