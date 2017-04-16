package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"bytes"

	"github.com/nikhil264/whereismy212/location"
)

func main() {
	t, _ := strconv.ParseUint(time.Now().Format("20060102150405"), 10, 0)
	l := location.Loc{Lng: 17.555556, Lat: 78.563919, Time: t, Dest: "bs"}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(l)

	res, _ := http.Post("http://localhost:9090/onbus", "application/json", b)
	io.Copy(os.Stdout, res.Body)
	res, _ = http.Get("http://localhost:9090/track")
	io.Copy(os.Stdout, res.Body)

	// s := "20060102151405"
	// s = "ab" + s[len(s)-4:]
	// if s > "bs1406" {
	// 	fmt.Println("gg")
	// }
	// fmt.Println(s)
	// m := make([]string, 1)
	// m =
	// m["gg"] = []string{}
	// tmp := m["gg"]
	// if len(m["gg"]) == 0 {
	// 	fmt.Println("grfg")
	// }
	// fmt.Println(m)
	// // s := []string{"gg", "wp", "we"}
	// m["gg"] = append(m["gg"], "ff")
	// fmt.Println(m)
}
