package serial_port

import (
	"IMU901ComServer/decoder"
	"flag"
	"go.bug.st/serial"
	"io"
	"log"
)

var ComPort = flag.String("device", "COM6", "serial device name")

type SerialPortControl struct {
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SetupSerialPort(closeSignal chan struct{}) {
	log.Println("Setting up serial port for IMU901")
	if !stringInSlice(*ComPort, GetAvailableSerialPorts()) {
		log.Fatalf("No such serial port: %s", *ComPort)
	}

	port := OpenSerialPort("COM6", GetDefaultMode())
	defer func(port serial.Port) {
		log.Println("Closing serial port...")
		err := port.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(port) // shutdown if something went wrong
	log.Println("Serial port setup complete.")
	mainLoop(port, closeSignal)
	log.Println("serial decoder shutdown complete.")
}

func GetAvailableSerialPorts() []string {
	log.Println("Discovering available serial ports...")
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		log.Printf("Found port: %v\n", port)
	}
	return ports
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
	log.Println("Opening serial port...")
	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serial port opened.")
	return port
}

func mainLoop(port serial.Port, closeSignal chan struct{}) {
	var err error
	log.Println("IMU901 device initialization...")
	gyroRange, errGyro := decoder.GetDeviceGyroRange(port)
	if errGyro != nil {
		log.Fatal(errGyro) // quit here
	}
	log.Println("IMU901 start polling data.")
	for {
		err = decoder.DecodeStart(port, gyroRange)

		if err == io.EOF {
			return
		}

		select {
		case <-closeSignal:
			log.Println("Received close signal. Closing serial port...")
			return
		default:
			continue
		}
	}
}
