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

func TestVIN_Manufacturer(t *testing.T) {

	manufacturer := validVIN.Manufacturer()
	if manufacturer != "W0L" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, validVIN)
	}

	invalidVIN.Manufacturer()
}
