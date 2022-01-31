package main

import (
	"IMU901ComServer/decoder"
	"fmt"
	"go.bug.st/serial"
	"io"
	"log"
)

func main() {
	GetAvailableSerialPorts()
	main_loop(OpenSerialPort("COM6", GetDefaultMode()))
}

func GetAvailableSerialPorts() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}

func GetDefaultMode() *serial.Mode {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	return mode
}

func OpenSerialPort(port_name string, mode *serial.Mode) serial.Port {
	//mode := &serial.Mode{
	//	BaudRate: 115200,
	//}
	port, err := serial.Open(port_name, mode)
	if err != nil {
		log.Fatal(err)
	}

	return port
}

func main_loop(port serial.Port) {
	var err error
	for {
		err = decoder.Decode_Start(port)

		if err == io.EOF {
			return
		}
	}
}

//func print_char(port serial.Port) error {
//	char := make([]byte, 1)
//	n, err := port.Read(char)
//	if err != nil {
//		return err
//	}
//	if n == 0 {
//		println("Not reading character...")
//		return nil
//	}
//
//	fmt.Printf("0x%.2x (%d)\n", char[0], char[0])
//	return nil
//}
