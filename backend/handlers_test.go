package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/goleak"
)

func httpToWS(t *testing.T, u string) string {
	t.Helper()

	wsURL, err := url.Parse(u)
	if err != nil {
		t.Fatal(err)
	}

	switch wsURL.Scheme {
	case "http":
		wsURL.Scheme = "ws"
	case "https":
		wsURL.Scheme = "wss"
	}

	return wsURL.String()
}

func newWSServer(t *testing.T, h http.Handler, userName string) (*httptest.Server, *websocket.Conn) {
	t.Helper()

	s := httptest.NewServer(h)
	wsURL := httpToWS(t, s.URL) + fmt.Sprintf("?id=%s&name=%v", randomString(10), userName)

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	return s, ws
}

func randomString(l int) string {
	bytes := make([]byte, l)
	pool := "abcdefghijklm"
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}

	return string(bytes)
}

// func TestRegisterUser(t *testing.T) {
// 	tcs := []struct {
// 		name string
// 		userName string
// 	}{
// 		{
// 			name: "with user1",
// 			userName: "test1",
// 		},
// 		{
// 			name: "with user2",
// 			userName: "test2",
// 		},
// 		{
// 			name: "with user3",
// 			userName: "test3",
// 		},
// 	}

// 	h := NewWshandler()
// 	go h.run()
// 	pass := true
// 	for _, tt := range tcs {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s, ws := newWSServer(t, h, tt.userName)
// 			time.Sleep(time.Second * 1)
// 			defer s.Close()
// 			defer ws.Close()
// 		})
// 	}

// 	for _, tt := range tcs {
// 		exist := false
// 		for _, room := range h.rooms.keys() {
// 			for _, client := range room.clients.values() {
// 				if client.name == tt.userName {
// 					exist = true
// 				}
// 			}
// 		}
// 		pass = pass && exist
// 	}
// 	if !pass {
// 		t.Fatalf("cannot create user")
// 	}
// }

func TestGoRoutine(t *testing.T) {
	defer goleak.VerifyNone(t)
	tcs := []struct {
		name string
		users []struct {
			userName string
		}
		roomCount int
	}{
		{
			name: "1 user",
			users: []struct {
				userName string
			}{{userName: "usr1"},},
			roomCount: 1,
		},
		{
			name: "2 users",
			users: []struct {
				userName string
			}{{userName: "usr1"}, {userName: "usr2"}},
			roomCount: 1,
		},
		{
			name: "3 users",
			users: []struct {
				userName string
			}{{userName: "usr1"}, {userName: "usr2"}, {userName: "usr3"}},
			roomCount: 2,
		},
		{
			name: "10 users",
			users: []struct {
				userName string
			}{{userName: "usr1"}, {userName: "usr2"},
			{userName: "usr3"}, {userName: "usr4"},
			{userName: "usr5"}, {userName: "usr6"},
			{userName: "usr7"}, {userName: "usr8"},
			{userName: "usr9"}, {userName: "usr10"}},
			roomCount: 5,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			h := NewWshandler()
			// servers := []*httptest.Server{}
			// connections := []*websocket.Conn{}
			// defer func () {
			// 	log.Printf("handler close")
			// 	// close(h.close)
			// 	for _, s := range servers {
			// 		s.Close()
			// 	}
			// 	for _, ws := range connections {
			// 		ws.Close()
			// 	}
			// }()
			defer close(h.close)
			go h.run()
			for _, v := range tt.users {
				s, ws := newWSServer(t, h, v.userName)
				defer func(){
					s.Close()
					ws.Close()
				}()
				// servers = append(servers, s)
				// connections = append(connections, ws)
				time.Sleep(time.Second * 1)
			}
			log.Println(h.rooms)
			if len(h.rooms) != tt.roomCount {
				t.Fatalf("Expexted '%v', got '%v'", tt.roomCount, len(h.rooms))
			}
		})
	}
}

