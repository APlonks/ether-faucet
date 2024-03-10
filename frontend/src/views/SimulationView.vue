<script setup lang="ts">
import CreateSimulation from '../components/actions/CreateSimulation.vue'
import HelpSimulation from '../components/actions/HelpSimulation.vue'

</script>


<template>
    <div class="container_simulation_view">
        <HelpSimulation/>
        <CreateSimulation/>
    </div>
</template>

<style>
.container_simulation_view{
    display: flex;
    flex-direction: column;
}

</style>