## Chile Location Codes - Go Library

This repository contains a Go library that provides an easy way to obtain location codes in Chile. 

### What are location codes?
```json5
[
  {
    "locCode": "ACZ",
    "locName": "CODPA",
    "regCode": "15",
    "regName": "Arica y Parinacota",
    "comCode": "15101",
    "comName": "ARICA"
  },
  {
    "locCode": "AGN",
    "locName": "GUANACAHUA",
    "regCode": "15",
    "regName": "Arica y Parinacota",
    "comCode": "15101",
    "comName": "ARICA"
  },
  //...
]
```
The library exposes functions to search for locations by their LOCCode, RegCode, or COMCode, as well as to retrieve all available locations.

Features
- Get location information by LOCCode (unique identifier)
- Get a list of locations by RegCode (region code)
- Get a list of locations by COMCode (commune code)
- Get a list of all available locations

### Struct: Location


type Location struct {
    LOCCode string `json:"locCode"` // Unique identifier for the location (e.g. "ACZ")
    LOCName string `json:"locName"` // Name of the location (e.g. "CODPA")
    RegCode string `json:"regCode"` // Code of the region (e.g. "15")
    RegName string `json:"regName"` // Name of the region (e.g. "Arica y Parinacota")
    COMCode string `json:"comCode"` // Code of the commune (e.g. "15101")
    COMName string `json:"comName"` // Name of the commune (e.g. "ARICA")
}

### Functions
```go
func GetByLocCode(locCode string) (Location, error)
func GetByRegCode(regCode string) ([]Location, error)
func GetByComCode(comCode string) ([]Location, error)
func GetAll() []Location
```

### Installation

```go
go get github.com/Eitol/geocl
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/Eitol/geocl/v1"
)

func main() {
	loc, err := geocl.GetByLocCode("ACZ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Location: %v\n", loc)
}
// Output: Location: {ACZ CODPA 15 Arica y Parinacota 15101 ARICA}
```

### License

Use as you wish
