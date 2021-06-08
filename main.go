package main

import (
	"time"

	"github.com/sreesanthv/hotel_ec_controller/config"
	"github.com/sreesanthv/hotel_ec_controller/hotel"
	"github.com/sreesanthv/hotel_ec_controller/motion"
	"github.com/sreesanthv/hotel_ec_controller/services"
)

func main() {
	motionChan := make(chan bool)
	conf := config.Get()

	watcher := motion.NewWatcher(motionChan)
	h := hotel.Initialize(conf)

	service := services.NewHotelService(conf, h, watcher, 1*time.Minute)
	service.Start()
}
