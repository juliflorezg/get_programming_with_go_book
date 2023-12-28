package main

import (
	"log"
	"time"
)

// earthReceiver receives messages sent from Mars.
// As connectivity is limited, it only receives messages for some time every Mars day.
func earthReceiver(msgc chan []Message) {
	for {
		time.Sleep(dayLength - receiveTimePerDay)
		receiveMarsMessages(msgc)
	}
}

// receiveMarsMessages receives messages sent from Mars for the given duration.
func receiveMarsMessages(msgc chan []Message) {
	finished := time.After(receiveTimePerDay)
	for {
		select {
		case <-finished:
			return
		case ms := <-msgc:
			for _, m := range ms {
				log.Printf("earth received report of life sign level %d from %s at %v", m.LifeSigns, m.Rover, m.Pos)
			}
		}
	}
}