func TestCreateRoom(t *testing.T) {
	defer goleak.VerifyNone(t)
	tcs := []struct {
		name string
		users []struct {
			userName string
		}
		roomCount int
	}{
		{
			name: "1 user",
			users: []struct {
				userName string
			}{{userName: "usr1"},},
			roomCount: 1,
		},
		{
			name: "2 users",
			users: []struct {
				userName string
			}{{userName: "usr1"}, {userName: "usr2"}},
			roomCount: 1,
		},
		{
			name: "3 users",
			users: []struct {
				userName string
			}{{userName: "usr1"}, {userName: "usr2"}, {userName: "usr3"}},
			roomCount: 2,
		},
		{
			name: "10 users",
			users: []struct {
				userName string
			}{{userName: "usr1"}, {userName: "usr2"},
			{userName: "usr3"}, {userName: "usr4"},
			{userName: "usr5"}, {userName: "usr6"},
			{userName: "usr7"}, {userName: "usr8"},
			{userName: "usr9"}, {userName: "usr10"}},
			roomCount: 5,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			h := NewWshandler()
			defer close(h.close)
			go h.run()
			for _, v := range tt.users {
				s, ws := newWSServer(t, h, v.userName)
				defer s.Close()
				defer ws.Close()
				time.Sleep(time.Second * 1)
			}
			log.Println(h.rooms)
			if len(h.rooms) != tt.roomCount {
				t.Fatalf("Expexted '%v', got '%v'", tt.roomCount, len(h.rooms))
			}
		})
	}
}

func receiveWSMessage(t *testing.T, ws *websocket.Conn, count int) []outbound {
	t.Helper()

	var replys []outbound
	for i := 0; i < count; i++ {
		_, m, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("%v", err)
		}

		var reply outbound
		err = json.Unmarshal(m, &reply)
		if err != nil {
			t.Fatal(err)
		}
		replys = append(replys, reply)
	}

	return replys
}

func sendMessage(t *testing.T, ws *websocket.Conn, msg inbound) {
	t.Helper()

	m, err := json.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}

	if err := ws.WriteMessage(websocket.BinaryMessage, m); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestFillPlayerMsg(t *testing.T) {
	defer goleak.VerifyNone(t)
	tcs := []struct {
		name string
		users []struct {
			userName string
		}
		reply outbound
	}{
		{
			name: "1 user",
			users: []struct {
				userName string
			}{{userName: "usr1"},{userName: "usr2"},},
			reply: outbound{
				Users:     []struct{Id string; Name string}{
					struct{Id string; Name string}{Id: "", Name: "usr1"},
					struct{Id string; Name string}{Id: "", Name: "usr2"},
				},
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			h := NewWshandler()
			servers := []*httptest.Server{}
			connections := []*websocket.Conn{}
			defer func () {
				log.Printf("handler close")
				close(h.close)
				for _, s := range servers {
					s.Close()
				}
				for _, ws := range connections {
					ws.Close()
				}
			}()
			go h.run()
			for _, v := range tt.users {
				s, ws := newWSServer(t, h, v.userName)
				servers = append(servers, s)
				connections = append(connections, ws)

				if v.userName == "usr1" {
					continue
				}
				replys := receiveWSMessage(t, ws, 1)
				reply := replys[0]
				replyUserNames := []string{}
				for _, u := range reply.Users {
					replyUserNames = append(replyUserNames, u.Name)
				}
				ttReplyUserNames := []string{}
				for _, u := range tt.reply.Users {
					ttReplyUserNames = append(ttReplyUserNames, u.Name)
				}
				if !reflect.DeepEqual(replyUserNames, ttReplyUserNames) {
					t.Fatalf("Expexted '%v', got '%v'", tt.reply, reply)
				}
			}
		})
	}
}

