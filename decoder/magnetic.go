package decoder

import "encoding/json"

const MagneticPayloadSize int = 8

type MagneticData struct {
	MagnetX     float64 `json:"magnet_x"`
	MagnetY     float64 `json:"magnet_y"`
	MagnetZ     float64 `json:"magnet_z"`
	Temperature float64 `json:"temperature"`
}

func (data *MagneticData) ToJson() (string, error) {
	j, e := json.Marshal(decodedDataStruct{
		Source: "magnetic",
		Data:   *data,
	})
	return string(j), e

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

	magneticData.MagnetX = float64(int16(mXH)<<8|int16(mXL)) / 1000.0
	magneticData.MagnetY = float64(int16(mYH)<<8|int16(mYL)) / 1000.0
	magneticData.MagnetZ = float64(int16(mZH)<<8|int16(mZL)) / 1000.0
	magneticData.Temperature = float64(int16(tH)<<8|int16(tL)) / 100.0

	return magneticData
}
