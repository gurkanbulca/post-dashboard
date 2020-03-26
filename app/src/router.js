import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from "./components/Login"
import Homepage from "./components/Homepage"
import NewPost from "./components/NewPost"
import { store } from "./store"

Vue.use(VueRouter)

const routes = [
    {
        path: "/", component: Homepage
    },
    { path: "/login", component: Login, beforeEnter(to, from, next) {
        next(!store.getters.isAuthenticated ? true : "/login")
    } },
    { path: "/new-post", component: NewPost, beforeEnter(to, from, next) {
        next(store.getters.isAuthenticated?true:"/login")
    } }
]

export const router = new VueRouter({
    routes,
    mode: 'history'
})
