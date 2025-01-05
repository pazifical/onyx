import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ReminderView from '@/views/ReminderView.vue'
import InfoView from '@/views/InfoView.vue'

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
    {
      path: '/onyx/info',
      name: 'info',
      component: InfoView,
    },
  ],
})

export default router
