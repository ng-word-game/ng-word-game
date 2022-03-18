<template>
  <div class="d-flex justify-content-center align-items-center flex-column" style="height: 100vh;">
    <div v-for="user in anotherUsers" class="mb-3">
      <p class="text-center" style="font-size: 1.2rem; font-weight: bold;">
        {{ user.Name }}のワード
      </p>
      <SquareCom v-if="user && anotherUsersCharInfo !== [] && anotherUsersCharInfo.filter(item => item.id === user.Id) !== [] && anotherUsersCharInfo.filter(item => item.id === user.Id)[0] && anotherUsersCharInfo.filter(item => item.id === user.Id)[0].chars" :chara-info="anotherUsersCharInfo.filter(item => item.id === user.Id)[0].chars" />
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
            文字を選ぶ
          </b-button>
          <p v-else>
            相手を待っています...
          </p>
        </div>
        <div class="d-flex">
          <div class="text-center" style="width: 50%;">
            <p>制限時間</p>
            <p v-if="nextTurn === user.Id">残り<span class="font-weight-bold" style="color: red;">{{ timer }}</span>秒</p>
          </div>
          <div class="text-center" style="width: 50%;">
            <p>ターン</p>
            <p class="font-weight-bold">{{ turn }}</p>
          </div>
        </div>
      </div>
      <div class="card" style="height: 100%; width: 60%;">
        <div class="card-body">
          <h5 class="text-center">
            NGリスト
          </h5>
          <div class="table-responsive" style="height: 80%;">
            <table class="table table-sm table-striped">
              <thead class="table-light" style="position: sticky; top: 0;">
                <tr>
                  <th class="text-center" scope="col">
                    名前
                  </th>
                  <th class="text-center" scope="col">
                    NG文字
                  </th>
                </tr>
              </thead>
              <tbody id="ng-table-body">
                <tr ref="tableRef" v-for="(v, idx) in ngCharas.filter(item => item.char !== 'L')" :key="idx">
                  <td class="text-center">
                    {{ v.name }}
                  </td>
                  <td v-if="v.char === ''" class="text-center">
                    時間切れ
                  </td>
                  <td v-else class="text-center">
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
      <div v-if="checkDuplicateNgChar" class="form-text" style="color: red;">すでにNG文字です</div>
      <b-button :disabled="ngCharValid" class="mt-2" variant="outline-info" block @click="registerNgChar">
        決定
      </b-button>
    </b-modal>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, ref, watch, onMounted, reactive, useRouter, useSlots } from '@nuxtjs/composition-api'
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
    const user = store.data.users.filter(u => u.Id === store.state.clientId)[0]
    const anotherUsers = store.data.users.filter(u => u.Id !== store.state.clientId)
    const ngChar = ref('')
    const ngCharValid = ref(true)
    const checkDuplicateNgChar = ref(false)
    const modalRef = ref()
    const data = reactive(store.data)
    const nextTurn = ref(store.data.next_turn)
    const userCharInfo = ref<{char: string, isOpen: boolean, isHide: boolean}[]>()
    const anotherUsersCharInfo = reactive<{id: string, chars: {char: string, isOpen: boolean, isHide: boolean}[]}[]>([])
    const thema = ref(store.data.thema)
    const ngCharas = ref(store.data.ng_chars)
    const timer = ref(0)
    const turn = ref(0)
    const tableRef = ref()
    const registerNgChar = () => {
      clearInterval(timerId.value)
      modalRef.value.hide()
      if (!store.state.socket) {
        console.log('websocket is not found')
        return
      }
      store.state.socket.send(JSON.stringify({ type: SET.SetNgChar, ng_char: ngChar.value }))
      ngChar.value = ''
    }

    const timerId = ref()

    const timerDown = () => {
      timer.value = 60
      const id = setInterval(() => {
        if (timer.value < 1) {
          ngChar.value = ''
          registerNgChar()
          return
        }
        timer.value -= 1
      }, 1000)
      timerId.value = id
    }

    onMounted(() => {
      setCharInfos()
      if (nextTurn.value === user.Id) {
        timerDown()
      }
    })
    watch(data, () => {
      setCharInfos()
      nextTurn.value = data.next_turn
      turn.value = data.turn
      ngCharas.value = data.ng_chars
      if (nextTurn.value === user.Id && checkLose()) {
        ngChar.value = 'L'
        registerNgChar()
        return
      }
      if (nextTurn.value === user.Id) {
        clearInterval(timerId.value)
        timerDown()
      }
    }, { deep: true })

    watch(tableRef, () => {
      const ngLast = document.querySelector('#ng-table-body > tr:last-child')
      ngLast!.scrollIntoView()
    })

    watch(ngChar, () => {
      ngCharValid.value = ngChar.value !== '' && ngChar.value.match(/^[ぁ-んー　]{1}$/) == null
      checkDuplicateNgChar.value = ngChar.value !== '' && ngCharas.value.filter(item => item.char === ngChar.value).length > 0
    })

    const checkLose = () => {
      if (!userCharInfo.value) {
        return false
      }
      return userCharInfo.value.filter(item => item.isOpen).length === userCharInfo.value.length
    }
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
      anotherUsers.forEach((user) => {
        const anUsrCharInfo: {char: string, isOpen: boolean, isHide: boolean}[] = []
        store.data.word_state[user.Id].forEach((item) => {
          for (const c in item) {
            anUsrCharInfo.push({ char: c, isOpen: item[c] === WORDSTATE.OpenedWord, isHide: item[c] === WORDSTATE.HiddenWord })
          }
        })
        if (anotherUsersCharInfo.filter(item => item.id === user.Id).length === 0) {
          anotherUsersCharInfo.push({ id: user.Id, chars: anUsrCharInfo })
        } else {
          anotherUsersCharInfo.filter(item => item.id === user.Id)[0].chars = anUsrCharInfo
        }
      })
    }

    return {
      user,
      userCharInfo,
      anotherUsers,
      anotherUsersCharInfo,
      nextTurn,
      ngCharas,
      ngChar,
      ngCharValid,
      checkDuplicateNgChar,
      modalRef,
      registerNgChar,
      thema,
      timer,
      turn,
      tableRef
    }
  }
})
</script>
