package serial_port

import (
	"IMU901ComServer/decoder"
	"fmt"
	"go.bug.st/serial"
	"io"
	"log"
)

type SerialPortControl struct {
}

func SetupSerialPort(closeSignal chan struct{}) {
	GetAvailableSerialPorts()
	port := OpenSerialPort("COM6", GetDefaultMode())
	defer func(port serial.Port) {
		err := port.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(port) // shutdown if something went wrong
	mainLoop(port, closeSignal)
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

func OpenSerialPort(portName string, mode *serial.Mode) serial.Port {
	//mode := &serial.Mode{
	//	BaudRate: 115200,
	//}
	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatal(err)
	}

	return port
}

func mainLoop(port serial.Port, closeSignal chan struct{}) {
	var err error
	gyroRange, errGyro := decoder.GetDeviceGyroRange(port)
	if errGyro != nil {
		log.Fatal(errGyro) // quit here
	}
	for {
		err = decoder.DecodeStart(port, gyroRange)

		if err == io.EOF {
			return
		}

		select {
		case <-closeSignal:
			return
		default:
			continue
		}
	}
}
