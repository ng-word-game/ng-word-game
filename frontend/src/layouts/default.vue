<template>
  <transition name="page">
    <nuxt />
  </transition>
</template>

<script lang="ts">
import { defineComponent, provide, inject, watch, ref, reactive, useRouter } from '@nuxtjs/composition-api'
import store, { key } from '../utils/store'
import { STATE } from '../utils/socket'
import { blobToJson } from '../utils/blobReader'

export default defineComponent({
  setup () {
    provide(key, store as any)
    const injectedStore = inject(key)
    if (!injectedStore) {
      throw new Error('store undefined')
    }
    const state = ref(injectedStore.state)
    const data = reactive(injectedStore.data)
    const once = ref(0)
    const router = useRouter()
    // const gameState = ref(injectedStore.state.data.game_state)
    watch(state, () => {
      if (state.value.socket && once.value === 0) {
        state.value.socket.addEventListener('message', (e) => {
          blobToJson(e.data).then((blobText) => {
            console.log(blobText)
            injectedStore.setData(blobText)
          })
        })
        once.value = 1
      }
    }, { deep: true })
    watch(data, () => {
      if (data.game_state === STATE.GameStop || data.game_state === STATE.GameEnd) {
        router.push({ name: 'result' })
      }
    }, { deep: true })
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
