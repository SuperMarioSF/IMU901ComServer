package decoder

import (
	"fmt"
	"go.bug.st/serial"
	"io"
)

const EULER uint8 = 0x01
const QUATERNION uint8 = 0x02
const GYRO uint8 = 0x03
const MAGNETIC uint8 = 0x04
const PRESSURE uint8 = 0x05
const PORTSTAT uint8 = 0x06

func readChar(port serial.Port, size int) ([]byte, error) {
	charArr := make([]byte, size)
	n, err := port.Read(charArr)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, io.EOF
	}
	return charArr, nil
}

func DecodeStart(port serial.Port, gyroRange DeviceGyroRange) error {
	charABBytearray, errA := readChar(port, 1)
	if errA != nil {
		return errA
	}
	if charABBytearray[0] != 0x55 {
		return nil // ignore desync case
	}
	charBBBytearray, errB := readChar(port, 1)
	if errB != nil {
		return errB
	}
	if charBBBytearray[0] != 0x55 {
		return nil // ignore desync case
	}

	// from now on, we have synced on frame header.
	frameIdBytearray, errC := readChar(port, 1)
	if errC != nil {
		return errC
	}

	frameId := frameIdBytearray[0]

	dataLengthByteArray, errD := readChar(port, 1)
	if errD != nil {
		return errD
	}
	dataLength := int(dataLengthByteArray[0])

	dataBlock, errE := readChar(port, dataLength)
	if errE != nil {
		return errE
	}

	checksumBytearray, errF := readChar(port, 1)
	if errF != nil {
		return errF
	}

	wholePayload := make([]byte, 1+1+1+1+dataLength+1)
	copy(wholePayload[0:1], charABBytearray)
	copy(wholePayload[1:2], charBBBytearray)
	copy(wholePayload[2:3], frameIdBytearray)
	copy(wholePayload[3:4], dataLengthByteArray)
	copy(wholePayload[4:4+dataLength], dataBlock)
	copy(wholePayload[4+dataLength:5+dataLength], checksumBytearray)

	if !checksumCheck(wholePayload) {
		return nil // checksum failed, skip frame
	}

	switch frameId {
	case EULER:
		eulerData := DecodeEuler(dataBlock)
		if eulerData == nil {
			break
		}
		fmt.Printf("EULER: \t\troll=%3.3f\tpitch=%3.3f\tyaw=%3.3f\n",
			eulerData.roll, eulerData.pitch, eulerData.yaw)
		break
	case QUATERNION:
		quaternionData := DecodeQuatenion(dataBlock)
		if quaternionData == nil {
			break
		}
		fmt.Printf("QUATERNION: \tq0=%3.3f\tq1=%3.3f\tq2=%3.3f\tq3=%3.3f\n",
			quaternionData.q0, quaternionData.q1, quaternionData.q2, quaternionData.q3)
		break
	case GYRO:
		gyroData := DecodeGyro(dataBlock, gyroRange)
		if gyroData == nil {
			break
		}
		fmt.Printf("ACCEL: \t\taccX=%3.3f\taccY=%3.3f\taccZ=%3.3f\nGYRO: \t\tgyroX=%3.3f\tgyroY=%3.3f\tgyroZ=%3.3f\n",
			gyroData.accX, gyroData.accY, gyroData.accZ, gyroData.gyroX, gyroData.gyroY, gyroData.gyroZ)
		break
	case MAGNETIC:
		magneticData := DecodeMagnetic(dataBlock)
		if magneticData == nil {
			break
		}
		fmt.Printf("MAGNETIC: \tmagX=%3.3f\tmagY=%3.3f\tmagZ=%3.3f\ttemp=%3.3f°C\n",
			magneticData.magnetX, magneticData.magnetY, magneticData.magnetZ, magneticData.temperature)
		break
	case PRESSURE:
		pressureData := DecodePressure(dataBlock)
		if pressureData == nil {
			break
		}
		fmt.Printf("PRESSURE: \tpre=%dPa\talt=%dcm\ttemp=%3.3f°C\n",
			pressureData.pressure, pressureData.altitude, pressureData.temperature)
		break
	case PORTSTAT:
		portStatusData := DecodePortstat(dataBlock)
		if portStatusData == nil {
			break
		}
		portVoltData := portStatusData.ToAdcVoltage()
		fmt.Printf("PORTSTAT: \td0=%6d\td1=%6d\td2=%6d\td3=%6d\nPORTVOLT: \tUd0=%.3fv\tUd1=%.3fv\tUd2=%.3fv\tUd3=%.3fv\n",
			portStatusData.d0, portStatusData.d1, portStatusData.d2, portStatusData.d3,
			portVoltData.Ud0, portVoltData.Ud1, portVoltData.Ud2, portVoltData.Ud3)
		break
	default:
		fmt.Printf("UKNOWN: FrameID=0x%.2X\n", frameId)
		return nil // unknown data type
	}

	return nil // end of section
}

func checksumCheck(payload []byte) bool {
	checksum := payload[len(payload)-1]
	sumPayload := payload[0 : len(payload)-1]
	var actualSum uint8 = 0x00

	for _, element := range sumPayload {
		actualSum += element
	}

	if checksum == actualSum {
		return true
	} else {
		return false
	}
}
