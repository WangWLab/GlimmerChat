package main

import (
	"flag"
	"fmt"
	"net/http"

	"GlimmerChat/log"
	"time"

	"github.com/gorilla/websocket"
)

var serverAddr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.RootLogger.Errorf("upgrade: %s", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.RootLogger.Errorf("read: %s", err)
			break
		}
		log.RootLogger.Infof("recv: %s", message)
		err = c.WriteMessage(mt, []byte(fmt.Sprintf("message received, time: %s", time.Now())))
		if err != nil {
			log.RootLogger.Errorf("write: %s", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.RootLogger.Warning("test server started")
	http.HandleFunc("/echo", echo)
	log.RootLogger.Fatal(http.ListenAndServe(*serverAddr, nil))
}
