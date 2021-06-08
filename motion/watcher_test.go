package motion

import (
	"os"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	now := time.Now()
	ipFile, _ := os.OpenFile(MOTION_INPUT_FILE, os.O_CREATE|os.O_RDWR, 0644)
	ipFile.WriteString(`{"floor":1,"sub":2}`)
	ipFile.Close()
	defer os.Remove(MOTION_INPUT_FILE)

	var motionSignal chan bool
	w := &MotionWatcher{MotionChan: motionSignal}
	w.read()

	if w.Motion.Floor != 1 || w.Motion.Sub != 2 {
		t.Error("Failed to parse the Motion input file")
	}

	if getInputLastModified().Round(1*time.Second).Equal(now.Round(1*time.Second)) != true {
		t.Error("Failed to set last modifed time of input file", getInputLastModified().Sub(now), now)
	}
}
