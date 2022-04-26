package main

var roomMaxPlayer int = 2

var themas []string = []string{"動物", "お菓子", "食べ物", "地名", "日用品", "飲み物", "乗り物"}

const (
	Initial = iota
	PlayerFilled
	GameStart
	GameEnd
	GameStop
	GameDrawEnd
)

const (
	HiddenWord = iota
	OpenedWord
)

const maxMessageSize = 512
