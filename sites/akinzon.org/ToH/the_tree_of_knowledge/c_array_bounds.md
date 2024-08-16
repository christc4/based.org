# C Array Bounds

Up: [C Arrays](c_arrays)
Brother(s):
TARGET DECK

C does not check array bounds (so no `IndexOutOfBoundsException`!), meaning that whatever is in that out of bounds index will be used as if it was in the array.

This is because arrays are stored as a [C Pointers|pointer](c_pointers|pointer) to the first element in the array, with all other elements being stored in an offset of `array[0]`. And since C does not check bounds, if there is something 13 blocks away from `array[0]`, C is able to use it as if there wasn't an INSANE problem with this!!!


































#### Why:
#### How:









