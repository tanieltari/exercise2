package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"
)

type Vehicle struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
}

var stops = []string{
	"Virmalise",
	"Luite",
	"Renniotsa",
	"Balti jaam",
	"Marja",
	"Tihase",
	"Kelmiküla",
	"Vabaduse väljak",
	"Ahtri",
	"Pronksi",
	"Ülase",
	"Koskla",
}

func main() {
	// Exercise number 1 🛑
	http.HandleFunc("GET /stops", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		query = strings.ToLower(strings.TrimSpace(query))
		var matches []string
		for _, stop := range stops {
			if strings.Contains(strings.ToLower(stop), query) {
				matches = append(matches, stop)
			}
		}
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(matches)
	})

	// Exercise number 2 🚀
	http.HandleFunc("GET /realtime", func(w http.ResponseWriter, r *http.Request) {
		var vehicles []Vehicle
		for i := 0; i < 5; i++ {
			vehicles = append(vehicles, Vehicle{
				Latitude:  randomFloat(1, 10),
				Longitude: randomFloat(1, 10),
				Type:      "bus",
				Name:      fmt.Sprintf("Bus %d", i+1),
			})
		}
		for i := 0; i < 2; i++ {
			vehicles = append(vehicles, Vehicle{
				Latitude:  randomFloat(1, 10),
				Longitude: randomFloat(1, 10),
				Type:      "train",
				Name:      fmt.Sprintf("Train %d", i+1),
			})
		}
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vehicles)
	})

	fmt.Println("Server started on http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
