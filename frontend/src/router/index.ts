import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UsersQuiz from '../views/UsersQuizes.vue'
import LoginViewVue from '@/views/LoginView.vue'
import createQuizVue from '@/views/createQuiz.vue'
import signupViewVue from '@/views/signupView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },

    {
      path: '/quiz',
      name: 'quiz',
      component: UsersQuiz
    },
    {
      path: '/login',
      name: 'login',
      component: LoginViewVue
    },
    {
      path: '/quiz/new',
      name: 'quiznew',
      component: createQuizVue
    },
    {
      path: '/signup',
      name: 'signup',
      component: signupViewVue,
    },
  ]
})

export default router
