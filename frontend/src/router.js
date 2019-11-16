import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import UserPage from './views/UserPage.vue'
import ArticleList from './components/ArticleList.vue'
import ArticlePage from './views/ArticlePage.vue'
import Editor from './components/Editor.vue'
import ArticleManage from './components/ArticleManage.vue'
import ManageHome from './views/ManageHome.vue'
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
      path: '/articlepage/:id',
      name: 'articlepage',
      component: ArticlePage
    },
    {
      path: '/editor',
      name: 'editor',
      component: Editor
    },
    {
      path: '/articlemanage',
      name: 'articlemanage',
      component: ArticleManage
    },
    {
      path: '/managehome',
      name: 'managehome',
      component: ManageHome
    },
    {
      path: '/about',
      name: 'about',
      component: () => import(/* webpackChunkName: "about" */ './views/profile/index.vue')
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
