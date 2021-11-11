import { createRouter, createWebHistory } from '@ionic/vue-router';
import { RouteRecordRaw } from 'vue-router';
import AdminDashboard from '../views/AdminDashboard.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/admin/dashboard/stats'
  },
  {
    path: '/admin/dashboard/',
    component: AdminDashboard,
    children: [
      {
        path: '',
        redirect: '/admin/dashboard/stats'
      },
      {
        path: 'stats',
        component: () => import('@/views/Stats.vue')
      },
      {
        path: 'charts',
        component: () => import('@/views/Charts.vue')
      },
      {
        path: 'orders',
        component: () => import('@/views/Orders.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
