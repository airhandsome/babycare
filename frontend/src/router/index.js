import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import PregnancyCare from '../views/PregnancyCare.vue'
import NewbornCare from '../views/NewbornCare.vue'
import PreschoolEducation from '../views/PreschoolEducation.vue'
import ExpertColumns from '../views/ExpertColumns.vue'

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
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router