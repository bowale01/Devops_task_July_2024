package main

import (
	"bytes"
	"context"
	"encoding/json"
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
	cancel()
}

func runEntrance(ctx context.Context, name string, minWaitTime, maxWaitTime int) {
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-ctx.Done():
			return
		default:
			waitTime := time.Duration(rand.Intn(maxWaitTime-minWaitTime)+minWaitTime) * time.Millisecond
			time.Sleep(waitTime)
			d, _ := json.Marshal(struct {
				Entrance string `json:"entrance"`
			}{
				Entrance: name,
			})
			resp, err := http.Post("http://localhost:3000", "application/json", bytes.NewBuffer(d))
			if err != nil {
				log.Printf("Error sending request: %v", err)
				continue
			}
			resp.Body.Close()
		}
	}
}
