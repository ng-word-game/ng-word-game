export type InterfaceSocketData = {
    game_state: number
    next_turn: string
    turn: number
    result: number
    thema: string
    users: [{Id: string, Name: string}] | null
    winner: string
    word_state: { [key: string]: [{[key: string]: number}]}
    words: { [key: string]: string },
    ng_chars: [{
      name: '',
      char: ''
    }]
}

export enum STATE {
    Initial,
    PlayerFilled,
    GameStart,
    GameEnd,
    GameStop,
}

export enum SET {
    SetWord,
    SetNgChar,
    PING
}

export enum WORDSTATE {
    HiddenWord,
    OpenedWord,
}
