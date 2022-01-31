package decoder

const MagneticPayloadSize int = 8

type MagneticData struct {
	magnetX     float64
	magnetY     float64
	magnetZ     float64
	temperature float64
}

func DecodeMagnetic(payload []byte) *MagneticData {
	if len(payload) != MagneticPayloadSize {
		return nil
	}

	mXL := payload[0]
	mXH := payload[1]
	mYL := payload[2]
	mYH := payload[3]
	mZL := payload[4]
	mZH := payload[5]
	tL := payload[6]
	tH := payload[7]

	magneticData := new(MagneticData)

	magneticData.magnetX = float64(int16(mXH)<<8|int16(mXL)) / 1000.0
	magneticData.magnetY = float64(int16(mYH)<<8|int16(mYL)) / 1000.0
	magneticData.magnetZ = float64(int16(mZH)<<8|int16(mZL)) / 1000.0
	magneticData.temperature = float64(int16(tH)<<8|int16(tL)) / 100.0

	return magneticData
}
