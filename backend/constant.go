package main

var roomMaxPlayer int = 2

var themas []string = []string{"動物", "新宿", "渋谷", "お菓子", "バイト"}

const (
	Initial      = iota
	PlayerFilled = iota
	GameStart    = iota
	GameEnd      = iota
	GameStop     = iota
)

const (
	HiddenWord = iota
	OpenedWord = iota
)

const maxMessageSize = 512
