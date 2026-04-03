package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"://github.com"
	"://github.com"
	"://github.com/promhttp"
)

var (
	streamingActive = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "obs_streaming_active",
		Help: "1 if OBS is streaming, 0 otherwise",
	})
	obsFps = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "obs_fps",
		Help: "Current FPS of OBS Studio",
	})
)

func main() {
	prometheus.MustRegister(streamingActive, obsFps)

	obsHost := os.Getenv("OBS_HOST")
	obsPassword := os.Getenv("OBS_PASSWORD")

	if obsHost == "" {
		obsHost = "172.16.10.197:4455"
	}

	client, err := goobs.New(obsHost, goobs.WithPassword(obsPassword))
	if err != nil {
		log.Fatalf("Failed to connect to OBS: %v", err)
	}

	go func() {
		for {
			status, _ := client.Stream.GetStreamStatus()
			if status != nil && status.OutputActive {
				streamingActive.Set(1)
			} else {
				streamingActive.Set(0)
			}

			stats, _ := client.General.GetStats()
			if stats != nil {
				obsFps.Set(stats.Fps)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9407", nil))
}