func TestSetWord(t *testing.T) {
	defer goleak.VerifyNone(t)
	tcs := []struct {
		name string
		users []struct {
			userName string
			word string
		}
		reply outbound
	}{
		{
			name: "1 user",
			users: []struct {
				userName string
				word string
			}{{userName: "usr1", word: "word1"},},
			reply: outbound{
				Result:    resultOK,
				GameState: Initial,
				Thema:     "",
				Users:     []struct{Id string; Name string}{
					struct{Id string; Name string}{Id: "", Name: "usr1"},
				},
				Words:     map[string]string{"usr1": "word1"},
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			h := NewWshandler()
			servers := []*httptest.Server{}
			connections := []*websocket.Conn{}
			defer func () {
				log.Printf("handler close")
				close(h.close)
				for _, s := range servers {
					s.Close()
				}
				for _, ws := range connections {
					ws.Close()
				}
			}()
			go h.run()
			for _, v := range tt.users {
				s, ws := newWSServer(t, h, v.userName)
				servers = append(servers, s)
				connections = append(connections, ws)

				receiveWSMessage(t, ws,1)
				sendMessage(t, ws, inbound{
					Type: setWord,
					Word: v.word,
					NgChar: "",
				})
				replys := receiveWSMessage(t, ws,1)
				reply := replys[0]
				replyUserNames := []string{}
				for _, u := range reply.Users {
					replyUserNames = append(replyUserNames, u.Name)
				}
				ttReplyUserNames := []string{}
				for _, u := range tt.reply.Users {
					ttReplyUserNames = append(ttReplyUserNames, u.Name)
				}
				if !reflect.DeepEqual(replyUserNames, ttReplyUserNames) {
					t.Fatalf("Expexted '%v', got '%v'", ttReplyUserNames, replyUserNames)
				}
				if !reflect.DeepEqual(reply.Words, tt.reply.Words) {
					t.Fatalf("Expexted '%v', got '%v'", tt.reply.Words, reply.Words)
				}
			}
		})
	}
}

func TestStartGame(t *testing.T) {
	defer goleak.VerifyNone(t)
	tcs := []struct {
		name string
		users []struct {
			userName string
			word string
		}
		reply outbound
	}{
		{
			name: "2 user",
			users: []struct {
				userName string
				word string
			}{{userName: "usr1", word: "word1"},{userName: "usr2", word: "word2"}},
			reply: outbound{
				Result:    resultOK,
				GameState: GameStart,
				Thema:     "",
				Users:     []struct{Id string; Name string}{
					struct{Id string; Name string}{Id: "", Name: "usr1"},
				},
				Words:     map[string]string{"usr1": "word1", "usr2": "word2"},
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			h := NewWshandler()
			servers := []*httptest.Server{}
			connections := []*websocket.Conn{}
			defer func () {
				log.Printf("handler close")
				close(h.close)
				for _, s := range servers {
					s.Close()
				}
				for _, ws := range connections {
					ws.Close()
				}
			}()
			go h.run()
			for _, v := range tt.users[:len(tt.users)-1] {
				s, ws := newWSServer(t, h, v.userName)
				servers = append(servers, s)
				connections = append(connections, ws)

				receiveWSMessage(t, ws,1)
				sendMessage(t, ws, inbound{
					Word: v.word,
				})
				receiveWSMessage(t, ws,1)
			}
			s, ws := newWSServer(t, h, tt.users[len(tt.users)-1].userName)
			servers = append(servers, s)
			connections = append(connections, ws)
			receiveWSMessage(t, ws,1)
			sendMessage(t, ws, inbound{
				Word: tt.users[len(tt.users)-1].word,
			})
			replys := receiveWSMessage(t, ws, 2)
			reply := replys[1]
			if !(reply.GameState == tt.reply.GameState && reflect.DeepEqual(reply.Words, tt.reply.Words)) {
				t.Fatalf("Expexted '%v', got '%v'", tt.reply, reply)
			}
		})
	}
}


