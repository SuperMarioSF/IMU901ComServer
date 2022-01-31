package decoder

const EulerPayloadSize int = 6 //bytes

type EulerData struct {
	roll  float64
	pitch float64
	yaw   float64
}

func DecodeEuler(payload []byte) *EulerData {
	if len(payload) != EulerPayloadSize {
		return nil // incorrect payload.
	}
	rollL := payload[0]
	rollH := payload[1]
	pitchL := payload[2]
	pitchH := payload[3]
	yawL := payload[4]
	yawH := payload[5]

	eulerData := new(EulerData)
	eulerData.roll = float64(int16((uint16(rollH)<<8)|uint16(rollL))) / 32768 * 180
	eulerData.pitch = float64(int16((uint16(pitchH)<<8)|uint16(pitchL))) / 32768 * 180
	eulerData.yaw = float64(int16((uint16(yawH)<<8)|uint16(yawL))) / 32768 * 180

	return eulerData
}
