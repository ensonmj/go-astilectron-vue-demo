<template>
  <div id="app">
    <Notifier/>
    <Modal/>
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script>
import Notifier from '@/components/Notifier.vue'
import Modal from '@/components/Modal.vue'
import { mapMutations } from 'vuex'

export default {
  components: {
    Notifier,
    Modal
  },
  mounted () {
    // Wait for astilectron to be ready
    document.addEventListener('astilectron-ready', this.onAstilectronReady)
  },
  methods: {
    onAstilectronReady: function () {
      // eslint-disable-next-line
      astilectron.onMessage((message) => {
        console.log(message)
        switch (message.name) {
          case 'menu.about':
            this.toggleModal({ show: true, content: message.payload })
            return { payload: 'payload' }
          case 'go.check.out.menu':
            this.toggleNotifier({ show: true, type: 'info', msg: message.payload })
            break
        }
      })
    },
    ...mapMutations(['toggleNotifier', 'toggleModal'])
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
  a {
    font-weight: bold;
    color: #2c3e50;
    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
