import { createRouter, createWebHistory } from 'vue-router'
import Auth from '@/core/check.js'
const routes = [{
        path: '/',
        redirect: '/login',
        component: () =>
            import ('@/views/Home.vue')
    },
    {
        path: '/admin/dashboard/',
        component: () =>
            import ('@/views/SuperAdminDashboard.vue'),
    },
    {
        path: '/admin/bars/',
        component: () =>
            import ('@/views/SuperAdminBars.vue'),
    },
    {
        path: '/admin/bars/:id',
        component: () =>
            import ('@/views/SuperAdminBarManagement.vue'),

    },
    {
        path: '/admin/incidences/',
        component: () =>
            import ('@/views/SuperAdminIncidencesManagement.vue'),
    },
    {
        path: '/register',
        component: () =>
            import ('@/views/Register')
    },
    {
        path: '/login',
        component: () =>
            import ('@/views/Login')
    },
    {
        path: '/bars',
        component: () =>
            import ('@/views/YourBars'),
        beforeEnter: Auth.checkAdmin
    },
    {
        path: '/bars/createBars',
        component: () =>
            import ('@/views/CreateBars'),
        beforeEnter: Auth.checkAdmin
    }

]

const router = createRouter({
    linkActiveClass: 'active',
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router