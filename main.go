package main

import (
	"IMU901ComServer/serial_port"
	"IMU901ComServer/ws_server"
	"os"
	"os/signal"
)

func main() {

	serialCloseSignal := make(chan struct{}, 1)
	wsServerCloseSignal := make(chan struct{}, 1)
	wsServerQuitWaitSignal := make(chan struct{}, 1)

	go serial_port.SetupSerialPort(serialCloseSignal)
	go ws_server.StartWSServerAndWait(wsServerCloseSignal, wsServerQuitWaitSignal)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	// do shutdown
	println("starting shutdown")
	go close(serialCloseSignal)
	go close(wsServerCloseSignal)
	<-wsServerQuitWaitSignal
	println("shutdown complete")
}
