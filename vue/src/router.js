import Vue from 'vue'
import Router from 'vue-router'

import Index from '@/components/index/Index'
import NarBar from './components/header/NavBar.vue'
import Home from './views/Home.vue'
import CommonFooter from '@/components/footer/CommonFooter'
import ManageHome from './views/ManageHome.vue'
import ArticlePage from './components/article/ArticlePage.vue'
import UpdateArticlePage from './components/article/UpdateArticlePage.vue'
import CreateArticlePage from './components/article/CreateArticlePage.vue'
import About from  './views/profile/index.vue'
import SignIn from  './components/user/SignIn.vue'
import SignUp from  './components/user/SignUp.vue'
Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
    path: '/',
    component: Index,
    children: [
    {
        path: '/',
        name: 'home',
        components: {
          header: NarBar,
          content: Home,
          footer: CommonFooter
        },
        meta: {
          title: 'yfblog的个人博客 | 野生数据科学家'
        }
    },
    {
      path: '/articlepage/:id',
      name: 'articlepage',
      components: {
        header: NarBar,
        content: ArticlePage,
        footer: CommonFooter
      },
        meta: {
          title: '阅读文章 | 野生数据科学家'
        }
    },
    {
      path: '/createarticlepage',
      name: 'createarticlepage',
      components: {
        header: NarBar,
        content: CreateArticlePage,
        footer: CommonFooter
      },
        meta: {
          title: '创建文章 | 野生数据科学家'
        }
    },
    {
      path: '/managehome',
      name: 'managehome',
      components: {
        header: NarBar,
        content: ManageHome,
        footer: CommonFooter
      },
      meta: {
        title: '管理文章 | 野生数据科学家'
      }
    },
    {
      path: '/updatearticlepage/:id',
      name: 'updatearticlepage',
      components: {
        header: NarBar,
        content: UpdateArticlePage,
        footer: CommonFooter
      },
      meta: {
        title: '管理文章 | 野生数据科学家'
      }
    },
    {
      path: '/about',
      name: 'about',
      components: {
        header: NarBar,
        content: About,
        footer: CommonFooter
      },
      meta: {
        title: '关于我 | 野生数据科学家'
      }
    },
    {
      path: '/signin',
      name: 'signin',
      components: {
        header: NarBar,
        content: SignIn,
        footer: CommonFooter
      },
      meta: {
        title: '登录 | 野生数据科学家'
      }
    },
    {
      path: '/signup',
      name: 'signup',
      components: {
        header: NarBar,
        content: SignUp,
        footer: CommonFooter
      },
      meta: {
        title: '注册 | 野生数据科学家'
      }
    }
  ]
}
]
})

