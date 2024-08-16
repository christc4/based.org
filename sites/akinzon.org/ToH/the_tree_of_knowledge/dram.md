# DRAM - Dynamic RAM

Up: [RAM](ram)
Brother(s): [SRAM](sram)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems


Dynamic RAM stores bits as a charge using capacitors. It is used in general main memory.

These capacitors slowly leak said charge and need refreshing to maintain their contents.

To refresh DRAM, the current value inside the capacitor must be read and then rewritten.

A benefit of DRAM's use of capacitors in its architecture is that it allows it to be more dense.

Its also cheaper and smaller, but is slower than [SRAM|Static RAM](sram|static_ram), with a speed of 6-60 ns.

![Pasted image 20231123192441.png](pasted_image_20231123192441.png)

All the DRAM together ^

![Pasted image 20231123192656.png](pasted_image_20231123192656.png)

^ This image represents a typical DRAM cell. To write to the cell, a voltage is applied to the bit line and a signal is also applied to the address line, which allows it to be stored in the capacitor.

^ To read from the cell, once the address line is selected the transistor turns on and the charge from the capacitor goes to the bit line and a sense amplifier, which checks if the signal is a logical 1 (high voltage) or logical 0 (low voltage). This process drains the capacitor, and so the value must be refreshed.

![Pasted image 20231123194742.png](pasted_image_20231123194742.png)

























#### Why:
#### How:









