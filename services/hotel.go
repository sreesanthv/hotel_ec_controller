package services

import (
	"log"
	"time"

	"github.com/sreesanthv/hotel_ec_controller/config"
	"github.com/sreesanthv/hotel_ec_controller/hotel"
	"github.com/sreesanthv/hotel_ec_controller/motion"
)

type HotelService struct {
	config             *config.Config
	Hotel              hotel.Hotel
	Watcher            *motion.MotionWatcher
	motionWaitDuration time.Duration
}

func NewHotelService(conf *config.Config, hot hotel.Hotel, w *motion.MotionWatcher, waitTime time.Duration) *HotelService {
	return &HotelService{conf, hot, w, waitTime}
}

func (s *HotelService) Start() {
	log.Println("Controller started with the default state")

	s.Hotel.PrintStatus()

	// watch for motion
	for {
		select {
		case <-s.Watcher.MotionChan:
			flrId := s.Watcher.Motion.Floor
			subId := s.Watcher.Motion.Sub
			if flrId > s.config.NoOfFloors || subId > s.config.SubCorridorsPerFloor {
				log.Println("Invalid motion input. Floor: ", flrId, " Sub Corridor:", subId)
				continue
			}

			if subId > 0 {
				// motion in SubCorridor -> turn on	 sub-cor light
				floor := s.Hotel.Floor(flrId)
				floor.ManageSubBulb(subId, true)
				log.Println("Motion detected in Sub Corridor")
				log.Println("\tFloor:", flrId, ", Sub:", subId)
				floor.OptimisePowerUsage()
				s.Hotel.PrintStatus()

				//schedule -> light turn off
				motionTime := time.Now().Unix()
				time.AfterFunc(s.motionWaitDuration, func() {
					if motionTime == s.Hotel.GetSubLastMotion(flrId, subId) {
						floor.ManageSubBulb(subId, false)
						log.Println("No motion in Sub Corridor")
						log.Println("Floor:", flrId, ", Sub:", subId)
						floor.OptimisePowerUsage()
						s.Hotel.PrintStatus()
					}
				})
				s.Hotel.SetSubLastMotion(flrId, subId, motionTime)
			}
		}
	}
}
