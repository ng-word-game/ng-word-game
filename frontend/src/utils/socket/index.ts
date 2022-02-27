export type InterfaceSocketData = {
    game_state: number
    next_turn: string
    result: number
    thema: string
    users: string[]
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
}

export enum WORDSTATE {
    HiddenWord,
    OpenedWord,
}
