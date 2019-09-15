import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import UserPage from './views/UserPage.vue'
import ArticleList from './components/ArticleList.vue'
import Editor from './components/Editor.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/user',
      name: 'user',
      component: UserPage
    },
    {
      path: '/articlelist',
      name: 'articlelist',
      component: ArticleList
    },
    {
      path: '/editor',
      name: 'editor',
      component: Editor
    },
    {
      path: '/about',
      name: 'about',
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/signin',
      name: 'signin',
      component: () => import(/* webpackChunkName: "signin" */ './components/SignIn.vue')
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import(/* webpackChunkName: "signup" */ './components/SignUp.vue')
    },
  ]
})
