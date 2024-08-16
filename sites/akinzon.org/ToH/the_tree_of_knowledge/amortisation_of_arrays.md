# Amortisation of Arrays

Up: [Variable-Length Arrays](variable-length_arrays)
Brother(s):
TARGET DECK

Say, we had an array that was of initial size 2, and wanted to add 16 elements, we would have to double the array size each time the array fills up. The amount of operations needed is:
	adds: 16
	copies: 2 + 4 + 8 + 16
	`new Array[]`: 4

This is 46 copy + add operations + 4 `new` operations!

[General Time Analysis for Array Amortisation](general_time_analysis_for_array_amortisation)



































#### Why:
#### How:









