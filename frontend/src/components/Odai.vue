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
          <p class="text-center mt-3">候補一覧</p>
          <div v-for="(word, idx) in suggestionsRef" :key="idx" class="border mx-auto text-center" style="width: 50%;">
            <div>{{ word }}</div>
          </div>
          <div class="text-center">
            <a v-if="suggestionsRef" target="_blank" href="http://www.goo.ne.jp/">
              <img src="//u.xgoo.jp/img/sgoo.png" alt="supported by goo" title="supported by goo" style="width: 15%;">
            </a>
          </div>
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
import { defineComponent, inject, ref, useRouter, onMounted, watch, reactive, useContext } from '@nuxtjs/composition-api'
import axios from 'axios'
import { key } from '../utils/store'
import { STATE, SET } from '../utils/socket'

export default defineComponent({
  name: 'OdaiCom',
  setup () {
    const { app } = useContext()
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
    const suggestionsRef = ref<string[]>([])
    onMounted(() => {
      thema.value = store.data.thema
    })

    const toHiragana = (s: string) => {
      return new Promise((resolve) => {
        axios.post('https://labs.goo.ne.jp/api/hiragana', { app_id: app.$config.apiKey, sentence: s, output_type: 'hiragana' }).then((response) => {
          resolve(response.data.converted)
        })
      })
    }

    watch(data, () => {
      if (data.game_state === STATE.GameStart) {
        waiting.value = false
        console.log('start')
        router.push({ name: 'play' })
      }
    }, { deep: true })
    watch(wordRef, () => {
      wordValid.value = wordRef.value.match(/^[ぁ-んー　]{4,6}$/) == null
      if (wordRef.value !== '') {
        axios.get(`${location.protocol}//${location.host}/api/suggest/search?hl=ja&q=${wordRef.value}&output=toolbar`, { responseType: 'document' }).then((response) => {
          const xml = response.data as XMLDocument
          const suggestions = xml.querySelectorAll('suggestion')
          let suggestionValues: string[] = []
          suggestionsRef.value = []
          for (let i = 0; i < suggestions.length; i++) {
            if (suggestions[i].getAttribute('data') == null) {
              continue
            }
            if (suggestions[i].getAttribute('data')!.match(/.* .*/) !== null) {
              continue
            }
            suggestionValues.push(suggestions[i].getAttribute('data') as string)
          }
          suggestionValues = suggestionValues.sort((a, b) => a.length - b.length)
          for (let i = 0; i < suggestionValues.length; i++) {
            if (i > 3) {
              break
            }
            toHiragana(suggestionValues[i]).then(r => suggestionsRef.value.push(r as string))
          }
        })
      }
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
      waiting,
      suggestionsRef
    }
  }
})
</script>
