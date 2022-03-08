package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type Client struct {
	wsHandler *wsHandler
	room      *Room
	id        string
	name      string
	Word      string
	conn      *websocket.Conn
	send      chan []byte
	close     chan struct{}
	wordState WordState
}

type Clients map[string]*Client

func (c Clients) values() []*Client {
	vs := []*Client{}
	for _, v := range c {
		vs = append(vs, v)
	}
	return vs
}

func (c *Client) write() {
	// c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	for {
		select {
		case <-c.close:
			log.Printf("close client.write()")
			return
		case msg, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (c *Client) read() {
	c.conn.SetReadLimit(maxMessageSize)
	// c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	defer func() {
		log.Println(c.name, " leave")
		c.room.gameStop()
		close(c.close)
	}()
	for {
		if _, msg, err := c.conn.ReadMessage(); err == nil {
			c.room.send <- struct {
				from *Client
				body []byte
			}{from: c, body: msg}
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
