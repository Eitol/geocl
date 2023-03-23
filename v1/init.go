package geocl

import (
	_ "embed"
	"encoding/json"
	"sync"
)

/*
Content of geo.json:
[
   {
      "locCode": "ACZ",
      "locName": "CODPA",
      "regCode": "15",
      "regName": "Arica y Parinacota",
      "comCode": "15101",
      "comName": "ARICA"
   },
   ...
]
*/

//go:embed geo.json
var jsonGeo string
var isInitOnce sync.Once
var list []Location
var byLocCode map[string]Location
var byRegCode map[string][]Location
var byComCode map[string][]Location

func initStore() {
	byLocCode = make(map[string]Location)
	byRegCode = make(map[string][]Location)
	byComCode = make(map[string][]Location)
	err := json.Unmarshal([]byte(jsonGeo), &list)
	if err != nil {
		panic(err)
	}
	for _, loc := range list {
		byLocCode[loc.LOCCode] = loc
		byRegCode[loc.RegCode] = append(byRegCode[loc.RegCode], loc)
		byComCode[loc.COMCode] = append(byComCode[loc.COMCode], loc)
	}
}

func assertInit() {
	isInitOnce.Do(initStore)
}
