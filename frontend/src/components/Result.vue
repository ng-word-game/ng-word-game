<template>
  <div class="d-flex justify-content-center align-items-center" style="height: 100vh;">
    <div class="card" style="width: 50%;">
      <div class="card-body">
        <h3 v-if="resultMode()" class="card-title text-center">
          Result
        </h3>
        <h3 v-else-if="stopMode()" class="card-title text-center">
          対戦相手が退出しました
        </h3>
        <h4 v-if="resultMode()" class="text-center">
          {{ store.data.winner }}の勝利
        </h4>
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
      resultMode: () => store.data.game_state === STATE.GameEnd,
      stopMode: () => store.data.game_state === STATE.GameStop,
      goIndex: () => router.push({ name: 'index' })
    }
  }
})
</script>
