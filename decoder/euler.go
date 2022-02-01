package decoder

import (
	json "encoding/json"
)

const EulerPayloadSize int = 6 //bytes

type EulerData struct {
	Roll  float64 `json:"roll"`
	Pitch float64 `json:"pitch"`
	Yaw   float64 `json:"yaw"`
}

func (data *EulerData) ToJson() (string, error) {
	j, e := json.Marshal(decodedDataStruct{
		Source: "euler",
		Data:   *data,
	})
	return string(j), e
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
	eulerData.Roll = float64(int16((uint16(rollH)<<8)|uint16(rollL))) / 32768 * 180
	eulerData.Pitch = float64(int16((uint16(pitchH)<<8)|uint16(pitchL))) / 32768 * 180
	eulerData.Yaw = float64(int16((uint16(yawH)<<8)|uint16(yawL))) / 32768 * 180

	return eulerData
}
