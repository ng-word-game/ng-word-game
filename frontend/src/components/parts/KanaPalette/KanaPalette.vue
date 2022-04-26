<template>
  <div class="kana-palette">
    <div v-for="(section, sIndex) in keys" :key="sIndex" class="section">
      <div v-for="(row, rIndex) in section" :key="rIndex" class="row">
        <div v-for="(key, index) in row" :key="index" class="key">
          <KanaPaletteKeyCom
            v-if="key != null"
            :key="sIndex+rIndex+index"
            :char="key.char"
            :is-opened="key.isOpened"
            :is-dangerous="key.isDangerous"
            :is-selected="key.isSelected"
            @click="onKeyClick"
          />
          <KanaPaletteKeyCom
            v-else
            :key="sIndex+rIndex+index"
            :is-enabled="false"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from '@nuxtjs/composition-api'
import KanaPaletteKeyCom from './KanaPaletteKey.vue'

interface KanaPalette {
  sections: KanaPaletteSection[];
}

// 清音セクションと濁音・半濁音セクションを縦に並べる用途
interface KanaPaletteSection {
  rows: KanaPaletteRow[];
}

// 五十音表に合わせて、縦の行を右から並べる
interface KanaPaletteRow {
  keys: (KanaPaletteKey | null)[];
}

interface KanaPaletteKey {
  char: string;
  isOpened: boolean; // 解放済み？
  isDangerous: boolean; // 自分の単語に含まれている？
  isSelected: boolean; // 選択中？
}

// 外から Sections -> Rows -> Keys
type KanaPaletteDefinition = string[][][];

type KanaPaletteBuilder = (openedChars: string[], dangerousChars: string[], selectedChar: string) => KanaPalette;

const defineKanaPalette = (chars: KanaPaletteDefinition) => {
  return ((openedChars: string[], dangerousChars: string[], selectedChar: string) => {
    const toKey: (c: string) => KanaPaletteKey | null = (c: string) => {
      if (c === '') {
        return null
      }

      return ({
        char: c,
        isOpened: openedChars.includes(c),
        isDangerous: dangerousChars.includes(c),
        isSelected: selectedChar === c
      } as KanaPaletteKey)
    }
    const toRow = (chars: string[]) => ({ keys: chars.map(toKey) })
    const toSection = (rows: string[][]) => ({ rows: rows.map(toRow) })
    const palette = { sections: chars.map(toSection) } as KanaPalette

    return palette
  }) as KanaPaletteBuilder
}

const transformToVertical = (section: KanaPaletteSection) => {
  const width = section.rows.length
  const height = Math.max(...section.rows.map(r => r.keys.length))
  const keys =
    [...Array(height)].map((_, v) =>
      [...Array(width)].map((_, h) =>
        section.rows[width - h - 1].keys[v]
      ))
  return keys
}

const kanaPaletteBuilder = defineKanaPalette([
  [
    ['あ', 'い', 'う', 'え', 'お'],
    ['か', 'き', 'く', 'け', 'こ'],
    ['さ', 'し', 'す', 'せ', 'そ'],
    ['た', 'ち', 'つ', 'て', 'と'],
    ['な', 'に', 'ぬ', 'ね', 'の'],
    ['は', 'ひ', 'ふ', 'へ', 'ほ'],
    ['ま', 'み', 'む', 'め', 'も'],
    ['や', '', 'ゆ', '', 'よ'],
    ['ら', 'り', 'る', 'れ', 'ろ'],
    ['わ', 'を', 'ん', '', 'ー']
  ],
  [
    ['が', 'ぎ', 'ぐ', 'げ', 'ご'],
    ['ざ', 'じ', 'ず', 'ぜ', 'ぞ'],
    ['だ', 'ぢ', 'づ', 'で', 'ど'],
    ['ば', 'び', 'ぶ', 'べ', 'ぼ'],
    ['ぱ', 'ぴ', 'ぷ', 'ぺ', 'ぽ'],
    ['ぁ', 'ぃ', 'ぅ', 'ぇ', 'ぉ'],
    ['ゃ', 'ゅ', 'ょ', '', 'っ'],
    ['ゔ', 'ゐ', 'ゑ']
  ]
])

export default defineComponent({
  name: 'KanaPaletteCom',
  components: {
    KanaPaletteKeyCom
  },
  props: {
    openedChars: {
      type: Array as () => string[],
      default: () => []
    },
    dangerousChars: {
      type: Array as () => string[],
      default: () => []
    },
    selectedChar: {
      type: String as () => string,
      default: ''
    }
  },
  setup (props, context) {
    const keys = computed(() => {
      const palette = kanaPaletteBuilder(props.openedChars, props.dangerousChars, props.selectedChar)
      const transformedSections = palette.sections.map(s => transformToVertical(s))
      return transformedSections
    })

    const onKeyClick = (char: string) => {
      context.emit('selected', char)
    }

    return {
      props,
      keys,
      onKeyClick
    }
  }
})
</script>

<style scoped>
.kana-palette {
  padding: 10px;
}

.kana-palette .section {
  margin-bottom: 0.5em;
}

.kana-palette .row {
  display: flex;
  justify-content: center;
}
.kana-palette .key {
  width: 2em;
  height: 2em;
  line-height: 2em;
}
</style>
