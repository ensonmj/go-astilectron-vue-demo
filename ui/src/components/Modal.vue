<template>
  <div class="modal" id="modal" v-show="modal.show">
    <div class="modal-background"></div>
    <div class="modal-table">
      <div class="modal-wrapper">
        <div class="modal-body">
          <font-awesome-icon icon="times" @click="onClose"/>
          <div id="modal-content" v-html="modal.content"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    modal () {
      return this.$store.state.modal
    }
  },
  methods: {
    onClose () {
      this.$store.commit('toggleModal', {
        show: false,
        content: ''
      })
    }
  }
}
</script>

<style scoped>
.modal {
    color: #fff;
    height: 100%;
    left: 0;
    position: fixed;
    top: 0;
    width: 100%;
    z-index: 1;
}
.modal-background {
    background-color: #000;
    height: 100%;
    left: 0;
    opacity: 0.7;
    position: absolute;
    top: 0;
    width: 100%;
}
.modal-table {
    display: table;
    height: 100%;
    left: 0;
    position: absolute;
    top: 0;
    width: 100%;
}
.modal-wrapper {
    display: table-cell;
    height: 100%;
    padding: 2em;
    text-align: center;
    vertical-align: middle;
    width: 100%;
}
.modal-body {
    background-color: #f1f1f1;
    border-radius: 5px;
    color: #333;
    padding: 2em;
    position: relative;
}
.modal-close {
    cursor: pointer;
    position: absolute;
    right: 0.65em;
    top: 0.65em;
}
</style>
