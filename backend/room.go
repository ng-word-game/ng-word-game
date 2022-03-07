package main

import (
	"errors"
	"log"
	"math/rand"
	"reflect"
	"sync"
)

type Room struct {
	mux sync.RWMutex
	available bool
	clients Clients
	gameState int
	thema string
	nextClientIdx int
	clientIds []string
	ngChars []NgChar
	winner string
	send chan struct{
		from *Client
		body []byte
	}
	leave chan *Client
}

func (r *Room) WithLockRoom(fn func()) {
	r.mux.Lock()
	fn()
	defer r.mux.Unlock()
}

type Rooms map[*Room]bool

func (r Rooms) keys() []*Room {
    ks := []*Room{}
    for k, _ := range r {
        ks = append(ks, k)
    }
    return ks
}

func NewRoom() *Room {
	return &Room{
		mux:           sync.RWMutex{},
		available:     true,
		clients:       map[string]*Client{},
		gameState:     Initial,
		thema:         themas[rand.Intn(len(themas))],
		nextClientIdx: 0,
		clientIds:     []string{},
		ngChars:       []NgChar{},
		winner:        "",
		send: make(chan struct {
			from *Client
			body []byte
		}),
		leave: make(chan *Client),
	}
}

func (r *Room) join(client *Client) {
	log.Printf("%v join room", client.name)
	client.room = r
	r.clients[client.id] = client
	r.clientIds = append(r.clientIds, client.id)
	if len(r.clients) >= roomMaxPlayer {
		r.WithLockRoom(func() {
			r.available = false
			r.gameState = PlayerFilled
		})
	}
}

func (r *Room) gameStart() {
	r.WithLockRoom(func() {
		log.Printf("Game Start")
		r.gameState = GameStart
		r.nextClientIdx = 0
	})
}

func (r *Room) gameEnd(c *Client) {
	r.WithLockRoom(func() {
		log.Printf("Game End")
		r.gameState = GameEnd
		r.winner = c.id
	})
}

func (r *Room) checkEndGame() (bool, *Client) {
	passedPlayer := 0
	var winner *Client
	for _, client := range r.clients {
		pass := true
		for _, v := range client.wordState {
			for _, i := range v{
				if i != OpenedWord {
					pass = false
				}
			}
		}
		if pass {
			passedPlayer++
		} else {
			winner = client
		}
	}
	if passedPlayer == len(r.clients) - 1 {
		return true, winner
	}
	return false, nil
}

func (r *Room) changeTurn() {
	r.WithLockRoom(func() {
		if r.nextClientIdx + 1 < roomMaxPlayer {
			r.nextClientIdx++
		} else {
			r.nextClientIdx = 0
		}
	})
}

func (r *Room) addNgChar(from *Client, ngChar string) {
	r.WithLockRoom(func() {
		r.ngChars = append(r.ngChars, NgChar{Name: from.name, Char: ngChar})
	})
}

func (r *Room) applyNgChar(ngChar string) {
	for _, item := range(r.ngChars) {
		for _, client := range r.clients.values() {
			client.wordState.applyNgchar(item.Char)
		}
	}
	for _, client := range r.clients.values() {
		log.Println("name: ", client.name, "word: ", client.wordState)
	}
}

func(r *Room) checkAllWords() bool {
	all := true
	for _, v := range r.clients {
		if v.Word == "" {
			all = false
		}
	}

	return all
}

func (r *Room) checkMaxPlayer() bool {
	if len(r.clients) == roomMaxPlayer {
		return true
	}

	return false
}

func (r *Room) gameStop() {
	r.WithLockRoom(func() {
		r.available = false
		r.gameState = GameStop
	})
}

func (r Rooms) SearchAvailableRoomIdx() (bool, *Room) {
	for _, v := range r.keys() {
		if v.available {
			return true, v
		}
	}

	return false, nil
}

func (r Rooms) SearchRoomByClient(client *Client) (error, *Room) {
	for _, room := range r.keys() {
		for _, v := range room.clients {
			if reflect.DeepEqual(client, &v) {
				return nil, room
			}
		}
	}

	return errors.New("can't find room"), nil
}
