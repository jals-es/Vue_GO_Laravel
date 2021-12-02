import { createRouter, createWebHistory } from 'vue-router'

const routes = [{
        path: '/',
        redirect: '/login',
        component: () =>
            import ('@/views/Home.vue')

    },
    {
        path: '/home',
        component: () =>
            import ('@/views/Home.vue')
    },
    {
        path: '/about',
        component: () =>
            import ('@/views/About.vue')
    },
    {
        path: '/login',
        loader: 'css-loader',
        component: () =>
            import ('@/views/Login.vue')
    },
    {
        path: '/register',
        loader: 'css-loader',
        component: () =>
            import ('@/views/Register.vue')
    },
    {
        path: '/admin/dashboard/',
        component: () =>
            import ('@/views/SuperAdminDashboard.vue'),
        children: [
            // {
            //   path: '',
            //   redirect: '/admin/dashboard/stats'
            // }
            // {
            //   path: 'stats',
            //   component: () => import('@/components/Stats.vue')
            // },
            // {
            //   path: 'charts',
            //   component: () => import('@/components/Charts.vue')
            // },
            // {
            //   path: 'orders',
            //   component: () => import('@/components/Orders.vue')
            // }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router