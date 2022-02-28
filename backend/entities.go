package main

type WordState []map[string]int

func (ws WordState) applyNgchar(char string) {
	for _, item := range ws {
		if _, ok := item[char]; ok {
			item[char] = OpenedWord
		}
	}
}
