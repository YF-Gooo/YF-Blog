import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)


export default new Vuex.Store({
  state: {
    kw:"",
    amkw:""
  },
  mutations: {
    changeSearchKW(state,kw){
      state.kw=kw
  },
    changeAMSearchKW(state,amkw){
      state.amkw=amkw
  },
}})
