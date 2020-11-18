import Vue from 'vue'
import VueRouter from 'vue-router'
import Parte1 from "./components/Parte1.vue";
import Parte2 from "./components/Parte2.vue";
import Parte3 from "./components/Parte3.vue";
Vue.use(VueRouter)

const routes = [
    { path: '/parte1', component: Parte1 },
    { path: '/parte2', component: Parte2 },
    { path: '/parte3', component: Parte3 },
]

const router = new VueRouter({
    mode: 'history',
    routes // short for `routes: routes`
})

export default router