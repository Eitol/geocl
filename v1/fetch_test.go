package geocl

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetByLocCode(t *testing.T) {
	type args struct {
		locCode string
	}
	tests := []struct {
		name string
		args args
		want Location
		err  error
	}{
		{
			name: "found",
			args: args{locCode: "ACZ"},
			want: Location{
				LOCCode: "ACZ",
				LOCName: "CODPA",
				RegCode: "15",
				RegName: "Arica y Parinacota",
				COMCode: "15101",
				COMName: "ARICA",
			},
			err: nil,
		},
		{
			name: "not found",
			args: args{locCode: "XXX"},
			want: Location{},
			err:  errLocCodeNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLoc, err := GetByLocCode(tt.args.locCode)

			if err != nil && !errors.Is(err, tt.err) {
				t.Errorf("GetByLocCode() error = %v, wantErr %v", err, tt.err)
			}
			if !reflect.DeepEqual(gotLoc, tt.want) {
				t.Errorf("GetByLocCode() = %v, want %v", gotLoc, tt.want)
			}
		})
	}
}

func TestGetByRegCodeFound(t *testing.T) {
	var regCode = "16"
	const numberOfLocationsInRegion = 65
	gotLoc, err := GetByRegCode(regCode)
	if err != nil {
		t.Errorf("GetByRegCode() error = %v", err)
	}
	if len(gotLoc) == 0 {
		t.Errorf("GetByRegCode() = %v, want > 0", gotLoc)
	}
	if len(gotLoc) != numberOfLocationsInRegion {
		t.Errorf("GetByRegCode() = %v, want 5", len(gotLoc))
	}
}

func TestGetByRegCodeNotFound(t *testing.T) {
	var regCode = "XX"
	gotLoc, err := GetByRegCode(regCode)
	if err != nil && !errors.Is(err, errRegCodeNotFound) {
		t.Errorf("GetByRegCode() error = %v, wantErr %v", err, errRegCodeNotFound)
	}
	if len(gotLoc) != 0 {
		t.Errorf("GetByRegCode() = %v, want 0", len(gotLoc))
	}
}

func TestGetByComCodeFound(t *testing.T) {
	var comCode = "15101"
	const numberOfLocationsInCommune = 15
	gotLoc, err := GetByComCode(comCode)
	if err != nil {
		t.Errorf("GetByComCode() error = %v", err)
	}
	if len(gotLoc) == 0 {
		t.Errorf("GetByComCode() = %v, want > 0", gotLoc)
	}
	if len(gotLoc) != numberOfLocationsInCommune {
		t.Errorf("GetByComCode() = %v, want 1", len(gotLoc))
	}
}

func TestGetByComCodeNotFound(t *testing.T) {
	var comCode = "XXXXX"
	gotLoc, err := GetByComCode(comCode)
	if err != nil && !errors.Is(err, errComCodeNotFound) {
		t.Errorf("GetByComCode() error = %v, wantErr %v", err, errComCodeNotFound)
	}
	if len(gotLoc) != 0 {
		t.Errorf("GetByComCode() = %v, want 0", len(gotLoc))
	}
}

func TestGetAll(t *testing.T) {
	const numberOfLocations = 1670
	gotLoc := GetAll()
	if len(gotLoc) != numberOfLocations {
		t.Errorf("GetAll() = %v, want %v", len(gotLoc), numberOfLocations)
	}
}
