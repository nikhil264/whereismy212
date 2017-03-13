package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nikhil264/whereismy212/location"
)

var err error

func onbusHandler(w http.ResponseWriter, r *http.Request) {
	l := location.Loc{}
	// err := json.NewDecoder(r.Body).Decode(&l)
	// if err != nil {
	// 	http.Error(w, err.Error(), 400)
	// 	return
	// }
	if r.Method == "POST" && (r.PostFormValue("dest") == "bs" || r.PostFormValue("dest") == "sb") {
		err = r.ParseForm()
		tmp := r.PostFormValue("lat") + " " + r.PostFormValue("lng") + " " + r.PostFormValue("time") + " " + r.PostFormValue("dest")
		m = append(m, tmp)
		l.Lat, _ = strconv.ParseFloat(r.PostFormValue("lat"), 0)
		l.Lng, _ = strconv.ParseFloat(r.PostFormValue("lng"), 0)
		l.Time, _ = strconv.ParseUint(r.PostFormValue("time"), 10, 0)
		l.Dest = r.PostFormValue("dest")
		fmt.Fprintln(w, "OK")
	}
	location.UpdateLoc(l)
}

var m []string

func trackHandler(w http.ResponseWriter, r *http.Request) {
	l := location.GetLoc()
	c := json.NewEncoder(w)
	//bus locations are encoded as json
	for k, v := range l {
		v.Dest = k
		fmt.Fprintln(w, k)
		c.Encode(v)
	}
	fmt.Fprintln(w, len(m))
	for _, v := range m {
		fmt.Fprintln(w, v)
	}
}
func main() {
	location.InitBusLoc()
	http.HandleFunc("/track", trackHandler)
	http.HandleFunc("/onbus", onbusHandler)
	http.ListenAndServe(":9090", nil)
}
