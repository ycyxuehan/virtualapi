import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    activedMenu:"default",
  },
  mutations: {
    activeMenu(state, menu){
      // console.info(menu)
      state.activedMenu = menu;
    }
  },
  actions: {
    ActiveMenu({commit}, menu){
      commit("activeMenu", menu)
    }
  },
  modules: {},
  getters:{
    GetActivedMenu: state => {state.activedMenu},
  }
});
