declare type InterfaceSocketData = {
    game_state: number
    next_turn: string
    turn: number
    result: number
    thema: string
    users: [{ Id: string, Name: string }] | null
    winner: string
    word_state: { [key: string]: {[key: string]: number}}
    words: { [key: string]: string }
    ng_chars: [{ name: string, char: string }]
}
