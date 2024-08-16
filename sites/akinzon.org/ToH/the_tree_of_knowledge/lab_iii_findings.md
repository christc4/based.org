# Lab III Findings

Up: [Computer Systems Labs](computer_systems_labs)
Brother(s):
TARGET DECK

255 + 2 = 1 (wrap-around)

4 -  5 = 255 (wrap-around)

`muls` command allows for wrap-around (s stands for signed)

```
CMP r1, #0
beq .L5
```

branches to tag `L5` if contents of register 1 = 0

n-factorial in assembly:
n = r1
counter = r2

```
CMP r1, #0
BEQ .specialcase
CMP r1, #1
BEQ .specialcase
B mainloop

mainloop:
	CMP r2, #1
	BEQ .end
	MULS r1, r1, r2
	SUBS r2, r2, #1
	B mainloop
```

For our lab, 12! was the highest possible factorial representable.
































#### Why:
#### How:









