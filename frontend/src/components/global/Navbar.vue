<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref } from "vue";
import Menubar from 'primevue/menubar';
import Badge from 'primevue/badge';
import InputText from 'primevue/inputtext';
import Avatar from 'primevue/avatar';
import Image from 'primevue/image';


const items = ref([
    {
        label: 'Home',
        icon: 'pi pi-home'
    },
    {
        label: 'Features',
        icon: 'pi pi-star'
    },
    {
        label: 'Projects',
        icon: 'pi pi-search',
        items: [
            {
                label: 'Core',
                icon: 'pi pi-bolt',
                shortcut: '⌘+S'
            },
            {
                label: 'Blocks',
                icon: 'pi pi-server',
                shortcut: '⌘+B'
            },
            {
                label: 'UI Kit',
                icon: 'pi pi-pencil',
                shortcut: '⌘+U'
            },
            {
                separator: true
            },
        ]
    },
]);
</script>


<template>
    <div class="card">
        <Menubar :model="items" class="container_navbar">
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
                </a>
            </template>
            <template #end>
                DARK MODE
            </template>
        </Menubar>
    </div>
</template>

<style>
.container_navbar{
    display: flex;
    /* flex-direction: column; */
    justify-content: space-between;
}
</style>
