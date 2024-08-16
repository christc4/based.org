# The Link Register (and return addresses)

Up: [ARM](arm),[Registers](registers)
Brother(s):
TARGET DECK

This register is used whenever BL or BLX is called, and the return address (# of instruction line to return to) is stored in the Link Register (LR), and is copied from the PC.

BL - branch and link

BX - branch and exchange, allows to switch between Thumb/ARM code

the return address can be stored on a stack; pushed onto it at a subroutine call and popped off to overwrite the contents of the PC.


































#### Why:
#### How:









