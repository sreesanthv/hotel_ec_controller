package hotel

import (
	"fmt"
	"log"

	"github.com/sreesanthv/hotel_ec_controller/config"
)

type MainCorridor struct {
	AC   *Device
	Bulb *Device
}

type SubCorridor struct {
	AC         *Device
	Bulb       *Device
	LastMotion int64
}

type Floor struct {
	MainCorridors []MainCorridor
	SubCorridors  []SubCorridor
}

type Hotel []Floor

func (h Hotel) Floor(flrId int) Floor {
	return h[flrId-1]
}

func Initialize(conf *config.Config) Hotel {
	hotel := make(Hotel, conf.NoOfFloors)
	for fId := range hotel {
		hotel[fId].MainCorridors = make([]MainCorridor, conf.MainCorridorsPerFloor)
		for mId := range hotel[fId].MainCorridors {
			hotel[fId].MainCorridors[mId].AC = NewDevice(10, true)
			hotel[fId].MainCorridors[mId].Bulb = NewDevice(5, true)
		}

		hotel[fId].SubCorridors = make([]SubCorridor, conf.SubCorridorsPerFloor)
		for sId := range hotel[fId].SubCorridors {
			hotel[fId].SubCorridors[sId].AC = NewDevice(10, true)
			hotel[fId].SubCorridors[sId].Bulb = NewDevice(5, false)
		}
	}

	return hotel
}

// Print status in console
func (h Hotel) PrintStatus() {
	for flrId, floor := range h {
		fmt.Println("Floor:", flrId+1)
		for mId, main := range floor.MainCorridors {
			fmt.Println("\tMain Corridor:", mId+1)
			fmt.Println("\t\tAC - ", main.AC.State())
			fmt.Println("\t\tBulb - ", main.Bulb.State())
		}
		for sId, sub := range floor.SubCorridors {
			fmt.Println("\tSub Corridor:", sId+1)
			fmt.Println("\t\tAC - ", sub.AC.State())
			fmt.Println("\t\tBulb - ", sub.Bulb.State())
		}

		fmt.Println("\tPower consumption:", floor.PowerConsumption(), "Units")
	}
}

// get floor power consumption
func (floor Floor) PowerConsumption() int {
	var power int

	for _, main := range floor.MainCorridors {
		power += main.AC.InstantPowerConsumption()
		power += main.Bulb.InstantPowerConsumption()
	}
	for _, sub := range floor.SubCorridors {
		power += sub.AC.InstantPowerConsumption()
		power += sub.Bulb.InstantPowerConsumption()
	}

	return power
}

// to turn on / off Main Cor - Bulb
func (floor Floor) ManageSubBulb(subId int, doTurnOn bool) {
	if doTurnOn {
		floor.SubCorridors[subId-1].Bulb.On()
	} else {
		floor.SubCorridors[subId-1].Bulb.Off()
	}
}

// to turn on / off Main Cor - AC
func (floor Floor) ManageSubAC(subId int, doTurnOn bool) {
	if doTurnOn {
		floor.SubCorridors[subId-1].AC.On()
	} else {
		floor.SubCorridors[subId-1].AC.Off()
	}
}

// optimise power usage
func (f Floor) OptimisePowerUsage() {
	mainCount := len(f.MainCorridors)
	subCount := len(f.SubCorridors)
	power := f.PowerConsumption()

	if (mainCount*15 + subCount*10) < power {
		log.Println("High power consumption, turning off Sub Corridor - AC")
		f.ManageAllSubAC(false)
	} else {
		// power consumption - back to normal
		f.ManageAllSubAC(true)
	}
}

// to on/off all ACs under
func (f Floor) ManageAllSubAC(doTurnOn bool) {
	for sId, _ := range f.SubCorridors {
		f.ManageSubAC(sId+1, doTurnOn)
	}
}

func (h Hotel) SetSubLastMotion(fId, subId int, t int64) {
	h[fId-1].SubCorridors[subId-1].LastMotion = t
}

func (h Hotel) GetSubLastMotion(fId, subId int) int64 {
	return h[fId-1].SubCorridors[subId-1].LastMotion
}
