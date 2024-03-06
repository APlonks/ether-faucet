import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/wallets',
      name: 'wallets',
      component: () => import('../views/WalletsView.vue')
    },
    {
      path: '/SendEthersToSpecificAddress',
      name: 'SendEthersToSpecificAddress',
      component: () => import('../views/SendEthersToSpecificAddressView.vue')
    },
    {
      path: '/Simulation',
      name: 'Simulation',
      component: () => import('../views/SimulationView.vue')
    }
  ]
})

export default router
