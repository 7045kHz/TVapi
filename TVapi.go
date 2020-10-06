package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tarm/serial"
	"net/http"
)

func main() {
	FIRE2ALL := "MT00SW0100NT"
	BLUERAY2ALL := "MT00SW0200NT"
	//STATUS := "MT00RD0000NT"

	r := mux.NewRouter()
	r.HandleFunc("/TV/{options}", func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	options := vars["options"]
	switch options {
		case "dvd", "blueray":
			c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
			s, err := serial.OpenPort(c)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Error opening serial port %v", err)
			}
			// Send Command to sent BlueRay to all Outputs
			_, err = s.Write([]byte(BLUERAY2ALL))
			if err != nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Error Setting Matric to BlueRay|DVD to all %v", err)
			}
			s.Close()
		case "fire":
			c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
			s, err := serial.OpenPort(c)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Error opening serial port %v", err)
			}
			// Send Command to send Fire Cube to all Outputs
			_, err = s.Write([]byte(FIRE2ALL))
			if err != nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Error Setting Matric to FIRE Cube to all %v", err)
			}
			s.Close()
		}
	})
	http.ListenAndServe("192.168.15.250:8000", r)
}
