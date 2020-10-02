import Index from '@/components/pages/Index.vue'

export default [
    {
        path: '/',
        name: 'Index',
        redirect:to=>{
          return   "/histories"
        },
        component: Index
    },
    // {
    //     path:'*',
    //     redirect:'/'
    // },
    {
        path: '/articles/create',
        name: 'ArticleCreate',
        component: () => import('@/views/articles/Create')
    },
    {
        path: '/category/:category_id/articles',
        name: 'ArticleList',
        component: () => import('@/views/articles/List')
    },
    {
        path: '/category/:category_id/articles/:article_id/edit',
        name: 'ArticleEdit',
        component: () => import('@/views/articles/Edit')
    },
    {
        path: '/categories',
        name: 'CategoryList',
        component: () => import('@/views/categories/Index')
    },
    {
        path: '/category/create',
        name: 'CategoryCreate',
        component: () => import('@/views/categories/Create')
    },
    {
        path: '/category/:category_id/edit',
        name: 'CategoryEdit',
        component: () => import('@/views/categories/Edit')
    },
    {
        path: '/history/create',
        name: 'HistoryEdit',
        component: () => import('@/views/histories/Create')
    },
    {
        path: '/histories',
        name: 'HistoryList',
        component: () => import('@/views/histories/List')
    },
    {
        path: '/history/:history_id',
        name: 'HistoryShow',
        component: () => import('@/views/histories/Show')
    }
]