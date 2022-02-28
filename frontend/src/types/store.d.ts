import { InterfaceSocketData } from '~/utils/socket'

export interface Store {
  state: {
    readonly name: string
    readonly myword: string
    readonly socket: WebSocket|null
  },
  data: {
    readonly game_state: number
    readonly next_turn: string
    readonly result: number
    readonly thema: string
    readonly users: [{ Id: string, Name: string }]
    readonly winner: string
    readonly word_state: { [key: string]: [{[key: string]: number}]}
    readonly words: { [key: string]: string }
    readonly ng_chars: [{ name: string, char: string }]
  }
  setName: (n: string) => void
  setMyWord: (n: string) => void
  setSocket: (s: WebSocket|null) => void
  setData: (d: InterfaceSocketData) => void
}
