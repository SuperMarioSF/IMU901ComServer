package decoder

const PRESSURE_PAYLOAD_SIZE = 10

type PressureData struct {
	pressure    int32   // unit: Pa
	altitude    int32   // unit: cm
	temperature float64 // unit: °C
}

func Decode_Pressure(payload []byte) *PressureData {
	if len(payload) < PRESSURE_PAYLOAD_SIZE {
		return nil
	}

	p0 := payload[0]
	p1 := payload[1]
	p2 := payload[2]
	p3 := payload[3]
	a0 := payload[4]
	a1 := payload[5]
	a2 := payload[6]
	a3 := payload[7]
	tL := payload[8]
	tH := payload[9]

	pressureData := new(PressureData)

	pressureData.pressure = int32(p0) + int32(p1)<<8 + int32(p2)<<16 + int32(p3)<<24
	pressureData.altitude = int32(a0) + int32(a1)<<8 + int32(a2)<<16 + int32(a3)<<24
	pressureData.temperature = float64(int16(tH)<<8|int16(tL)) / 100.0

	return pressureData
}
