import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    notifier: {
      show: false,
      type: 'info',
      msg: ''
    },
    modal: {
      show: false,
      content: ''
    }
  },
  mutations: {
    toggleNotifier (state, payload) {
      state.notifier.show = payload.show
      state.notifier.type = payload.type
      state.notifier.msg = payload.msg
    },
    toggleModal (state, payload) {
      state.modal.show = payload.show
      state.modal.content = payload.content
    }
  },
  actions: {

  }
})
