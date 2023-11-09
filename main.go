package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const PORT string = "0.0.0.0:8080"

var LEAD_TIME string = os.Getenv("LEAD_TIME")

func main() {
	startTime := time.Now()
	leadTime, err := strconv.Atoi(LEAD_TIME)

	log.Print(leadTime)

	if err != nil {
		log.Fatal("Could not read or convert lead time")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		_, _ = writer.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/health", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		for {
			elapsed := time.Since(startTime).Seconds()
			log.Printf("elapsed: %f", elapsed)
			if elapsed > float64(leadTime) && rand.Intn(100) == 1 {
				os.Exit(1)
			} else {
				time.Sleep(time.Second * 5)
			}
		}

	}()

	log.Print("Staring server")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
