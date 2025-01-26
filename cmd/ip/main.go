package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const DEFAULT_PORT = 80

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", getIPHandler)

	port, err := portFromEnvOrDefault()
	if err != nil {
		log.Fatalf("Invalid port: %s", err.Error())
	}

	hostport := fmt.Sprintf("0.0.0.0:%d", port)

	log.Printf("Server started on %s ...", hostport)

	log.Fatal(http.ListenAndServe(hostport, mux))
}

func portFromEnvOrDefault() (int, error) {
	portEnv := os.Getenv("SERVICE_PORT")

	if portEnv == "" {
		return DEFAULT_PORT, nil
	}

	return strconv.Atoi(portEnv)
}

func getIPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	clientIP := readUserIP(r)

	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write([]byte(clientIP))

	if writeErr != nil {
		log.Printf("Error on answer")
	}
}

func readUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = strings.Split(r.RemoteAddr, ":")[0]
    }
    return IPAddress
}
