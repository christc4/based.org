# Microcontrollers

Up: [Basic Concepts and History of the Computer](basic_concepts_and_history_of_the_computer)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Basic Concepts and History of the Computer

A microcontroller is a single chip that contains a processor, ROM memory for the program,  
RAM, and an I/O control unit. - self contained

![Pasted image 20231109200449.png](pasted_image_20231109200449.png)

Typiacally 32 bit, used to be 8 and 16 bit

They have much slower [The Clock|clocks](the_clock|clocks) from 32 kHz to 100 MHz

A benefit of this slow speed is their low power consumption
	> from microwatts to milliwatts
	> To put this into perspective, modern CPUs use ~2-280W

Small RAM, typically bytes/kilobytes of RAM, storage

With all this tiny storage/speed, they can be VERY CHEAP

What are some characteristics of a microcontroller #flashcard 
Microcontrollers are really cheap, use low power, and have comparatively small amounts of storage/[SRAM](sram).
<!--ID: 1701371854402-->


See [Microcontroller memory](microcontroller_memory)

Microcontrollers have longevity - the same build can be in production for years, the technology doesn't need to change much

**Microcontrollers generally have GPIO, serial interfaces, timers, etc.**

See [Microprocessors vs Microcontrollers](microprocessors_vs_microcontrollers)

**Usages of Microcontrollers** - Everywhere - Phones, embedded controllers, cameras, home appliances, etc.

Some microcontrollers can have Embedded OSes - FreeRTOS, Mbed OS, RIOT-OS

**IMPORTANT:** - [Microcontroller programming platforms|Microcontrollers need to be programmed to do anything](microcontroller_programming_platforms|microcontrollers_need_to_be_programmed_to_do_anything)

Conceptually similar to a [System on a Chip](system_on_a_chip)