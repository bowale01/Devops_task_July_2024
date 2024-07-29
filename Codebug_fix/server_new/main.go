package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

type rider struct {
	id        int64
	name      string
	entrance  string
	vipStatus bool
}

type rollercoaster struct {
	nextId      int64
	idMu        sync.Mutex
	rideQueue   []*rider
	rideQueueMu sync.Mutex
	ride        []*rider
	rideMu      sync.Mutex
}

const (
	rideDuration = 5000
	numberOfCars = 8
	carCapacity  = 2
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rc := rollercoaster{
		ride: make([]*rider, 0, numberOfCars*carCapacity),
	}
	go rc.start(ctx)
	err := http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := &struct {
			Entrance string `json:"entrance"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(d); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		isVIP := rand.Float64() >= 0.7

		rc.idMu.Lock()
		r := &rider{id: rc.nextId, entrance: d.Entrance, name: namesgenerator.GetRandomName(0), vipStatus: isVIP}
		rc.nextId++
		rc.idMu.Unlock()

		rc.rideQueueMu.Lock()
		if r.vipStatus {
			rc.rideQueue = append([]*rider{r}, rc.rideQueue...)
		} else {
			rc.rideQueue = append(rc.rideQueue, r)
		}
		queueSize := len(rc.rideQueue)
		rc.rideQueueMu.Unlock()

		log.Printf("Entrance %s: %s entered the queue. Size: %d\n", d.Entrance, r.name, queueSize)
	}))
	if err != nil {
		panic(err)
	}
}

func (rc *rollercoaster) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			rc.rideQueueMu.Lock()
			for i := 0; i < numberOfCars*carCapacity && i < len(rc.rideQueue); i++ {
				rc.rideQueueMu.Unlock()

				rc.rideMu.Lock()
				r := rc.rideQueue[i]
				rc.ride = append(rc.ride, r)
				car := i / carCapacity
				carSeat := i % carCapacity
				rc.seatRider(r, car, carSeat)
				rc.rideMu.Unlock()

				rc.rideQueueMu.Lock()
			}
			rc.rideQueue = rc.rideQueue[min(numberOfCars*carCapacity, len(rc.rideQueue)):]
			rc.rideQueueMu.Unlock()

			log.Println("Ride: started")
			time.Sleep(rideDuration * time.Millisecond)
			log.Println("Ride: finished")

			rc.rideMu.Lock()
			rc.ride = rc.ride[:0] // Clear the ride
			rc.rideMu.Unlock()
		}
	}
}

func (rc *rollercoaster) seatRider(r *rider, car, carSeat int) {
	log.Printf("Ride: %s entering the car %d in seat %d\n", r.name, car+1, carSeat+1)
	time.Sleep(400 * time.Millisecond)
	log.Printf("Ride: %s entered the car %d in seat %d\n", r.name, car+1, carSeat+1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
