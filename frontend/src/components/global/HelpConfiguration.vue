<script setup lang="ts">
import { ref } from "vue";
import Sidebar from 'primevue/sidebar';
import Button from "primevue/button";

const visible = ref(false);
</script>

<template>
    <div class="help-config card flex justify-content-center">
        <Sidebar v-model:visible="visible" header="Help Configuration" class="w-full md:w-20rem lg:w-30rem">
            <h3 style="bold">How it works ?</h3>
            <p>Configuration is a prerequisite for using ether-faucet functions</p>
            <p>Your frontend will not communicate directly with your blockchain node. It will make requests to the API (backend), which will then create a client to communicate with your node.</p>
            <br>
            <h3 style="bold">Examples of values using default configuration should be :</h3>
            <p>Backend API Address: http://X.X.X.X:5002</p>
            <p>Node HTTP endpoint: http://X.X.X.X:8545</p>
            <p>Node WS endpoint: http://X.X.X.X:8546</p>
            <p>X.X.X.X corresponds to IP address like : 192.168.3.124</p>
            <br>
            <h3>To know</h3>
            <p>Currently only "Backend API Address" is required</p>
        </Sidebar>
        <Button outlined label="Help Configuration" @click="visible = true" />
    </div>
</template>

<style>

</style>