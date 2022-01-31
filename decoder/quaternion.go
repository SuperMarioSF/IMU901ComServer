package decoder

const QuaternionPayloadSize int = 8

type QuaternionData struct {
	q0 float64
	q1 float64
	q2 float64
	q3 float64
}

func DecodeQuatenion(payload []byte) *QuaternionData {
	if len(payload) != QuaternionPayloadSize {
		return nil
	}

	q0L := payload[0]
	q0H := payload[1]
	q1L := payload[2]
	q1H := payload[3]
	q2L := payload[4]
	q2H := payload[5]
	q3L := payload[6]
	q3H := payload[7]

	quaternionData := new(QuaternionData)
	quaternionData.q0 = float64(int16((uint16(q0H)<<8)|uint16(q0L))) / 32768
	quaternionData.q1 = float64(int16((uint16(q1H)<<8)|uint16(q1L))) / 32768
	quaternionData.q2 = float64(int16((uint16(q2H)<<8)|uint16(q2L))) / 32768
	quaternionData.q3 = float64(int16((uint16(q3H)<<8)|uint16(q3L))) / 32768

	return quaternionData
}
