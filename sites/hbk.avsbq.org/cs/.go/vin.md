# Car vin

[source](https://www.toptal.com/go/golang-oop-tutorial)

	package vin

	func Manufacturer(vin string) string {

	manufacturer := vin[: 3]
	// if the last digit of the manufacturer ID is a 9
	// the digits 12 to 14 are the second part of the ID
	if manufacturer[2] == '9' {
	manufacturer += vin[11: 14]
	}

	return manufacturer
	} 

## Explanation

`manufacturer := vin[: 3]`

First three characters of vin are extractred and assigned to the `manufacturer` variable

`vin[:3]` means "take the substring of vin from the start (index 0) to index 3, but not including index 3." This gives us the first three characters.

`if manufacturer[2] == '9' {`

This function checks if the third character (at index 2) of the `manufacturer` string is `9`

In some VIN schemes, if the third character of the manufacturer identifier is '9', it indicates that the identifier is longer than three characters and requires additional digits

`manufacturer += vin[11: 14]`

If the condition is true (i.e., the third character is '9'), the function appends three more characters from the VIN, specifically characters at positions 12 to 14 (vin[11:14]).
vin[11:14] means "take the substring of vin starting at index 11 and ending at index 14, but not including index 14."


