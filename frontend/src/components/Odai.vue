<template>
  <div class="d-flex justify-content-center align-items-center" style="height: 100vh;">
    <div class="card" style="width: 50%;">
      <div class="card-body">
        <h5 class="card-title text-center">
          お題
        </h5>
        <h1 class="my-4 text-center">
          {{ thema }}
        </h1>
        <div class="form-group d-flex flex-column justify-content-center mt-4">
          <span class="form-text">4~6文字でひらがなのみ</span>
          <input v-model="wordRef" type="text" class="form-control" placeholder="ことばをかく">
          <div v-if="wordValid" class="form-text" style="color: red;">4~6文字でひらがなのみを使用してください</div>
          <div class="mx-auto mt-3">
            <button v-if="!waiting && wordValid" disabled type="submit" class="btn btn-outline-info" @click="registered">
              決定
            </button>
            <button v-else-if="!waiting" type="submit" class="btn btn-outline-info" @click="registered">
              決定
            </button>
            <p v-else>
              他のプレイヤーを待っています...
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, ref, useRouter, onMounted, watch, reactive } from '@nuxtjs/composition-api'
import { key } from '../utils/store'
import { STATE, SET } from '../utils/socket'

export default defineComponent({
  name: 'OdaiCom',
  setup () {
    const store = inject(key)
    const wordRef = ref<string>('')
    const wordValid = ref(false)
    const router = useRouter()
    const thema = ref('')
    const waiting = ref(false)
    if (!store) {
      throw new Error('store undefined')
    }
    const data = reactive(store.data)
    if (!store.state.socket) {
      router.push({ name: 'index' })
    }
    onMounted(() => {
      thema.value = store.data.thema
    })
    watch(data, () => {
      if (data.game_state === STATE.GameStart) {
        waiting.value = false
        console.log('start')
        router.push({ name: 'play' })
      }
    }, { deep: true })
    watch(wordRef, () => {
      wordValid.value = wordRef.value.match(/^[ぁ-んー　]{4,6}$/) == null
    })
    const registered = () => {
      store.setMyWord(wordRef.value)
      waiting.value = true
      if (!store.state.socket) {
        return
      }
      store.state.socket.send(JSON.stringify({ type: SET.SetWord, word: wordRef.value }))
    }

    return {
      wordRef,
      wordValid,
      thema,
      registered,
      waiting
    }
  }
})
</script>
