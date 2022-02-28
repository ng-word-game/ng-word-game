<template>
  <div class="d-flex justify-content-center align-items-center flex-column" style="height: 100vh;">
    <div class="mb-3">
      <p class="text-center" style="font-size: 1.2rem; font-weight: bold;">
        {{ anotherUser.Name }}のワード
      </p>
      <SquareCom v-if="anotherUserCharInfo" :chara-info="anotherUserCharInfo" />
    </div>
    <div class="d-flex align-items-center justify-content-around" style="height: 40vh; width: 80%;">
      <div style="width: 30%;">
        <div class="card">
          <p class="text-center" style="font-size: 100%;">
            お題
          </p>
          <p class="text-center" style="font-size: 200%; font-weight: bold;">
            {{ thema }}
          </p>
        </div>
        <div class="my-2 d-flex justify-content-center">
          <b-button v-if="nextTurn === user.Id" v-b-modal.modal-1 variant="outline-info" style="width: 80%;">
            NGを選ぶ
          </b-button>
          <p v-else>
            相手を待っています...
          </p>
        </div>
      </div>
      <div class="card" style="height: 100%; width: 60%;">
        <div class="card-body">
          <h5 class="text-center">
            NGリスト
          </h5>
          <div class="table-responsive" style="height: 80%;">
            <table class="table table-sm">
              <thead>
                <tr>
                  <th class="text-center" scope="col">
                    名前
                  </th>
                  <th class="text-center" scope="col">
                    NG文字
                  </th>
                </tr>
              </thead>
              <tbody v-for="v in ngCharas" :key="v.name + v.char">
                <tr>
                  <td class="text-center">
                    {{ v.name }}
                  </td>
                  <td class="text-center">
                    {{ v.char }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
    <div class="mt-3">
      <p class="text-center" style="font-size: 1.2rem; font-weight: bold;">
        {{ user.Name }}のワード
      </p>
      <SquareCom v-if="userCharInfo" :chara-info="userCharInfo" />
    </div>
    <b-modal id="modal-1" ref="modalRef" hide-footer>
      <div class="d-block text-center">
        <h3>NG文字を入力しよう</h3>
      </div>
      <input v-model="ngChar" type="text" class="form-control" placeholder="">
      <div v-if="ngCharValid" class="form-text" style="color: red;">ひらがな1文字のみ入力できます</div>
      <b-button :disabled="ngCharValid" class="mt-2" variant="outline-info" block @click="registerNgChar">
        決定
      </b-button>
    </b-modal>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, ref, watch, onMounted, reactive, useRouter } from '@nuxtjs/composition-api'
import { key } from '../utils/store'
import { SET, WORDSTATE } from '../utils/socket'
import SquareCom from './parts/Square.vue'

export default defineComponent({
  name: 'PlayCom',
  components: {
    SquareCom
  },
  setup () {
    const store = inject(key)
    if (!store) {
      throw new Error('store undefined')
    }
    const router = useRouter()
    if (!store.state.socket) {
      router.push({ name: 'index' })
    }
    const user = store.data.users.filter(u => u.Name === store.state.name)[0]
    const anotherUser = store.data.users.filter(u => u.Name !== user.Name)[0]
    const ngChar = ref('')
    const ngCharValid = ref(true)
    const modalRef = ref()
    const data = reactive(store.data)
    const nextTurn = ref(store.data.next_turn)
    const userCharInfo = ref<{char: string, isOpen: boolean, isHide: boolean}[]>()
    const anotherUserCharInfo = ref<{char: string, isOpen: boolean, isHide: boolean}[]>()
    const thema = ref(store.data.thema)
    const ngCharas = ref(store.data.ng_chars)
    const registerNgChar = () => {
      modalRef.value.hide()
      if (!store.state.socket) {
        console.log('websocket is not found')
        return
      }
      store.state.socket.send(JSON.stringify({ type: SET.SetNgChar, ng_char: ngChar.value }))
    }

    onMounted(() => {
      setCharInfos()
    })
    watch(data, () => {
      setCharInfos()
      nextTurn.value = data.next_turn
      ngCharas.value = data.ng_chars
      console.log(ngCharas)
    }, { deep: true })

    watch(ngChar, () => {
      ngCharValid.value = ngChar.value.match(/^[ぁ-んー　]{1}$/) == null
    })
    // const charaInfo: CharaInfoInterface[] = [{
    //   char: 'こ',
    //   isOpen: false,
    //   isHide: false
    // }, { char: 'ん', isOpen: true, isHide: false }, { char: 'に', isOpen: false, isHide: false }, { char: 'ち', isOpen: false, isHide: true }, { char: 'は', isOpen: false, isHide: false }]

    const setCharInfos = () => {
      const usrCharInfo: {char: string, isOpen: boolean, isHide: boolean}[] = []
      store.data.word_state[user.Id].forEach((item) => {
        for (const c in item) {
          usrCharInfo.push({ char: c, isOpen: item[c] === WORDSTATE.OpenedWord, isHide: false })
        }
      })
      userCharInfo.value = usrCharInfo
      const anUsrCharInfo: {char: string, isOpen: boolean, isHide: boolean}[] = []
      store.data.word_state[anotherUser.Id].forEach((item) => {
        for (const c in item) {
          anUsrCharInfo.push({ char: c, isOpen: item[c] === WORDSTATE.OpenedWord, isHide: item[c] === WORDSTATE.HiddenWord })
        }
      })
      anotherUserCharInfo.value = anUsrCharInfo
    }

    return {
      user,
      userCharInfo,
      anotherUser,
      anotherUserCharInfo,
      nextTurn,
      ngCharas,
      ngChar,
      ngCharValid,
      modalRef,
      registerNgChar,
      thema
    }
  }
})
</script>
