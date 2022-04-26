<template>
  <transition name="page">
    <nuxt />
  </transition>
</template>

<script lang="ts">
import { defineComponent, provide, inject, watch, ref, reactive, onDeactivated, onMounted } from '@nuxtjs/composition-api'
import store, { key } from '../utils/store'
import { STATE, SET } from '../utils/socket'

export default defineComponent({
  setup () {
    provide(key, store as any)
    const injectedStore = inject(key)
    if (!injectedStore) {
      throw new Error('store undefined')
    }
    const state = ref(injectedStore.state)
    const data = reactive(injectedStore.data)
    let once = true
    const pingFunc = () => {
      return setInterval(() => {
        if (injectedStore.state.socket !== null) {
          injectedStore.state.socket.send(JSON.stringify({ type: SET.PING }))
        }
      }, 10000)
    }
    // const gameState = ref(injectedStore.state.data.game_state)
    watch(data, () => {
      if (injectedStore.state.socket === null) {
        clearInterval(pingFunc())
        return
      }
      if (data.game_state === STATE.GameStop || data.game_state === STATE.GameEnd || data.game_state === STATE.GameDrawEnd) {
        injectedStore.state.socket.close()
      }
    }, { deep: true })

    onMounted(() => {
      if (once) {
        once = false
        console.log('here')
        pingFunc()
      }
    })

    onDeactivated(() => {
      if (state.value.socket) {
        state.value.socket.close()
      }
    })

    return {}
  }
})
</script>

<style>
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}
*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
}
</style>
