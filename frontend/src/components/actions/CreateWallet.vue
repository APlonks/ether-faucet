<script setup lang="ts">
import Web3 from 'web3';
import { ref } from 'vue';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

type Wallet = {
    publicKey: string;
    privateKey: string;
};

let wallets = ref<Wallet[]>([]);

function createWallet() {
    const httpEndpoint = localStorage.getItem('http_endpoint') || "http://127.0.0.1"
    const web3 = new Web3(httpEndpoint);
    const account = web3.eth.accounts.create();
    console.log("Public Key:", account.address);
    console.log("Private Key:", account.privateKey);
    const newWallet: Wallet = { publicKey: account.address, privateKey: account.privateKey };
    wallets.value.push(newWallet);
}

const columns = [
    { field: 'publicKey', header: 'Public Key' },
    { field: 'privateKey', header: 'Private Key' }
];
</script>

<template>
    <div class="container">
        <Button @click="createWallet" label="New Wallet" outlined />
        <br>
        <div class="card">
            <DataTable :value="wallets" :tableStyle="{ minWidth: '60rem' }">
                <Column v-for="col in columns" :key="col.field" :field="col.field" :header="col.header"></Column>
            </DataTable>
        </div>
    </div>
</template>

<style>
.container{
    display: flex;
    flex-direction: column;
}
</style>
