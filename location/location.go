package location

import "sync"
import "strconv"

//Loc represents the location of the bus
type Loc struct {
	Lng  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
	Time uint64  `json:"time"`
	Dest string  `json:"dest"`
}

var busLoc map[string][]Loc

var currBusses []string

var mutex = &sync.Mutex{}

//GetLoc return last updated bus location
func GetLoc() map[string]Loc {
	currLoc := make(map[string]Loc)
	//test locations
	// currLoc["bs0655"] = Loc{12.1, 12.3, 21321, "tapri"}
	// currLoc["bs0625"] = Loc{12.1, 12.3, 21321, "tapri"}
	// currLoc["bs0675"] = Loc{12.1, 12.3, 21321, "tapri"}
	for _, v := range currBusses {
		currLoc[v] = busLoc[v][len(busLoc[v])-1]
	}
	return currLoc
}

//UpdateLoc handles the new bus location received
func UpdateLoc(l Loc) {

	if validLoc(l) != true {
		return
	}

	whichBus(l)

}

//whichBus finds which current running bus location is received
func whichBus(l Loc) {
	t := strconv.FormatUint(l.Time, 10)
	t = t[len(t)-4:]
	if l.Dest == "bs" {
		//to check if the location is from any running bits to sec busses
		for _, v := range currBusses {
			if l.Dest+t > v {
				busLoc[v] = append(busLoc[v], l)
				return
			}
		}
		// to start a new bus from bits to sec
		for k := range busLoc {
			if l.Dest+t > k {
				mutex.Lock()
				currBusses = append(currBusses, k)
				busLoc[k] = append(busLoc[k], l)
				mutex.Unlock()
				return
			}
		}
	}
	if l.Dest == "sb" {
		//to check if the location is from any running sec to bits busses
		for _, v := range currBusses {
			if v[0:2] == "sb" && l.Dest+t > v {
				busLoc[v] = append(busLoc[v], l)
				return
			}
		}
		//to start a new bus from sec to bits
		for k := range busLoc {
			if k[0:2] == "sb" && l.Dest+t > k {
				mutex.Lock()
				currBusses = append(currBusses, k)
				busLoc[k] = append(busLoc[k], l)
				mutex.Unlock()
				return
			}
		}
	}
	//if user directly specifies which bus he is in.
	if _, ok := busLoc[l.Dest]; ok {
		busLoc[l.Dest] = append(busLoc[l.Dest], l)
		return
	}

}

func validLoc(l Loc) bool {
	return true
}

//gives the slice of current running busses from either sides
// func GetCurrBusses() []string {
// 	for k, v := range busLoc {
// 		if len(v) > 0 {
// 			currBusses = append(currBusses, k)
// 		}
// 	}
// 	return currBusses
// }

func initBusLoc() {
	busLoc = make(map[string][]Loc)

	busLoc["bs0655"] = []Loc{}
	busLoc["bs0755"] = []Loc{}
	busLoc["bs0850"] = []Loc{}
	busLoc["bs0920"] = []Loc{}
	busLoc["bs0950"] = []Loc{}
	busLoc["bs1100"] = []Loc{}
	busLoc["bs1200"] = []Loc{}
	busLoc["bs1300"] = []Loc{}
	busLoc["bs1400"] = []Loc{}
	busLoc["bs1505"] = []Loc{}
	busLoc["bs1605"] = []Loc{}
	busLoc["bs1720"] = []Loc{}
	busLoc["bs1815"] = []Loc{}
	busLoc["bs1915"] = []Loc{}
	busLoc["bs2015"] = []Loc{}
	busLoc["sb0750"] = []Loc{}
	busLoc["sb0820"] = []Loc{}
	busLoc["sb0850"] = []Loc{}
	busLoc["sb0950"] = []Loc{}
	busLoc["sb1100"] = []Loc{}
	busLoc["sb1200"] = []Loc{}
	busLoc["sb1300"] = []Loc{}
	busLoc["sb1410"] = []Loc{}
	busLoc["sb1505"] = []Loc{}
	busLoc["sb1605"] = []Loc{}
	busLoc["sb1720"] = []Loc{}
	busLoc["sb1815"] = []Loc{}
	busLoc["sb1915"] = []Loc{}
	busLoc["sb2015"] = []Loc{}
	busLoc["sb2115"] = []Loc{}

}
