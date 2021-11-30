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
    path: '/admin/bars/',
    component: () => import('@/views/SuperAdminBars.vue'),
  }

]

const router = createRouter({
  linkActiveClass: 'active',
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
