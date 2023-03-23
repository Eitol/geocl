package geocl

import (
	_ "embed"
	"errors"
)

var errLocCodeNotFound = errors.New("location code not found")
var errRegCodeNotFound = errors.New("region code not found")
var errComCodeNotFound = errors.New("commune code not found")

// GetByLocCode returns a Location by its LOCCode.
// If LOCCode is not found, it returns an error.
// LOCCode is a unique identifier.
// There is only one Location per LOCCode.
// LOCCode is a 3-character string. e.g. "ACZ"
func GetByLocCode(locCode string) (Location, error) {
	assertInit()
	loc, err := byLocCode[locCode]
	if !err {
		return Location{}, errors.Join(errLocCodeNotFound, errors.New("invalid locCode: "+locCode))
	}
	return loc, nil
}

// GetByRegCode returns a slice of Locations by its RegCode.
// If RegCode is not found, it returns an error.
// RegCode is not a unique identifier.
// There are many Locations per RegCode.
// RegCode is a 2-character string. e.g. "15"
func GetByRegCode(regCode string) ([]Location, error) {
	assertInit()
	loc, ok := byRegCode[regCode]
	if !ok {
		return nil, errors.Join(errRegCodeNotFound, errors.New("invalid regCode: "+regCode))
	}
	return loc, nil
}

// GetByComCode returns a slice of Locations by its COMCode.
// If COMCode is not found, it returns an error.
// COMCode is not a unique identifier.
// There are many Locations per COMCode.
// COMCode is a 5-character string. e.g. "15101"
func GetByComCode(comCode string) ([]Location, error) {
	assertInit()
	loc, ok := byComCode[comCode]
	if !ok {
		return nil, errors.Join(errComCodeNotFound, errors.New("invalid comCode: "+comCode))
	}
	return loc, nil
}

// GetAll returns a slice of all Locations.
func GetAll() []Location {
	assertInit()
	return list
}
