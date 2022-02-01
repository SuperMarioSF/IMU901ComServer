package decoder

import (
	"encoding/json"
	"go.bug.st/serial"
)

const GyroPayloadSize = 12

type GyroData struct {
	AccX  float64 `json:"acc_x"`
	AccY  float64 `json:"acc_y"`
	AccZ  float64 `json:"acc_z"`
	GyroX float64 `json:"gyro_x"`
	GyroY float64 `json:"gyro_y"`
	GyroZ float64 `json:"gyro_z"`
}

type DeviceGyroRange struct {
	accelRange int
	gyroRange  int
}

func (data *GyroData) ToJson() (string, error) {
	j, e := json.Marshal(decodedDataStruct{
		Source: "gyro",
		Data:   *data,
	})
	return string(j), e

}

func GetDeviceGyroRange(port serial.Port) (DeviceGyroRange, error) {
	// placeholder Data
	gyroRange := DeviceGyroRange{
		accelRange: 1,
		gyroRange:  2000,
	}
	return gyroRange, nil
}

func DecodeGyro(payload []byte, gyroRange DeviceGyroRange) *GyroData {
	if len(payload) != GyroPayloadSize {
		return nil
	}

	aXL := float64(payload[0])
	aXH := float64(payload[1])
	aYL := float64(payload[2])
	aYH := float64(payload[3])
	aZL := float64(payload[4])
	aZH := float64(payload[5])
	gXL := float64(payload[6])
	gXH := float64(payload[7])
	gYL := float64(payload[8])
	gYH := float64(payload[9])
	gZL := float64(payload[10])
	gZH := float64(payload[11])

	gyroData := new(GyroData)
	gyroData.AccX = float64(int16((uint16(aXH)<<8)|uint16(aXL))) / 32768 * float64(gyroRange.accelRange)
	gyroData.AccY = float64(int16((uint16(aYH)<<8)|uint16(aYL))) / 32768 * float64(gyroRange.accelRange)
	gyroData.AccZ = float64(int16((uint16(aZH)<<8)|uint16(aZL))) / 32768 * float64(gyroRange.accelRange)
	gyroData.GyroX = float64(int16((uint16(gXH)<<8)|uint16(gXL))) / 32768 * float64(gyroRange.gyroRange)
	gyroData.GyroY = float64(int16((uint16(gYH)<<8)|uint16(gYL))) / 32768 * float64(gyroRange.gyroRange)
	gyroData.GyroZ = float64(int16((uint16(gZH)<<8)|uint16(gZL))) / 32768 * float64(gyroRange.gyroRange)

	return gyroData
}
