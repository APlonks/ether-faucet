<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Navbar from './components/global/Navbar.vue'
import AccessWallets from './components/access/AccessWallets.vue'
import AccessSendEthersToSpecificAddress from './components/access/AccessSendEthersToSpecificAddress.vue';
import Faucet from './components/access/AccessSimulation.vue'

</script>

<template>
  <header>
    <div class="wrapper">
      <Navbar/>

      <nav>
        <AccessWallets/>
        <AccessSendEthersToSpecificAddress/>
        <Faucet/>
      </nav>
      <RouterView class="RouterView" />
    </div>
  </header>
  
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
  display: flex;
  justify-content: center;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

.RouterView {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%; 
  margin-top: 3em;
}

nav * {
  margin: 1rem;
}

/* @media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    flex-direction: column;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
} */
</style>
