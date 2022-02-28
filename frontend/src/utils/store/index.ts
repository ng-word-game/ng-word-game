import { InjectionKey, reactive, readonly } from '@vue/composition-api'
import { Store } from '../../types/store'
import { InterfaceSocketData, STATE } from '../socket'

interface InterfaceState {
  name: string
  myword: string
  socket: WebSocket|null
}

const initialData = reactive<InterfaceSocketData>({
  game_state: STATE.Initial,
  next_turn: '',
  result: 0,
  thema: '',
  users: null,
  winner: '',
  word_state: {},
  words: { '': '' },
  ng_chars: [{ name: '', char: '' }]
})

const state = reactive<InterfaceState>({
  name: '',
  myword: '',
  socket: null
})

const setName = (n: string): void => {
  state.name = n
}

const setMyWord = (n: string): void => {
  state.myword = n
}

const setSocket = (s: WebSocket): void => {
  state.socket = s
}

const setData = (d: InterfaceSocketData): void => {
  Object.assign(initialData, d)
}

export default {
  state: readonly(state),
  data: readonly(initialData),
  setName,
  setMyWord,
  setSocket,
  setData
}

export const key: InjectionKey<Store> = Symbol('key')
