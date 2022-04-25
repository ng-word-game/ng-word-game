package main

var roomMaxPlayer int = 2

var themas []string = []string{"動物", "お菓子", "食べ物", "地名", "日用品", "飲み物", "乗り物"}

const (
	Initial      = iota
	PlayerFilled = iota
	GameStart    = iota
	GameEnd      = iota
	GameStop     = iota
	GameDrawEnd  = iota
)

const (
	HiddenWord = iota
	OpenedWord = iota
)

const maxMessageSize = 512
