package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
)

var rtp float64

const addr = ":64333"

func main() {
	flag.Float64Var(&rtp, "rtp", 1.0, "target RTP value in [0.0, 1.0]")
	flag.Parse()

	if rtp <= 0.0 || rtp > 1.0 {
		log.Fatalf("Invalid RTP value: must be in [0.0, 1.0], got %f", rtp)
	}

	http.HandleFunc("/get", handleGet)

	log.Printf("Starting server on http://localhost%s with RTP=%.4f", addr, rtp)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	multiplier := generateMultiplier(rtp)

	resp := map[string]float64{"result": multiplier}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func generateMultiplier(rtp float64) float64 {
	x := 1.0 + rand.Float64()*(10000.0-1.0)

	if rand.Float64() < rtp {
		return x + rand.Float64()*(10000.0-x)
	} else {
		return 1.0 + rand.Float64()*(x-1.0)
	}
}
