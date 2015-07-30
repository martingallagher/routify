// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"./website"
)

func main() {
	addr := "0.0.0.0:1337"
	server := &http.Server{
		Addr:         addr,
		Handler:      website.Routes,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Fire up goroutine so we can capture signals
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	log.Printf("server started: http://%s", addr)

	// Capture signals e.g. CTRL+C
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Wait for signal
	<-sig
}
