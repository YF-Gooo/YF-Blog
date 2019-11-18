import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import ManageHome from './views/ManageHome.vue'
import ArticlePage from './components/article/ArticlePage.vue'
import UpdateArticlePage from './components/article/UpdateArticlePage.vue'
import CreateArticlePage from './components/article/CreateArticlePage.vue'
import ArticleManage from './components/article/ArticleManage.vue'
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
      path: '/articlepage/:id',
      name: 'articlepage',
      component: ArticlePage
    },
    {
      path: '/createarticlepage',
      name: 'createarticlepage',
      component: CreateArticlePage
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
      path: '/updatearticlepage/:id',
      name: 'updatearticlepage',
      component: UpdateArticlePage
    },
    {
      path: '/about',
      name: 'about',
      component: () => import(/* webpackChunkName: "about" */ './views/profile/index.vue')
    },
    {
      path: '/signin',
      name: 'signin',
      component: () => import(/* webpackChunkName: "signin" */ './components/user/SignIn.vue')
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import(/* webpackChunkName: "signup" */ './components/user/SignUp.vue')
    },
  ]
})
