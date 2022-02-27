<template>
  <PlayCom />
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import PlayCom from '../components/Play.vue'

export default defineComponent({
  name: 'PlayPage',
  components: {
    PlayCom
  },
  setup () {
    return {}
  }
})
</script>
