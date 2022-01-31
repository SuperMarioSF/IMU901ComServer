package decoder

const PORTSTAT_PAYLOAD_SIZE = 8

type PortStatData struct {
	d0 uint16
	d1 uint16
	d2 uint16
	d3 uint16
}

type PortStatVoltage struct {
	Ud0 float64
	Ud1 float64
	Ud2 float64
	Ud3 float64
}

func (portStat *PortStatData) ToAdcVoltage() PortStatVoltage {
	return PortStatVoltage{
		Ud0: (float64(portStat.d0) / 4095) * 3.3, // unit: volt
		Ud1: (float64(portStat.d1) / 4095) * 3.3, // unit: volt
		Ud2: (float64(portStat.d2) / 4095) * 3.3, // unit: volt
		Ud3: (float64(portStat.d3) / 4095) * 3.3, // unit: volt
	}
}

func Decode_PortStat(payload []byte) *PortStatData {
	if len(payload) != PORTSTAT_PAYLOAD_SIZE {
		return nil
	}

	d0L := payload[0]
	d0H := payload[1]
	d1L := payload[2]
	d1H := payload[3]
	d2L := payload[4]
	d2H := payload[5]
	d3L := payload[6]
	d3H := payload[7]

	portstatData := new(PortStatData)

	portstatData.d0 = uint16(d0H)<<8 | uint16(d0L)
	portstatData.d1 = uint16(d1H)<<8 | uint16(d1L)
	portstatData.d2 = uint16(d2H)<<8 | uint16(d2L)
	portstatData.d3 = uint16(d3H)<<8 | uint16(d3L)

	return portstatData
}
