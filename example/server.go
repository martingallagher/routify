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

	// Local
	"./website"

	// External
	"github.com/martingallagher/routify/router"
)

func main() {
	http.HandleFunc("/", globalHandler)

	addr := "0.0.0.0:1337"

	// Fire up goroutine so we can capture signals
	go func() {
		log.Fatal(http.ListenAndServe(addr, nil))
	}()

	log.Printf("server started: http://%s", addr)

	// Capture signals e.g. CTRL+C
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Wait for signal
	<-sig
}

func globalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain; charset=utf-8")

	h, p, err := website.Routes.Get(r)

	if err != nil {
		if e, ok := err.(*router.Error); ok {
			w.WriteHeader(e.StatusCode())
			w.Write([]byte(e.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Wups, internal error!"))
		}

		return
	}

	h(w, r, p)
}
