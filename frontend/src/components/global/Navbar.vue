<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref } from "vue";
import Image from 'primevue/image';


const items = ref([
    
]);
</script>


<template>
    <div class="navbar">
        <div>
            <h3 >Ether-faucet</h3>
        </div>
        <div class="ethereum_logo">
            <RouterLink class="home_link" to="/">
                <Image src="../../../public/Ethereum_logo_.png" alt="Image" width="40" />
            </RouterLink>
            </div>
        <div>
            DARK MODE
        </div>

        <!-- <Menubar :model="items" class="container_navbar">
            <template #start>
                <RouterLink to="/">
                    <div class="card flex justify-content-center">
                        <Image src="../../../public/Ethereum_logo_.png" alt="Image" width="30" />
                    </div>
                </RouterLink>
            </template>
            <template #item="{ item, props, hasSubmenu, root }">
                <a class="flex align-items-center" v-bind="props.action">
                    <span :class="item.icon" />
                    <span class="ml-2">{{ item.label }}</span>
                    <Badge v-if="item.badge" :class="{ 'ml-auto': !root, 'ml-2': root }" :value="item.badge" />
                    <span v-if="item.shortcut" class="ml-auto border-1 surface-border border-round surface-100 text-xs p-1">{{ item.shortcut }}</span>
                    <i v-if="hasSubmenu" :class="['pi pi-angle-down', { 'pi-angle-down ml-2': root, 'pi-angle-right ml-auto': !root }]"></i>
                    azeaze
                </a>
            </template>
            <template #end>
                DARK MODE
            </template>
        </Menubar> -->
    </div>
</template>

<style>
.navbar{
    display: flex;
    /* flex-direction: column; */
    justify-content: space-between;
    margin: 1%;
}

.home_link{
    text-decoration:none;
}
</style>
