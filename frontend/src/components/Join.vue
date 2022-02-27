<template>
  <div class="d-flex justify-content-center align-items-center" style="height: 100vh;">
    <div class="card" style="width: 50%;">
      <div class="card-body">
        <h5 class="card-title text-center">
          ワードゲーム
        </h5>
        <div class="form-group d-flex flex-column justify-content-center">
          <input v-model="name" type="email" class="form-control" placeholder="ユーザー名">
          <div class="mx-auto mt-3">
            <button v-if="!waiting" type="submit" class="btn btn-outline-info" :disabled="waiting" @click="registered">
              参加する
            </button>
            <p v-else>
              参加者を待っています....
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, reactive, ref, useRouter, watch } from '@nuxtjs/composition-api'
import { key } from '../utils/store'
import { STATE } from '../utils/socket'

export default defineComponent({
  name: 'PlayerJoin',
  setup () {
    const store = inject(key)
    const name = ref<string>('')
    const router = useRouter()
    if (!store) {
      throw new Error('store undefined')
    }
    const data = reactive(store.data)
    const waiting = ref(false)

    watch(data, () => {
      console.log('here')
      if (data.game_state === STATE.PlayerFilled) {
        waiting.value = false
        router.push({ name: 'odai' })
      }
    }, { deep: true })

    onMounted(() => {
      if (store.state.socket) {
        store.state.socket.addEventListener('close', () => {
          store.setSocket(null)
          console.log('close')
        })
        store.state.socket.close()
      }
    })

    const registered = () => {
      store.setName(name.value)
      waiting.value = true
      const socket = new WebSocket(`${location.protocol}//${location.hostname}:8080/?name=` + name.value)
      console.log(socket)
      socket.addEventListener('open', () => {
        console.log('open')
        store.setSocket(socket)
      })
    }

    return {
      name,
      registered,
      waiting
    }
  }
})
</script>
