package hotel

type Device struct {
	isOn             bool
	powerConsumption int
}

func (d *Device) On() {
	d.isOn = true
}

func (d *Device) Off() {
	d.isOn = false
}

func (d *Device) State() string {
	st := "Off"
	if d.isOn {
		st = "On"
	}
	return st
}

// get instant power usage
func (d *Device) InstantPowerConsumption() int {
	p := 0
	if d.isOn {
		p = d.powerConsumption
	}

	return p
}

func NewDevice(power int, isOn bool) *Device {
	return &Device{
		isOn:             isOn,
		powerConsumption: power,
	}
}
