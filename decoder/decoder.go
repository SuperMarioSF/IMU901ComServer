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

func readchar(port serial.Port, size int) ([]byte, error) {
	char_arr := make([]byte, size)
	n, err := port.Read(char_arr)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, io.EOF
	}
	return char_arr, nil
}

func Decode_Start(port serial.Port) error {
	char_a_bytearr, errA := readchar(port, 1)
	if errA != nil {
		return errA
	}
	if char_a_bytearr[0] != 0x55 {
		return nil // ignore unsync case
	}
	char_b_bytearr, errB := readchar(port, 1)
	if errB != nil {
		return errB
	}
	if char_b_bytearr[0] != 0x55 {
		return nil // ignore unsync case
	}

	// from now on, we have synced on frame header.
	frame_id_bytearr, errC := readchar(port, 1)
	if errC != nil {
		return errC
	}

	framr_id := uint8(frame_id_bytearr[0])

	data_length_byte, errD := readchar(port, 1)
	if errD != nil {
		return errD
	}
	data_length := int(uint8(data_length_byte[0]))

	data_block, errE := readchar(port, data_length)
	if errE != nil {
		return errE
	}

	checksum_bytearr, errF := readchar(port, 1)
	if errF != nil {
		return errF
	}

	whole_payload := make([]byte, 1+1+1+1+data_length+1)
	copy(whole_payload[0:1], char_a_bytearr)
	copy(whole_payload[1:2], char_b_bytearr)
	copy(whole_payload[2:3], frame_id_bytearr)
	copy(whole_payload[3:4], data_length_byte)
	copy(whole_payload[4:4+data_length], data_block)
	copy(whole_payload[4+data_length:5+data_length], checksum_bytearr)

	if !checksum_check(whole_payload) {
		return nil // checksum failed, skip frame
	}

	switch framr_id {
	case EULER:
		euler_data := Decode_Euler(data_block)
		if euler_data == nil {
			break
		}
		fmt.Printf("EULER: roll=%3.3f\tpitch=%3.3f\tyaw=%3.3f\n",
			euler_data.roll, euler_data.pitch, euler_data.yaw)
		break
	case QUATERNION:
		//println("QUATERNION")
		break
	case GYRO:
		//println("GYRO")
		break
	case MAGNETIC:
		//println("MAGNETIC")
		break
	case PRESSURE:
		//println("PRESSURE")
		break
	case PORTSTAT:
		//println("PORTSTAT")
		break
	default:
		println("UKNOWN")
		return nil // unknown data type
	}

	return nil // end of section
}

func checksum_check(payload []byte) bool {
	checksum := payload[len(payload)-1]
	sum_payload := payload[0 : len(payload)-1]
	var actual_sum uint8 = 0x00

	for _, element := range sum_payload {
		actual_sum += element
	}

	if checksum == actual_sum {
		return true
	} else {
		return false
	}
}