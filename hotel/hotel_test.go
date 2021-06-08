package hotel

import (
	"encoding/json"
	"testing"

	"github.com/sreesanthv/hotel_ec_controller/config"
)

func TestHotelDefaultState(t *testing.T) {
	conf := &config.Config{2, 1, 2}
	h := Initialize(conf)

	if h[0].SubCorridors[0].Bulb.isOn {
		a1, _ := json.Marshal(h)
		t.Error("Sub corridor bulb is on", string(a1))
	}

	if h[1].SubCorridors[0].AC.isOn == false {
		a1, _ := json.Marshal(h)
		t.Error("Sub corridor AC is off", string(a1))
	}

	if h[1].MainCorridors[0].AC.isOn == false {
		a1, _ := json.Marshal(h)
		t.Error("Sub corridor AC is off", string(a1))
	}

	if h[1].MainCorridors[0].Bulb.isOn == false {
		a1, _ := json.Marshal(h)
		t.Error("Sub corridor AC is off", string(a1))
	}

	if p := h[1].PowerConsumption(); p != 35 {
		t.Errorf("Wrong power consumption. Expected: %d, Got %d", 35, p)
	}
}

func TestManageDevice(t *testing.T) {
	conf := &config.Config{2, 1, 2}
	h := Initialize(conf)

	h[0].ManageSubAC(1, false)
	if h[0].SubCorridors[0].AC.isOn {
		t.Error("Failed to turn off Sub AC")
	}

	h[0].ManageSubBulb(1, true)
	if !h[0].SubCorridors[0].Bulb.isOn {
		t.Error("Failed to turn on Sub bulb")
	}
}

func TestOptimisePowerUsage(t *testing.T) {
	conf := &config.Config{2, 1, 2}
	h := Initialize(conf)

	h[0].ManageSubBulb(1, true)
	h[0].ManageSubBulb(2, true)
	p1 := h[0].PowerConsumption()
	if p1 != 45 {
		t.Errorf("Wrong power consumption before optimization. Expected: %d, Got %d", 45, p1)
	}

	h[0].OptimisePowerUsage()
	p2 := h[0].PowerConsumption()
	if p2 != 25 {
		t.Errorf("Wrong power consumption after optimization. Expected: %d, Got %d", 25, p2)
	}

	if h[0].SubCorridors[0].AC.isOn {
		t.Errorf("Sub cor AC should be off after optmization")
	}
	if h[0].SubCorridors[1].AC.isOn {
		t.Errorf("Sub cor AC should be off after optmization")
	}
	if !h[1].SubCorridors[1].AC.isOn {
		t.Errorf("Sub cor AC in other floor should not be affected after optmization")
	}
}
