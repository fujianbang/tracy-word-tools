export default [
    {
        path: '/',
        name: 'home',
        component: () => import('@/views/home')
    },
    {
        path: '/handle',
        name: 'handle',
        component: () => import('@/views/handle')
    }
]
