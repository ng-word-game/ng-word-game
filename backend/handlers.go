package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"runtime"

	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 256,
}

const (
	resultOK = iota
	resultErr = iota
)

type NgChar struct {
	Name string `json:"name"`
	Char string `json:"char"`
}

type outbound struct {
	Result int `json:"result"`
	GameState int `json:"game_state"`
	Thema string `json:"thema"`
	Users []string `json:"users"`
	Words map[string]string `json:"wors"`
	WordState map[string]WordState `json:"word_state"`
	NextTurn string `json:"next_turn"`
	Winner string `json:"winner"`
	NgChars []NgChar `json:"ng_chars"`
}

func createOutbound(result int, room *Room) ([]byte, error) {
	clientNames := []string{}
	words := map[string]string{}
	clientWordStates := map[string]WordState{}

	for _, v := range room.clients.values() {
		clientNames = append(clientNames, v.name)
		words[v.name] = v.Word
		clientWordStates[v.name] = v.wordState
	}

	out, err := json.Marshal(outbound{
		Result:    result,
		GameState: room.gameState,
		Thema:     room.thema,
		Users:     clientNames,
		Words:     words,
		WordState: clientWordStates,
		NextTurn:  room.nextClientName,
		Winner:    room.winner,
		NgChars:   room.ngChars,
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

const (
	setWord = iota
	setNgChar = iota
)

type inbound struct {
	Type int `json:"type"`
	Word string `json:"word"`
	NgChar string `json:"ng_char"`
}

type wsHandler struct {
	rooms Rooms
	join chan *Client
	leave chan *Client
	close chan struct{}
	roomSend chan struct{
		from *Client
		room *Room
		body []byte
	}
}

func NewWshandler() *wsHandler {
	return &wsHandler{
		rooms: map[*Room]bool{},
		join:  make(chan *Client),
		leave: make(chan *Client),
		close: make(chan struct{}),
		roomSend: make(chan struct {
			from *Client
			room *Room
			body []byte
		}),
	}
}

func (h *wsHandler) run() {
	log.Println("run handler goroutine: ", runtime.NumGoroutine())
	defer func(){
		log.Printf("close h.run()")
		for r := range h.rooms {
			for _, c := range r.clients.values() {
				close(c.close)
				c.conn.Close()
			}
			close(r.close)
		}
	}()

	for {
		select {
		case client := <-h.join:
			log.Printf("%v join", client.name)
			canEnter, room := h.rooms.SearchAvailableRoomIdx()
			if !canEnter  {
				room = h.createRoom()
			}
			room.join(client)
			out, err := createOutbound(resultOK, room)
			if err != nil {
				continue
			}
			for _, client := range room.clients {
				client.send <- out
			}
		case client := <- h.leave:
			room := client.room
			room.gameStop()
			delete(room.clients, client.id)
			out, err := json.Marshal(outbound{GameState: GameStop})
			if err != nil {
				continue
			}
			for _, client := range room.clients {
				client.send <- out
			}
			delete(h.rooms, client.room)
		case <- h.close:
			return
		}
	}
}

func (r *Room) run() {
	for {
		select {
		case <- r.close:
			log.Printf("close room")
			return
		case send := <- r.send:
			var in inbound
			err := json.Unmarshal(send.body, &in)
			if err != nil {
				send.from.handleError(err)
			}
			switch in.Type {
			case setWord:
				log.Printf("word: "+ in.Word)
				send.from.setWord(in.Word)
				out, err := createOutbound(resultOK, r)
				if err != nil {
					continue
				}
				send.from.send <- out
	
				if r.checkMaxPlayer() && r.checkAllWords() {
					r.gameStart()
					out, err := createOutbound(resultOK, r)
					if err != nil {
						continue
					}
					for _, client := range r.clients {
						client.send <- out
					}
				}
			case setNgChar:
				if r.gameState != GameStart {
					continue
				}
				// if r.clients.values()[r.nextClientIdx] != send.from {
				// 	continue
				// }
				log.Printf("ngChar: " + in.NgChar)
				r.ngChars = append(r.ngChars, NgChar{Name: send.from.name, Char: in.NgChar})
				r.changeTurn()
				r.applyNgChar(in.NgChar)
				if checkEnd, _ := r.checkEndGame(); checkEnd {
					r.gameEnd(send.from.name)
				}
				out, err := createOutbound(resultOK, r)
				if err != nil {
					continue
				}
				for _, client := range r.clients {
					client.send <- out
				}
			default:
				log.Printf(send.from.name + "'s message is unknown")
			}
		}
	}
}

func (h *wsHandler) createRoom() *Room {
	room := NewRoom()
	h.rooms[room] = true
	go room.run()

	return room
}

func (h *wsHandler) NewClient(id string, name string, c *websocket.Conn) *Client {
	log.Printf("new client: " + name)
	return &Client{
		wsHandler: h,
		room:      &Room{},
		id:        id,
		name:      name,
		Word:      "",
		conn:      c,
		send:      make(chan []byte),
		close:     make(chan struct{}),
		wordState: []map[string]int{},
	}
}

func (h *wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}

	client := h.NewClient(randomString(10), r.FormValue("name"), conn)
	h.join <- client
	defer func() {
		log.Printf("close ServeHTTP(), この時点でh.closeされている")
	}()
	go client.write()
	client.read()
}

func randomString(l int) string {
	bytes := make([]byte, l)
	pool := "_:$%&/()"
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}

	return string(bytes)
}
