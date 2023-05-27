import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router';
import page1 from "../pages/card.vue";


const routes: RouteRecordRaw[] = [

    {
        meta: {
            title: '标签后台页面',
            permiss: '1',
        },
        path: '/',
        name: 'page1',
        component: page1,
    },

    {
        path: '/:catchAll(.*)',
        name: '403',
        meta: {
            title: '没有权限',
        },
        component: () => import(/* webpackChunkName: "403" */ '../views/403.vue'),
    },
];

const router = createRouter({
        history: createWebHistory(),
        routes,
        scrollBehavior() {
            return {left: 0, top: 0}
        }
    }
);

router.beforeEach((to, from, next) => {
    document.title = `${to.meta.title}`;
    next()

});



export default router;
