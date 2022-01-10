import {
  createRouter,
  createWebHistory
} from 'vue-router'
import Auth from '@/core/check.js'
import SuperAdminAuth from '@/core/checkSuperAdmin.js'
const routes = [{
    path: '/',
    redirect: '/login',
    component: () =>
      import('@/views/Home.vue')
  },
  {
    path: '/admin/dashboard/',
    component: () => import('@/views/SuperAdminDashboard.vue'),
    beforeEnter: SuperAdminAuth.checkSuperAdmin
  },
  {
    path: '/admin/messages/',
    component: () => import('@/views/SuperAdminMessagesManagement.vue'),
    beforeEnter: SuperAdminAuth.checkSuperAdmin
  },
  {
    path: '/admin/bars/',
    component: () => import('@/views/SuperAdminBars.vue'),
    beforeEnter: SuperAdminAuth.checkSuperAdmin
  },
  {
    path: '/admin/bars/:id',
    component: () => import('@/views/SuperAdminBarManagement.vue'),
    beforeEnter: SuperAdminAuth.checkSuperAdmin

  },
  {
    path: '/admin/incidences/',
    component: () => import('@/views/SuperAdminIncidencesManagement.vue'),
    beforeEnter: SuperAdminAuth.checkSuperAdmin
  },
  {
    path: '/register',
    component: () => import('@/views/Register')
  },
  {
    path: '/login',
    component: () => import('@/views/Login')
  },
  {
    path: '/bars',
    component: () =>
      import('@/views/YourBars'),
    beforeEnter: Auth.checkAdmin
  },
  {
    path: '/bars/createBars',
    component: () =>
      import('@/views/CreateBars'),
    beforeEnter: Auth.checkAdmin
  }
]

const router = createRouter({
  linkActiveClass: 'active',
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router