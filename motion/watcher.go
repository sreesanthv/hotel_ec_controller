package motion

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

const MOTION_INPUT_FILE = "motion_ip.json"

type Motion struct {
	Floor int `json:"floor"`
	Main  int `json:"main"`
	Sub   int `json:"sub"`
}

type MotionWatcher struct {
	LastUpdatedAt time.Time
	Motion        *Motion
	MotionChan    chan bool
	m             sync.Mutex
}

func NewWatcher(motionSignal chan bool) *MotionWatcher {
	w := &MotionWatcher{MotionChan: motionSignal}
	w.read()
	go w.watch()
	return w
}

// check if motion iput file is changed
func (w *MotionWatcher) watch() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			if !w.LastUpdatedAt.Equal(getInputLastModified()) {
				w.read()
				w.MotionChan <- true
			}
		}
	}
}

// parse the motion input
// send the motion through channel
func (w *MotionWatcher) read() {
	w.m.Lock()
	defer w.m.Unlock()

	file, err := ioutil.ReadFile(MOTION_INPUT_FILE)
	if err != nil {
		log.Fatal("Error reading Motion input file", err)
	}

	m := new(Motion)
	err = json.Unmarshal(file, m)
	if err != nil {
		log.Fatal("Invalid Motion input file", err)

	}
	w.Motion = m
	w.LastUpdatedAt = getInputLastModified()
}

// get last modified time of motion input file
func getInputLastModified() time.Time {
	info, err := os.Stat(MOTION_INPUT_FILE)
	if err != nil {
		log.Fatal("Error reading Motion input file info", err)
	}
	return info.ModTime()
}
