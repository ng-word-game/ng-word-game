<template>
  <div class="d-flex justify-content-center align-items-center" style="height: 100vh;">
    <div class="card" style="width: 75%;">
      <div class="card-body">
        <h4 class="card-title text-center">
          NG単語ゲーム
        </h4>
        <div class="form-group d-flex flex-column justify-content-center">
          <h5 class="text-center">ユーザー名</h5>
          <input v-model="name" type="text" class="form-control" placeholder="ユーザー名">
          <div v-if="nameErr" class="form-text" style="color: red;">ユーザー名を入力してください</div>
          <h5 class="mt-3 text-center">ルームを作成</h5>
          <div class="mx-auto form-inline justify-content-center">
            <label for="inputMaxPlayer">募集人数: </label>
            <input id="inputMaxPlayer" type="number" v-model="maxPlayer" class="mx-2 form-control" placeholder="人数" style="width: 25%;">
            <button v-if="!waiting" type="submit" class="btn btn-outline-info" :disabled="waiting" @click="createRoom">
              ルームを作成
            </button>
            <p v-else-if="!connectError">
              参加者を待っています....
            </p>
            <p v-if="connectError">サーバーと接続できません</p>
          </div>
          <div v-if="maxPlayerErr" class="form-text text-center" style="color: red;">募集人数は2人以上にしてください</div>
        </div>
        <h5 class="py-3 text-center">募集中のルーム</h5>
        <table class="table table-borderless">
          <thead>
            <tr>
              <th>参加ユーザー</th>
              <th class="text-center">参加人数/募集人数</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="room in rooms" v-bind:key="room.id">
              <td>{{ room.players.join(', ') }}</td>
              <td class="text-center">{{ room.num }}/{{ room.max_player }}</td>
              <td style="width: 35%;"><button class="btn btn-outline-info" @click="joinRoom(room.id)">参加</button></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, reactive, ref, useContext, useFetch, useRouter, watch } from '@nuxtjs/composition-api'
import axios from 'axios'
import { key } from '../utils/store'
import { STATE } from '../utils/socket'
import { blobToJson } from '../utils/blobReader'

export default defineComponent({
  name: 'PlayerJoin',
  setup () {
    const { app } = useContext()
    const store = inject(key)
    const name = ref<string>('')
    const nameErr = ref(false)
    const maxPlayer = ref(0)
    const maxPlayerErr = ref(false)
    const router = useRouter()
    if (!store) {
      throw new Error('store undefined')
    }
    const data = reactive(store.data)
    const waiting = ref(false)
    const connectError = ref(false)
    const clientId = generateUuid()

    const rooms = ref()
    const fetch = () => {
      axios.get(`http://${app.$config.apiURL}/rooms`).then((response) => {
        rooms.value = response.data.rooms
      })
    }

    onMounted(() => {
      fetch()
      setInterval(() => {
        fetch()
      }, 5000)
    })

    watch(data, () => {
      console.log('here')
      if (data.game_state === STATE.PlayerFilled) {
        waiting.value = false
        router.push({ name: 'odai' })
      }
    }, { deep: true })

    function handleCloseEvent () {
      console.log('close')
      router.push({ name: 'result' })
    }

    function generateUuid() {
      // https://github.com/GoogleChrome/chrome-platform-analytics/blob/master/src/internal/identifier.js
      // const FORMAT: string = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx";
      const chars = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.split('')
      for (let i = 0, len = chars.length; i < len; i++) {
        switch (chars[i]) {
          case 'x':
            chars[i] = Math.floor(Math.random() * 16).toString(16)
            break
          case 'y':
            chars[i] = (Math.floor(Math.random() * 4) + 8).toString(16)
            break
        }
      }
      return chars.join('')
    }

    const createRoom = () => {
      if (name.value.length === 0) {
        nameErr.value = true
        return
      }
      if (maxPlayer.value < 2) {
        maxPlayerErr.value = true
        return
      }
      nameErr.value = false
      maxPlayerErr.value = false
      store.setName(name.value)
      store.setClientId(clientId)
      waiting.value = true
      const socket = new WebSocket(`ws://${app.$config.apiURL}/?id=${clientId}&name=${name.value}&roomId=${generateUuid()}&newRoom=true&maxPlayer=${maxPlayer.value}`)
      console.log(socket)
      socket.addEventListener('error', () => {
        connectError.value = true
      })
      socket.addEventListener('open', () => {
        console.log('open')
        socket.addEventListener('message', (e) => {
          blobToJson(e.data).then((blobText) => {
            console.log(blobText)
            store.setData(blobText)
          })
        })
        socket.addEventListener('close', handleCloseEvent, false)
        store.setSocket(socket)
      })
    }

    const joinRoom = (id: string) => {
      if (name.value.length === 0) {
        nameErr.value = true
        return
      }
      nameErr.value = false
      store.setName(name.value)
      store.setClientId(clientId)
      waiting.value = true
      const socket = new WebSocket(`ws://${app.$config.apiURL}/?id=${clientId}&roomId=${id}&newRoom=false&name=${name.value}&maxPlayer=0`)
      console.log(socket)
      socket.addEventListener('error', () => {
        connectError.value = true
      })
      socket.addEventListener('open', () => {
        console.log('open')
        socket.addEventListener('message', (e) => {
          blobToJson(e.data).then((blobText) => {
            console.log(blobText)
            store.setData(blobText)
          })
        })
        socket.addEventListener('close', handleCloseEvent, false)
        store.setSocket(socket)
      })
    }

    return {
      name,
      maxPlayer,
      connectError,
      createRoom,
      waiting,
      rooms,
      joinRoom,
      nameErr,
      maxPlayerErr
    }
  }
})
</script>
