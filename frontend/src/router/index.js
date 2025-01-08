import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import PregnancyCare from '../views/PregnancyCare.vue'
import NewbornCare from '../views/NewbornCare.vue'
import PreschoolEducation from '../views/PreschoolEducation.vue'
import ExpertColumns from '../views/ExpertColumns.vue'
import Login from '../views/auth/Login.vue'
import Register from '../views/auth/Register.vue'
import ForgotPassword from '../views/auth/ForgotPassword.vue'
import InfantGrowth from "../views/InfantGrowth.vue"
import CommunityForum from "../views/CommunityForum.vue"
import Privacy from '../views/Privacy.vue'
import Contact from '../views/Contact.vue'
import About from '../views/About.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/pregnancy-care',
    name: 'PregnancyCare',
    component: PregnancyCare
  },
  {
    path: '/newborn-care',
    name: 'NewbornCare',
    component: NewbornCare
  },
  {
    path: '/preschool',
    name: 'PreschoolEducation',
    component: PreschoolEducation
  },
  {
    path: '/expert-columns',
    name: 'ExpertColumns',
    component: ExpertColumns
  },
  {
    path: '/article/:id',
    name: 'ArticleDetail',
    component: () => import('../views/ArticleDetail.vue')
  },
  {
    path: '/privacy',
    name: 'Privacy',
    component: Privacy
  },
  {
    path: '/contact',
    name: 'Contact',
    component: Contact
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/infant-growth',
    name: 'InfantGrowth',
    component: InfantGrowth
  },
  {
    path: '/community',
    name: 'CommunityForum',
    component: CommunityForum
  },
  {
    path: '/auth/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/auth/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/auth/forgot-password',
    name: 'ForgotPassword',
    component: ForgotPassword
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router