<template>
  <div
    class="kana-palette-key"
    :class="className"
    @click="onClick"
  >
    {{ props.char }}
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "@nuxtjs/composition-api";

export default defineComponent({
  name: "KanaPaletteKeyCom",
  props: {
    char: {
      type: String as () => string,
      default: '',
    },
    isOpened: {
      type: Boolean as () => boolean,
      default: false,
    },
    isDangerous: {
      type: Boolean as () => boolean,
      default: false,
    },
    isSelected: {
      type: Boolean as () => boolean,
      default: false,
    },
    isEnabled: {
      type: Boolean as () => boolean,
      default: true,
    },
  },
  setup(props, context) {
    const className = computed(() => {
      if (!props.isEnabled) return ['is-disabled'];
      if (props.isOpened) return ['is-opened'];
      if (props.isDangerous) return ['is-dangerous'];
      if (props.isSelected) return ['is-selected'];
      return [];
    });
    
    const onClick = () => {
      if (!props.isEnabled || props.isOpened || props.isSelected) {
        return;
      }

      context.emit('click', props.char);
    };
    return {
      props,
      className,
      onClick,
    };
  }
});
</script>

<style scoped>
.kana-palette-key {
  border: solid 1px rgba(0, 0, 0, 0.125);
  border-radius: 0.25rem;
  text-align: center;
  cursor: pointer;
}

.kana-palette-key.is-selected {
  border-color: rgb(23, 162, 184);
  background-color: rgba(23, 162, 184, 0.125);
  cursor: default;
}

.kana-palette-key.is-dangerous {
  border-color: rgb(220, 53, 69);
  background-color: rgba(220, 53, 69, 0.125);
}

.kana-palette-key.is-opened {
  background-color: rgba(0, 0, 0, 0.125);
  cursor: default;
}

.kana-palette-key.is-disabled {
  border-color: rgba(0, 0, 0, 0);
  cursor:default;
}
</style>
