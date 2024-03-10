<script setup>
import { ref } from "vue";
import Sidebar from 'primevue/sidebar';
import Button from "primevue/button";

const visible = ref(false);
</script>

<template>
    <div class="help-simu card flex justify-content-center">
        <Sidebar v-model:visible="visible" header="Sidebar" class="w-full md:w-20rem lg:w-30rem">
            <h3 style="bold">How it works ?</h3>
            <p>Click on "Start Simulation" to create two wallet groups (with the same number of wallets).</p> 
            <p>Each wallets in these groups will receive the same number of ethers at start-up.</p>
            <p>These wallets will then continuously send transactions.</p>
            <p>Clicking on "Stop Simulation" will send a signal to the backend to stop the simulation.</p>
            <p>You must wait a few seconds before the pop-up window (top right) confirms that the simulation has stopped.</p>
            <br>
            <h3 style="bold">Default values are :</h3>
            <p>Accounts per wallets : 1</p>
            <p>Ethers per transactions : 0</p>
            <p>Transactions per blocks : 2</p>
            <br>
            <h3>To know</h3>
            <p>Each group of wallets will send the number of transactions to each block. This number is half the total configurable number of transactions.</p>
            <p>Then choose a pair number or a transaction will be removed from the total.</p>
        </Sidebar>
        <Button label="Help" @click="visible = true" />
    </div>
</template>

<style>
.help-simu{
    margin-bottom: 1rem;
}

</style>