package main

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/nikhil264/location"
)

func onbusHandler(w http.ResponseWriter, r *http.Request) {
	l := location.Loc{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	go location.UpdateLoc(l)
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	l := location.GetLoc()
	c := json.NewEncoder(w)

	for _, v := range l {
		c.Encode(v)

	}
}

func main() {
	http.HandleFunc("/track", trackHandler)
	http.HandleFunc("/onbus", onbusHandler)
	http.ListenAndServe(":9090", nil)
	fmt.Println("dd")
}
