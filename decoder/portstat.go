package decoder

import "encoding/json"

const PortstatPayloadSize = 8

type PortStatData struct {
	D0 uint16 `json:"d0"`
	D1 uint16 `json:"d1"`
	D2 uint16 `json:"d2"`
	D3 uint16 `json:"d3"`
}

type PortStatVoltage struct {
	Ud0 float64 `json:"ud0"`
	Ud1 float64 `json:"ud1"`
	Ud2 float64 `json:"ud2"`
	Ud3 float64 `json:"ud3"`
}

type FullPortStat struct {
	PortStat   PortStatData    `json:"port_stat"`
	ADCVoltage PortStatVoltage `json:"adc_voltage"`
}

func (data *FullPortStat) ToJson() (string, error) {
	j, e := json.Marshal(decodedDataStruct{
		Source: "full_port_stat",
		Data:   *data,
	})
	return string(j), e
}

func (portStat *PortStatData) ToJson() (string, error) {
	j, e := json.Marshal(decodedDataStruct{
		Source: "port_stat",
		Data:   *portStat,
	})
	return string(j), e
}

func (portStat *PortStatData) ToFullPortStat() *FullPortStat {
	data := new(FullPortStat)
	data.PortStat = *portStat
	data.ADCVoltage = portStat.ToAdcVoltage()
	return data
}

func (portStat *PortStatData) ToAdcVoltage() PortStatVoltage {
	return PortStatVoltage{
		Ud0: (float64(portStat.D0) / 4095) * 3.3, // unit: volt
		Ud1: (float64(portStat.D1) / 4095) * 3.3, // unit: volt
		Ud2: (float64(portStat.D2) / 4095) * 3.3, // unit: volt
		Ud3: (float64(portStat.D3) / 4095) * 3.3, // unit: volt
	}
}

func DecodePortstat(payload []byte) *PortStatData {
	if len(payload) != PortstatPayloadSize {
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

	portstatData.D0 = uint16(d0H)<<8 | uint16(d0L)
	portstatData.D1 = uint16(d1H)<<8 | uint16(d1L)
	portstatData.D2 = uint16(d2H)<<8 | uint16(d2L)
	portstatData.D3 = uint16(d3H)<<8 | uint16(d3L)

	return portstatData
}
