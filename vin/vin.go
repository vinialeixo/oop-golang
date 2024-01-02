package vin

import "fmt"

// func Manufacturer(vin string) string {
// 	manufacturer := vin[:3]
// 	// if the last digit of the manufacturer ID is a 9
// 	// the digits 12 to 14 are the second part of the ID

// 	if manufacturer[2] == '9' {
// 		manufacturer += vin[11:14]
// 	}
// 	return manufacturer
// }

/*
So this function works correctly when given the right input, but it has some problems:

There is no guarantee that the input string is a VIN.
For strings shorter than three characters, the function causes a panic.
The optional second part of the ID is a feature of European VINs only.
	The function would return wrong IDs for US cars having a 9 as the third digit of the manufacturer code.
*/

//Go OOP: Binding Functions to a Type

// first refactoring is to make VINs their own type and bind the Manufacturer() function to it
type VIN string

//Using Constructors

/*
A more elegant way is to put the validity checks in a constructor for the VIN type,
so that the Manufacturer() function is called for valid VINs only and does not need checks and error handling
*/
func NewVIN(code string) (VIN, error) {
	if len(code) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
	}
	return VIN(code), nil
}
func (v VIN) Manufacturer() string {
	manufacturer := v[:3]
	if manufacturer[2] == '9' {
		manufacturer += v[11:14]
	}

	return string(manufacturer)
}
