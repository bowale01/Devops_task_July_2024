package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go runEntrance(ctx, "North", 50, 3000)
	go runEntrance(ctx, "East", 100, 15000)
	go runEntrance(ctx, "South", 50, 3000)
	go runEntrance(ctx, "West", 400, 6000)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Shutdown signal received, exiting...")
}

func runEntrance(ctx context.Context, name string, minWaitTime, maxWaitTime int) {
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-ctx.Done():
			return
		default:
			sleepDuration := time.Duration(rand.Intn(maxWaitTime-minWaitTime+1) + minWaitTime) * time.Millisecond
			time.Sleep(sleepDuration)
			d, err := json.Marshal(struct {
				Entrance string `json:"entrance"`
			}{
				Entrance: name,
			})
			if err != nil {
				log.Printf("Error marshalling JSON: %v\n", err)
				continue
			}
			resp, err := http.Post("http://localhost:3000", "application/json", bytes.NewBuffer(d))
			if err != nil {
				log.Printf("Error making POST request: %v\n", err)
				continue
			}
			resp.Body.Close()
		}
	}
}
