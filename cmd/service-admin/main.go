package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func start() {
	cmd := exec.Command("./start.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func stop() {
	cmd := exec.Command("./stop.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func info() string {
	cmd := exec.Command("./info.sh")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func do() {
	out := info()
	if strings.Contains(out, "not") {
		// stuff is not running
		fmt.Println("stuff is not running, starting up.")
		start()
	} else {
		// stuff is running
		fmt.Println("stuff is running, shutting down.")
		stop()
	}
	fmt.Println(info())
}

var authStr = "spc1962"

func main() {

	http.Handle("/start", handleStart())
	http.Handle("/stop", handleStop())
	http.Handle("/info", handleInfo())
	log.Panic(http.ListenAndServe(":8080", nil))
}

func handleInfo() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			vals := r.URL.Query()
			val := vals.Get("auth")
			if val != authStr {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			out, text := info(), "server running: FALSE"
			if !strings.Contains(out, "not") {
				// server is running
				text = "server running: TRUE"
			}
			fmt.Fprintf(w, "%s", text)
			return
		}
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	return http.HandlerFunc(fn)
}

func handleStart() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			vals := r.URL.Query()
			val := vals.Get("auth")
			if val != authStr {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			out := info()
			if !strings.Contains(out, "not") {
				// server is running
				fmt.Fprintf(w, "server is already running: cannot start")
				return
			}
			start()
			fmt.Fprintf(w, "server start: executed")
			return
		}
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	return http.HandlerFunc(fn)
}

func handleStop() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			vals := r.URL.Query()
			val := vals.Get("auth")
			if val != authStr {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			out := info()
			if !strings.Contains(out, "not") {
				// server is running
				stop()
				fmt.Fprintf(w, "server stop: executed")
				return
			}
			fmt.Fprintf(w, "server is not running: cannot stop")
			return
		}
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	return http.HandlerFunc(fn)
}
