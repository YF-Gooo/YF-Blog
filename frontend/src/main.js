import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueQuillEditor  from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import ElementUI from 'element-ui' //element-ui的全部组件
import 'element-ui/lib/theme-chalk/index.css'//element-ui的css
import axios from 'axios';
Vue.prototype.$axios = axios;
Vue.use(ElementUI) //使用elementUI

Vue.use(VueQuillEditor);
Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
