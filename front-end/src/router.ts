import { createRouter, createWebHistory } from 'vue-router'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import Home from './views/Home.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter: (to: any, _from: any) => {
      if (!localStorage.getItem('token') && to.name != 'Login') {
        return { name: 'Login' }
      }
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
