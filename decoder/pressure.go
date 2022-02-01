package decoder

import "encoding/json"

const PressurePayloadSize = 10

type PressureData struct {
	Pressure    int32   `json:"pressure"`    // unit: Pa
	Altitude    int32   `json:"altitude"`    // unit: cm
	Temperature float64 `json:"temperature"` // unit: °C
}

func (data *PressureData) ToJson() (string, error) {
	j, e := json.Marshal(decodedDataStruct{
		Source: "pressure",
		Data:   *data,
	})
	return string(j), e
}

func DecodePressure(payload []byte) *PressureData {
	if len(payload) < PressurePayloadSize {
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

	pressureData.Pressure = int32(p0) + int32(p1)<<8 + int32(p2)<<16 + int32(p3)<<24
	pressureData.Altitude = int32(a0) + int32(a1)<<8 + int32(a2)<<16 + int32(a3)<<24
	pressureData.Temperature = float64(int16(tH)<<8|int16(tL)) / 100.0

	return pressureData
}
