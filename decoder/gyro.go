package decoder

import "go.bug.st/serial"

const GYRO_PAYLOAD_SIZE = 12

type GyroData struct {
	accX  float64
	accY  float64
	accZ  float64
	gyroX float64
	gyroY float64
	gyroZ float64
}

type DeviceGyroRange struct {
	accelRange int
	gyroRange  int
}

func GetDeviceGyroRange(port serial.Port) (DeviceGyroRange, error) {
	// placeholder data
	gyroRange := DeviceGyroRange{
		accelRange: 1,
		gyroRange:  2000,
	}
	return gyroRange, nil
}

func Decode_Gyro(payload []byte, gyroRange DeviceGyroRange) *GyroData {
	if len(payload) != GYRO_PAYLOAD_SIZE {
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
	gyroData.accX = float64(int16((uint16(aXH)<<8)|uint16(aXL))) / 32768 * float64(gyroRange.accelRange)
	gyroData.accY = float64(int16((uint16(aYH)<<8)|uint16(aYL))) / 32768 * float64(gyroRange.accelRange)
	gyroData.accZ = float64(int16((uint16(aZH)<<8)|uint16(aZL))) / 32768 * float64(gyroRange.accelRange)
	gyroData.gyroX = float64(int16((uint16(gXH)<<8)|uint16(gXL))) / 32768 * float64(gyroRange.gyroRange)
	gyroData.gyroY = float64(int16((uint16(gYH)<<8)|uint16(gYL))) / 32768 * float64(gyroRange.gyroRange)
	gyroData.gyroZ = float64(int16((uint16(gZH)<<8)|uint16(gZL))) / 32768 * float64(gyroRange.gyroRange)

	return gyroData
}
