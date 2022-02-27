<template>
  <OdaiCom />
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import OdaiCom from '../components/Odai.vue'

export default defineComponent({
  name: 'OdaiPage',
  components: {
    OdaiCom
  },
  setup () {
    return {}
  }
})
</script>
