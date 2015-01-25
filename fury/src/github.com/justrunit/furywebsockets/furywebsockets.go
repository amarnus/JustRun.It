/*
	Fury service websocket handler to
		LINT
		RUN
*/

package furywebsockets

import (
	"log"
	"encoding/json"
	"net/http"
	"io"
	"bufio"
	"regexp"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
	ReadBufferSize: 1024,
	WriteBufferSize: 1024};

type websocketmessage struct {
	Id string       `json:"id"`
	Data string     `json:"data"`
}

// Websocket connection and corresponding channels
// These channels read from the output readers
// above and pass data out via websocket conn
type connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // Buffered channel of outbound messages.
    send chan []byte

	// Session ID to monitor
	sid string
}

var SidWebsocketExists = make(map[string]int)
var SidToWebsocketChannel = make(map[string]chan []byte)
var SidToOperation = make(map[string]string)
var iohub = make(chan websocketmessage, 100);
var iohubInitialized = 0

// Global listener for all io events
func iohubListen() {
	for msg := range iohub {
		sid := msg.Id
		if SidWebsocketExists[sid] != 0 {
			SidToWebsocketChannel[sid] <- []byte(msg.Data)
		} else {
			log.Println("No websocket client for " + sid + ": " + msg.Data )
		}
	}
}

func ReaderToChannel(sid string, sout io.Reader, serr io.Reader) {
	numListeners := 2
	log.Println("Wiring up terminal io of session " + sid + " to its websocket")
	scanReaderFn := func(sio io.Reader) {
		s := bufio.NewScanner(sio)
		s.Split(bufio.ScanLines)

		cnt := 0

		for s.Scan() {
			iohub <- websocketmessage{
				Id: sid,
				Data: string(s.Text())}
			cnt++
		}
		numListeners--;
		if numListeners == 0 {
			iohub <- websocketmessage{
				Id: sid,
				Data: "op-complete"}
			delete( SidToOperation, sid )
		}
	}

	go scanReaderFn(sout)
	go scanReaderFn(serr)

	return
}

func (c *connection) reader() {
	for {
        _, message, err := c.ws.ReadMessage()
        if err != nil {
			log.Println( "Error reading websocket " + err.Error() )
            break
        }
		log.Println( "Websocket message: " + string(message) )

		// Parse incoming message from websocket client into go map
		var msgMap map[string]interface{}
		json.Unmarshal(message, &msgMap)

		// If sid not set, set it and start monitor
		if msgMap[ "id" ] == nil && c.sid == "" {

			c.ws.WriteMessage(websocket.TextMessage, []byte("{\"status\":0,\"error\":\"Send sid to monitor\"}"))

		} else if msgMap[ "id" ] != nil && c.sid == "" {

			c.sid = msgMap[ "id" ].(string)
			log.Println("Creating websocket connection for sid " + c.sid)
			SidToWebsocketChannel[ c.sid ] = c.send
			SidWebsocketExists[ c.sid ] = 1
		}
    }
    c.ws.Close()
}

func (c *connection) writer() {

	// Wait for incoming message from reader on channel
    for message := range c.send {

		// Retrieve sid from current websocket connection 
		// struct instance and message from current channel
		msg := websocketmessage{
			Id: c.sid,
			Data: string(message)}

		// Construct JSON from struct
		b, _ := json.Marshal(&msg)
		bstr := string(b)
		match, _ := regexp.MatchString("Error mounting", bstr)
		if match {
			bstr = "An error ocurred in the execution context. Please run again later"
			b = []byte(b)
		}

		// Write the json object to websocket
		err := c.ws.WriteMessage(websocket.TextMessage, b)

		if err != nil {
			break
		}
    }

	log.Println("Closing websocket connection for " + c.sid)
	delete( SidToWebsocketChannel, c.sid )
	delete( SidWebsocketExists, c.sid )
    c.ws.Close()
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	log.Println("Creating empty websocket connection")
	if err != nil {
		// if _, ok := err.(websocket.HandshakeError); !ok {
		// 	log.Println(err)
		// }
		log.Println(err)
		return
	}

	if iohubInitialized == 0 {
		go iohubListen()
		iohubInitialized = 1
	}

	c := &connection{
		send: make(chan []byte, 256),
		ws: ws,
		sid: ""}

	go c.writer()
	c.reader()
}

