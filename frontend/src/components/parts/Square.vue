/* eslint-disable object-property-newline */
<template>
  <div>
    <div class="d-flex">
      <div v-for="(v, idx) in props.charaInfo" :key="idx">
        <div v-if="v.isHide">
          <div v-if="isSmall(v.char)" class="circle--small hide">・</div>
          <div v-else-if="isDakuon(v.char)" class="mx-2" style="position: relative;">
            <div class="circle hide">・</div>
            <div class="dakuten hide" style="position: absolute; top: 0; right: -13px;"></div>
          </div>
          <div v-else-if="isHanDakuon(v.char)" class="mx-2" style="position: relative;">
            <div class="circle hide">・</div>
            <div class="dakuten hide" style="position: absolute; top: 0; right: -13px;"></div>
          </div>
          <div v-else class="circle hide mx-2">・</div>
        </div>
        <div v-else-if="isSmall(v.char)" class="circle--small">{{ changeChar[v.char] }}</div>
        <div v-else-if="isDakuon(v.char)" class="mx-2" style="position: relative;">
          <div class="circle">{{ changeChar[v.char] }}</div>
          <div class="dakuten" style="position: absolute; top: 0; right: -13px;">〃</div>
        </div>
        <div v-else-if="isHanDakuon(v.char)" class="mx-2" style="position: relative;">
          <div class="circle">{{ changeChar[v.char] }}</div>
          <div class="dakuten" style="position: absolute; top: 0; right: -13px; font-size: 10px;">○</div>
        </div>
        <div v-else class="circle mx-2">
          {{ v.char }}
        </div>
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
    const changeChar: {[key: string]: string} = {
      'ゃ': 'や', 'ゅ': 'ゆ', 'ょ': 'よ', 'ぁ': 'あ', 'ぃ': 'い', 'ぅ': 'う', 'ぇ': 'え', 'ぉ': 'お', 'っ': 'つ', 'ゎ': 'わ',
      'が': 'か', 'ぎ': 'き', 'ぐ': 'く', 'げ': 'け', 'ご': 'こ', 'ざ': 'さ', 'じ': 'し', 'ず': 'す', 'ぜ': 'せ',
      'ぞ': 'そ', 'だ': 'た', 'ぢ': 'ち', 'づ': 'つ', 'で': 'て', 'ど': 'と', 'ば': 'は', 'び': 'ひ', 'ぶ': 'ふ',
      'べ': 'へ', 'ぼ': 'ほ', 'ぱ': 'は', 'ぴ': 'ひ', 'ぷ': 'ふ', 'ぺ': 'へ', 'ぽ': 'ほ'
    }
    const isSmall = (c: string) => {
      return ['ゃ', 'ゅ', 'ょ', 'ぁ', 'ぃ', 'ぅ', 'ぇ', 'ぉ', 'っ', 'ゎ'].includes(c)
    }
    const isDakuon = (c: string) => {
      return ['が', 'ぎ', 'ぐ', 'げ', 'ご', 'ざ', 'じ', 'ず', 'ぜ', 'ぞ', 'だ', 'ぢ', 'づ', 'で', 'ど', 'ば', 'び', 'ぶ', 'べ', 'ぼ'].includes(c)
    }
    const isHanDakuon = (c: string) => {
      return ['ぱ', 'ぴ', 'ぷ', 'ぺ', 'ぽ'].includes(c)
    }
    return {
      props,
      isSmall,
      isDakuon,
      isHanDakuon,
      changeChar
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
.circle{
  display: inline-block;
  border: 1.2px solid;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  text-align:center;
  line-height: 40px;
  font-size: 30px;
}
.circle--small {
  display: inline-block;
  border: 1.2px solid;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  text-align:center;
  line-height: 16px;
  margin-top: 20px;
}
.circle--dakuten {
  display: inline-block;
  border: 1.2px solid;
  width: 35px;
  height: 35px;
  border-radius: 50%;
  text-align:center;
  line-height: 35px;
  margin-top: 5px;
}
.dakuten {
  display: inline-block;
  border: 1.2px solid;
  width: 15px;
  height: 15px;
  border-radius: 50%;
  text-align:center;
  line-height: 12px;
}
.opened {
  background-color: gray;
}
.hide {
  background-color: black;
}
</style>
