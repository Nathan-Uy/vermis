import { createRouter, createWebHistory } from "vue-router";
import Login from '../views/Login.vue';
import Dashboard from '../views/Dashboard.vue';
import AdminView from "../views/AdminView/AdminView.vue";
import UserView from "../views/UserView/UserView.vue";
import createUser from "../views/createUser.vue";
const routes = [
    { 
        path: '/', 
        name: 'Login',
        component: Login 
    },
    { 
        path: '/dashboard', 
        name: 'Dashboard',
        component: Dashboard 
    },
    {
        path: '/AdminView',
        name: 'AdminView',
        component: AdminView
    },
    {
        path: '/UserView',
        name: 'UserView',
        component: UserView
    },
    {
        path: '/createUser',
        name: 'createUser',
        component: createUser
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;