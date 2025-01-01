import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ReminderView from '@/views/ReminderView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/:path*',
      name: 'directory',
      component: HomeView,
    },
    {
      path: '/onyx/reminders',
      name: 'reminders',
      component: ReminderView,
    },
  ],
})

export default router
