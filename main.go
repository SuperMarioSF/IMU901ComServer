package main

import (
	"IMU901ComServer/serial_port"
	"IMU901ComServer/ws_server"
	"flag"
	"log"
	"os"
	"os/signal"
)

func main() {
	flag.Parse()

	serialCloseSignal := make(chan struct{}, 1)
	wsServerCloseSignal := make(chan struct{}, 1)
	wsServerQuitWaitSignal := make(chan struct{}, 1)

	log.Println("Starting serial data decoder...")
	go serial_port.SetupSerialPort(serialCloseSignal)

	log.Println("Starting websocket server...")
	go ws_server.StartWSServerAndWait(wsServerCloseSignal, wsServerQuitWaitSignal)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	log.Println("Gracefully closing serial port and websocket server")
	// do shutdown
	log.Println("")
	go close(serialCloseSignal)
	go close(wsServerCloseSignal)
	<-wsServerQuitWaitSignal
	log.Println("quitting...")
}
