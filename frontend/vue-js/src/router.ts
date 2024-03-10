import { createRouter, createWebHistory } from 'vue-router'
import ListView from './components/ListView.vue'
import WebsocketView from './components/WebsocketView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: ListView
    },
    {
      path: '/ws',
      name: "ws",
      component: WebsocketView
    }
  ]
})

export default router
