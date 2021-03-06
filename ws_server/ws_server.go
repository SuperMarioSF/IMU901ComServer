package ws_server

import (
	"context"
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
var serverCtl *http.Server = nil
var ServerReady bool = false
var GlobalMessagehub *Hub = newHub()

func StartWSServerAndWait(closeSignal chan struct{}, quitWaitChan chan struct{}) {
	log.Println("Starting message hub")
	hub := GlobalMessagehub
	go hub.run()
	log.Println("Starting websocket handler")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	serverChan := make(chan *http.Server, 1)
	waitChan := make(chan struct{}, 1)
	log.Printf("Starting websocket listener at '%s'", *addr)
	go getServerAndListen(*addr, nil, serverChan, waitChan)

	serverCtl = <-serverChan
	ServerReady = true
	defer internalShutdownServerAndWait(serverCtl)

	go func(closeSignal chan struct{}, serverCtl *http.Server) {
		if closeSignal != nil {
			<-closeSignal
			log.Println("Received close signal, shutting down websocket server...")
			go internalShutdownServerAndWait(serverCtl)
			log.Println("Wait for remaining connections to close")
		}
	}(closeSignal, serverCtl)
	<-waitChan
	log.Println("websocket server shutdown complete")

	if quitWaitChan != nil {
		close(quitWaitChan)
	}
}

func ShutdownServerAndWait() {
	go internalShutdownServerAndWait(serverCtl)
}

func getServerAndListen(addr string, handler http.Handler, server_chan chan<- *http.Server, waitChan chan<- struct{}) {
	server := &http.Server{Addr: addr, Handler: handler}
	server_chan <- server
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			ServerReady = false
			close(waitChan)
			return // handled elsewhere.
		}
		log.Fatal("ListenAndServe: ", err)
	}
}

func internalShutdownServerAndWait(serverCtl *http.Server) {
	ServerReady = false
	_shutdownServer(serverCtl)
}

func _shutdownServer(serverCtl *http.Server) {
	if err := serverCtl.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}
}

var line_number int = 0

func ProcessUpdateEvent(event string) {
	if !ServerReady {
		return
	}
	if len(GlobalMessagehub.clients) == 0 {

		return
	}
	GlobalMessagehub.broadcast <- []byte(event)
	line_number++
}
