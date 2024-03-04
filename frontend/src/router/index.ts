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
      path: '/sendtransaction',
      name: 'sendtransaction',
      component: () => import('../views/SendTransactionView.vue')
    }
  ]
})

export default router
