<script setup lang="ts">
import Web3 from 'web3'
import axios from 'axios'
import { ref, onMounted } from 'vue';
// import { ProductService } from '@/service/ProductService';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import ColumnGroup from 'primevue/columngroup';   // optional
import Row from 'primevue/row';                   // optional
import { todo } from 'node:test';

const web3 = new Web3(import.meta.env.VITE_HTTP_ENDPOINT_NODE);
// console.log(await web3.eth.net.getId());

let todos = ref ()
let array: { publicKey: string; privateKey: string }[] = []
// let newWallet:{ publicKey: string; privateKey: string };

function createWallet(){
    const account = web3.eth.accounts.create();
    console.log("public key:",account.address)
    console.log("private key",account.privateKey)
    const newWallet = {publicKey: account.address, privateKey: account.privateKey}
    array.push(newWallet)
    console.log(array)
}


const products = ref();
const columns = [
    { field: 'private key', header: 'private key' },
    { field: 'public key', header: 'public key' }
];

</script>

<template>
    <br> <br> <br>
    <Button @click="createWallet" label="New Wallet" />
    <div class="card">
        <DataTable :value="products" tableStyle="max-width: 60rem">
            <Column v-for="col of columns" :key="col.field" :field="col.field" :header="col.header"></Column>
        </DataTable>
    </div>

</template>

<style>
</style>