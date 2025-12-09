package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// Response in JSON structure
type Response struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}


func getClientIP(r *http.Request) string {
	
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For header check
		return strings.Split(forwarded, ",")[0]
	}
	
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	
	resp := Response{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		IP:        getClientIP(r),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)
	
	log.Printf("SimpleTimeService starting on port %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}