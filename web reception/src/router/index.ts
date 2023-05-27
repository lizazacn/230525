import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router';
import page1 from "../pages/page1.vue";
import page2 from "../pages/page2.vue";
import page3 from "../pages/page3.vue";
import page4 from "../pages/page4.vue";

const routes: RouteRecordRaw[] = [

    {
        meta: {
            title: '请输入姓名',
            permiss: '1',
        },
        path: '/',
        name: 'page1',
        component: page1,
    },
    {
        meta: {
            title: '请输入班级',
            permiss: '1',
        },
        path: '/page1',
        name: 'page2',
        component: page2,
    },
    {
        meta: {
            title: '请输入年龄',
            permiss: '1',
        },
        path: '/page2',
        name: 'page3',
        component: page3,
    },
    {
        meta: {
            title: '请输入性别',
            permiss: '1',
        },
        path: '/page3',
        name: 'page4',
        component: page4,
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
