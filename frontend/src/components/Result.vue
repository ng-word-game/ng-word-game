<template>
  <div class="d-flex justify-content-center align-items-center" style="height: 100vh;">
    <div class="card" style="width: 50%;">
      <div class="card-body">
        <h3 v-if="store.data == undefined" class="card-title text-center">
          接続が切断されました
        </h3>
        <h3 v-else-if="store.data.game_state == STATE.GameEnd" class="card-title text-center">
          Result
        </h3>
        <h3 v-else-if="store.data.game_state == STATE.GameStop" class="card-title text-center">
          対戦相手が退出しました
        </h3>
        <h3 v-else-if="store.data.game_state == STATE.GameDrawEnd" class="card-title text-center">
          引き分け
        </h3>
        <h4 v-if="store.data.game_state != undefined && store.data.game_state == STATE.GameEnd" class="text-center">
          {{ winner() }}の勝利
        </h4>
        <h5 class="mt-3 text-center">~結果~</h5>
        <div v-if="store.data.game_state != undefined && store.data.game_state == STATE.GameEnd" class="text-center">
          <div v-for="user in store.data.users" :key="user.Id">
            <div>{{ user.Name }}の単語: 「{{ store.data.words[user.Id] }}」</div>
          </div>
        </div>
        <div class="d-flex flex-column justify-content-center">
          <div class="mx-auto mt-3">
            <button type="submit" class="btn btn-outline-info" @click="goIndex">
              最初に戻る
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, useRouter } from '@nuxtjs/composition-api'
import { key } from '../utils/store'
import { STATE } from '../utils/socket'

export default defineComponent({
  name: 'ResultCom',
  setup () {
    const store = inject(key)
    if (!store) {
      throw new Error('store undefined')
    }
    const router = useRouter()

    return {
      store,
      STATE,
      goIndex: () => router.push({ name: 'index' }),
      winner: () => {
        if (store.data) {
          return store.data.users.filter(u => u.Id === store.data.winner)[0].Name
        }
      }
    }
  }
})
</script>
