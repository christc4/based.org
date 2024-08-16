# Flip-Flops

Up: [Computer Systems](computer_systems)
Brother(s):
TARGET DECK

Flip-flops are some of the most fundamental units of a computer.
	- Flip-flops are **bistable,** meaning that they can be in 1 of 2 states and remain in that state, even without input. As a result, they can be used as 1-bit memory.
	- Flip-flops have two outputs, which are complements of each other. These are labelled as $Q$ and $\bar Q$.
	
![Pasted image 20240109113744.png](pasted_image_20240109113744.png)

This a S-R flip-flop, where S is used to set the output and R is used to reset it. The gates above are NOR gates.

![Pasted image 20240109114016.png](pasted_image_20240109114016.png)

This is a clocked S-R flip-flop, where the output may only change in response to a clock pulse.

![Pasted image 20240109114107.png](pasted_image_20240109114107.png)

This is a D-type flip-flop, who only has one input to prevent S and R from being both 1, which would result in $Q$ and $\bar Q$ both being 0, which is inconsistent. Its input is delayed for a single clock pulse.

Flip-flops can be used to construct [Registers](registers)



























#### Why:
#### How:









