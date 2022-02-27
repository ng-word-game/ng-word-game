package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 3 * time.Minute

	// Time allowed to read the next pong message from the peer.
	pongWait = 3 * time.Minute

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var roomMaxPlayer int = 2

var themas []string = []string{"動物", "新宿", "渋谷", "お菓子", "バイト"}

const (
	Initial = iota
	PlayerFilled = iota
	GameStart = iota
	GameEnd = iota
	GameStop = iota
)

const (
	HiddenWord = iota
	OpenedWord = iota
)

type WordState []map[string]int

func (ws WordState) applyNgchar(char string) {
	for _, item := range ws {
		if _, ok := item[char]; ok {
			item[char] = OpenedWord
		}
	}
}

type Room struct {
	available bool
	clients Clients
	send chan []byte
	gameState int
	thema string
	nextClientIdx int
	ngChars []NgChar
	winner string
}

type Rooms map[*Room]bool
type Clients map[string]*Client

func (r Rooms) keys() []*Room {
    ks := []*Room{}
    for k, _ := range r {
        ks = append(ks, k)
    }
    return ks
}

type Client struct {
	wsHandler *wsHandler
	room *Room
	id string
	name string
	Word string
	conn *websocket.Conn
	send chan []byte
	wordState WordState
}

func (c Clients) values() []*Client {
    vs := []*Client{}
    for _, v := range c {
        vs = append(vs, v)
    }
    return vs
}

func NewRoom() *Room {
	return &Room{
		available:     true,
		clients:       map[string]*Client{},
		send:          make(chan []byte),
		gameState:     Initial,
		thema:         themas[rand.Intn(len(themas))],
		nextClientIdx: 0,
		ngChars: []NgChar{},
		winner: "",
	}
}

func (r *Room) join(client *Client) {
	log.Printf("%v join room", client.name)
	client.room = r
	r.clients[client.id] = client
	if len(r.clients) >= roomMaxPlayer {
		r.available = false
		r.gameState = PlayerFilled
	}
}

func (r *Room) gameStart() {
	log.Printf("Game Start")
	r.gameState = GameStart
	r.nextClientIdx = 0
}

func (r *Room) gameEnd(name string) {
	log.Printf("Game End")
	r.gameState = GameEnd
	r.winner = name
}

func (r *Room) checkEndGame() (bool, string) {
	allPass := true
	for _, client := range r.clients {
		pass := true
		for _, v := range client.wordState {
			for _, i := range v{
				if i != OpenedWord {
					pass = false
				}
			}
		}
		if allPass && pass {
			return true, client.name
		}
	}

	return false, ""
}

func (r *Room) changeTurn() {
	if r.nextClientIdx + 1 < roomMaxPlayer {
		r.nextClientIdx += 1
	} else {
		r.nextClientIdx = 0
	}
	log.Printf("change turn to " + r.clients.values()[r.nextClientIdx].name)
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
	r.available = false
	r.gameState = GameStop
}

func (c *Client) write() {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case msg, ok := <- c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				log.Println(err)
				return
			}
		case <- ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Client) read() {
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	defer func() {
		log.Println(c.name, " leave")
		c.wsHandler.leave <- c
		c.conn.Close()
	}()
	for {
		if _, msg, err := c.conn.ReadMessage(); err == nil {
			c.wsHandler.roomSend <- struct{from *Client; room *Room; body []byte}{from: c, room: c.room, body: msg}
		} else {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (c *Client) setWord(word string) {
	c.Word = word
	for _, s := range strings.Split(word, "") {
		c.wordState = append(c.wordState, map[string]int{s: HiddenWord})
	}
}

func (c *Client) handleError(err error) {
	log.Println("Error:", err)

	b, err := json.Marshal(struct{ Body string }{
		Body: err.Error(),
	})
	if err != nil {
		log.Println("Error:", err)
	}

	err = c.conn.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		log.Println("Error:", err)
	}
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
