<template>
  <div>
    <div class="d-flex">
      <div v-for="(v, idx) in props.charaInfo" :key="idx" class="rect mx-2" :class="[v.isOpen ? 'opened': '', v.isHide ? 'hide': '']">
        <p v-if="v.isHide">
          ・
        </p>
        <p v-else>
          {{ v.char }}
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import CharaInfoInterface from '@/types/CharaInfoInterface'

export default defineComponent({
  name: 'SquareCom',
  props: {
    charaInfo: {
      type: Array as () => CharaInfoInterface[],
      default: null
    }
  },
  setup (props) {
    return {
      props
    }
  }
})
</script>

<style scoped>
.rect {
  border: 1.2px solid;
  display: table-cell;
  text-align: center;
  vertical-align: middle;
}
.rect p {
  margin: 10px;
  font-size: 1.5rem;
  font-weight: bold;
}
.opened {
  background-color: gray;
}
.hide {
  background-color: black;
}
</style>
