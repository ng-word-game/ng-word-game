package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 256,
}

const (
	resultOK  = iota
	resultErr = iota
)

type NgChar struct {
	Name string `json:"name"`
	Char string `json:"char"`
}

type outbound struct {
	Result    int    `json:"result"`
	GameState int    `json:"game_state"`
	Thema     string `json:"thema"`
	Users     []struct {
		Id   string
		Name string
	} `json:"users"`
	Words     map[string]string    `json:"words"`
	WordState map[string]WordState `json:"word_state"`
	NextTurn  string               `json:"next_turn"`
	Winner    string               `json:"winner"`
	NgChars   []NgChar             `json:"ng_chars"`
	Turn      int                  `json:"turn"`
}

func createOutbound(result int, room *Room) ([]byte, error) {
	room.mux.RLock()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	defer room.mux.RUnlock()
	clientNames := []struct {
		Id   string
		Name string
	}{}
	words := map[string]string{}
	clientWordStates := map[string]WordState{}

	for _, v := range room.clients {
		clientNames = append(clientNames, struct {
			Id   string
			Name string
		}{Id: v.id, Name: v.name})
		words[v.id] = v.Word
		clientWordStates[v.id] = v.wordState
	}

	nextTurn := ""
	if 0 <= room.nextClientIdx && room.nextClientIdx < len(room.clientIds) {
		if room.clients[room.clientIds[room.nextClientIdx]] != nil {
			nextTurn = room.clients[room.clientIds[room.nextClientIdx]].id
		}
	}
	out, err := json.Marshal(outbound{
		Result:    result,
		GameState: room.gameState,
		Thema:     room.thema,
		Users:     clientNames,
		Words:     words,
		WordState: clientWordStates,
		NextTurn:  nextTurn,
		Winner:    room.winner,
		NgChars:   room.ngChars,
		Turn:      room.gameTurn,
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

const (
	setWord   = iota
	setNgChar = iota
	PING      = iota
)

type inbound struct {
	Type   int    `json:"type"`
	Word   string `json:"word"`
	NgChar string `json:"ng_char"`
}

type wsHandler struct {
	rooms Rooms
}

func NewWshandler() *wsHandler {
	return &wsHandler{
		rooms: map[*Room]bool{},
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			r.joinRoom(client)
			out, err := createOutbound(resultOK, r)
			if err != nil {
				continue
			}
			r.WithLockRoom(func() {
				for _, client := range r.clients {
					client.send <- out
				}
			})
		case client := <-r.leave:
			if r.clientIds[r.nextClientIdx] == client.id {
				r.changeTurn()
			}
			r.mux.RLock()
			delete(r.clients, client.id)
			clientsLen := len(r.clients)
			r.mux.RUnlock()
			if clientsLen == 1 {
				r.gameStop()
			}
			out, err := createOutbound(resultOK, r)
			if err != nil {
				continue
			}
			for _, client := range r.clients {
				client.send <- out
			}
			if clientsLen == 0 {
				r.available = false
				log.Println("close room, goroutine: ", runtime.NumGoroutine())
				return
			}
		case send := <-r.send:
			var in inbound
			err := json.Unmarshal(send.body, &in)
			if err != nil {
				send.from.handleError(err)
			}
			switch in.Type {
			case setWord:
				log.Printf("word: " + in.Word)
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
				r.changeTurn()
				// if r.clients.values()[r.nextClientIdx] != send.from {
				// 	continue
				// }
				log.Printf("ngChar: " + in.NgChar)
				r.addNgChar(send.from, in.NgChar)
				r.applyNgChar(in.NgChar)
				if checkEnd, winner := r.checkEndGame(); checkEnd {
					r.gameEnd(winner)
				}
				out, err := createOutbound(resultOK, r)
				if err != nil {
					continue
				}
				for _, client := range r.clients {
					client.send <- out
				}
			case PING:
				if r.gameState != GameStop {
					continue
				}
				out, err := json.Marshal(outbound{GameState: r.gameState})
				if err != nil {
					continue
				}
				send.from.send <- out
			default:
				log.Printf(send.from.name + "'s message is unknown")
			}
		}
	}
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

func (h *wsHandler) ServeWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}

	client := h.NewClient(r.FormValue("id"), r.FormValue("name"), conn)
	n, err := strconv.Atoi(r.FormValue("maxPlayer"))
	if err != nil {
		return
	}
	newRoom, err := strconv.ParseBool(r.FormValue("newRoom"))
	if err != nil {
		return
	}
	roomId := r.FormValue("roomId")
	if !newRoom {
		canEnter, room := h.rooms.filterById(roomId)
		if !canEnter {
			return
		}
		room.join <- client
	} else {
		log.Printf("create room")
		room := NewRoom(roomId, n)
		h.rooms[room] = true
		go room.run()
		room.join <- client
	}

	defer func() {
		log.Printf("close ServeHTTP()")
		log.Println("run handler goroutine: ", runtime.NumGoroutine())
		client.room.leave <- client
	}()
	go client.write()
	client.read()
}

func (h *wsHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodGet {
		return
	}
	type RoomType struct {
		Id        string   `json:"id"`
		Num       int      `json:"num"`
		MaxPlayer int      `json:"max_player"`
		Players   []string `json:"players"`
	}
	type outType struct {
		Rooms []RoomType `json:"rooms"`
	}
	var rooms []RoomType
	for _, room := range h.rooms.filterAvailable() {
		room.WithLockRoom(func() {
			var players []string
			for _, client := range room.clients {
				players = append(players, client.name)
			}
			rooms = append(rooms, RoomType{
				Id:        room.id,
				Num:       len(room.clients),
				MaxPlayer: room.maxPlayer,
				Players:   players,
			})
		})
	}
	out, err := json.Marshal(&outType{Rooms: rooms})
	if err != nil {
		return
	}
	w.WriteHeader(200)
	w.Write(out)
}
