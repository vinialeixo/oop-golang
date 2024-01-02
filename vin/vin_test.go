package vin_test

import (
	"testing"

	vin "github.com/vinialeixo/oop-golang/vin"
)

// const testVIN = "W09000051T2123456"

// func TestVIN_Manufacturer(t *testing.T) {
// 	manufacturer := vin.Manufacturer(testVIN)
// 	if manufacturer != "W09123" {
// 		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
// 	}
// }

const (
	validVIN   = vin.VIN("W0L000051T2123456")
	invalidVIN = vin.VIN("W0")
)

func TestVIN_NEW(t *testing.T) {
	_, err := vin.NewVIN(string(invalidVIN))
	if err != nil {
		t.Errorf("creating valid VIN returned an error: %s", err.Error())
	}
	_, err = vin.NewVIN(string(invalidVIN))
	if err == nil {
		t.Error("creating invalid VIN did not return an error")
	}
}

func TestVIN_Manufacturer(t *testing.T) {
	testVIN, _ := vin.NewVIN(string(validVIN))
	manufacturer := validVIN.Manufacturer()
	if manufacturer != "W0L" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}

/*
The test for the Manufacturer() function can now omit testing an invalid
VIN since it already would have been rejected by the NewVIN constructor.
*/
// func TestVIN_Manufacturer(t *testing.T) {

// 	manufacturer := validVIN.Manufacturer()
// 	if manufacturer != "W0L" {
// 		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, validVIN)
// 	}

// 	invalidVIN.Manufacturer() // panic!
// 	//validVIN.Manufacturer() //correct
// }
