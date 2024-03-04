<script setup lang="ts">
import { ref } from 'vue';
import InputText from 'primevue/inputtext';
import walletService from '../service/WalletService';

const wallet_to_send = ref("")

function submitForm (){
    walletService.sendTransaction(wallet_to_send.value).then(response => {
          console.log('Transaction successful', response);
        })
        .catch(error => {
          console.error('Error sending transaction', error);
        });
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