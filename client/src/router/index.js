import { createRouter, createWebHistory } from 'vue-router'
const routes = [
  {
    path: '/',
    redirect: '/home',
    component: () => import('@/views/Home.vue')

  },
  {
    path: '/admin/dashboard/',
    component: () => import('@/views/SuperAdminDashboard.vue'),
  },
  {
    path: '/admin/messages/',
    component: () => import('@/views/SuperAdminMessagesManagement.vue'),
  },
  {
    path: '/admin/bars/',
    component: () => import('@/views/SuperAdminBars.vue'),
  },
  {
    path: '/admin/bars/:id',
    component: () => import('@/views/SuperAdminBarManagement.vue'),

  },
  {
    path: '/admin/incidences/',
    component: () => import('@/views/SuperAdminIncidencesManagement.vue'),
  },
  {
    path: '/register',
    component: () => import('@/views/Register')
  },
  {
    path:'/login',
    component: () => import('@/views/Login')
  }

]

const router = createRouter({
  linkActiveClass: 'active',
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
